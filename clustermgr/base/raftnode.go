// Copyright 2022 The CubeFS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package base

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cubefs/blobstore/clustermgr/persistence/raftdb"
	"github.com/cubefs/blobstore/common/kvstore"
	"github.com/cubefs/blobstore/common/raftserver"
	"github.com/cubefs/blobstore/common/trace"
	"github.com/cubefs/blobstore/util/errors"
)

type RaftNodeConfig struct {
	FlushNumInterval    uint64            `json:"flush_num_interval"`
	FlushTimeIntervalS  int               `json:"flush_time_interval_s"`
	TruncateNumInterval uint64            `json:"truncate_num_interval"`
	NodeProtocol        string            `json:"node_protocol"`
	Nodes               map[uint64]string `json:"nodes"`
	ApplyIndex          uint64            `json:"-"`
}

type RaftNode struct {
	leaderHost string
	// currentApplyIndex is the memory current apply index, it's not stable
	currentApplyIndex uint64
	// stableApplyIndex is the persistent apply index and it's stable
	stableApplyIndex uint64
	// truncateApplyIndex record last truncated apply index
	truncateApplyIndex uint64
	// openSnapshotsNum record opening snapshot num currently
	openSnapshotsNum int32
	// appliers record all registered RaftApplier
	appliers []RaftApplier

	lock    sync.RWMutex
	closeCh chan interface{}
	raftDB  *raftdb.RaftDB

	raftserver.RaftServer
	*RaftNodeConfig
}

func NewRaftNode(cfg *RaftNodeConfig, raftDB *raftdb.RaftDB) (*RaftNode, error) {
	if cfg.FlushNumInterval == 0 {
		cfg.FlushNumInterval = defaultFlushNumInterval
	}
	if cfg.FlushTimeIntervalS == 0 {
		cfg.FlushTimeIntervalS = defaultFlushTimeIntervalS
	}
	if cfg.TruncateNumInterval == 0 {
		cfg.TruncateNumInterval = defaultTruncateNumInterval
	}

	raftNode := &RaftNode{
		raftDB:         raftDB,
		RaftNodeConfig: cfg,

		currentApplyIndex: cfg.ApplyIndex,
		stableApplyIndex:  cfg.ApplyIndex,
		// set truncateApplyIndex into last apply index - truncate num interval
		// it may not equal to the actual value, but it'll be fix by next truncation
		truncateApplyIndex: cfg.ApplyIndex - cfg.TruncateNumInterval,

		closeCh: make(chan interface{}),
	}

	return raftNode, nil
}

func (r *RaftNode) SetRaftServer(raftServer raftserver.RaftServer) {
	r.RaftServer = raftServer
}

// registRaftApplier use reflect to find out all RaftApplier and register
func (r *RaftNode) RegistRaftApplier(target interface{}) {
	// reflect all mgr's method, get the Applies and regist
	applies := make([]RaftApplier, 0)
	iface := reflect.TypeOf(new(RaftApplier)).Elem()
	vals := reflect.ValueOf(target).Elem()
	typs := vals.Type()
	for i := 0; i < vals.NumField(); i++ {
		field := typs.Field(i)
		if field.Type.Implements(iface) {
			applier := vals.Field(i).Interface().(RaftApplier)
			// set module name by reflect, no necessary to do it by self
			applier.SetModuleName(field.Name)
			applies = append(applies, applier)
		}
	}
	r.appliers = applies
}

func (r *RaftNode) GetStableApplyIndex() uint64 {
	return atomic.LoadUint64(&r.stableApplyIndex)
}

func (r *RaftNode) GetCurrentApplyIndex() uint64 {
	return atomic.LoadUint64(&r.currentApplyIndex)
}

func (r *RaftNode) RecordApplyIndex(ctx context.Context, index uint64, isFlush bool) (err error) {
	old := atomic.LoadUint64(&r.currentApplyIndex)
	if old < index {
		for {
			// update success, break
			if isSwap := atomic.CompareAndSwapUint64(&r.currentApplyIndex, old, index); isSwap {
				break
			}
			// already update, break
			old = atomic.LoadUint64(&r.currentApplyIndex)
			if old >= index {
				break
			}
			// otherwise, retry cas
		}
	}

	// no flush model, just record apply index into currentApplyIndex
	if !isFlush {
		return nil
	}
	err = r.flushAll(ctx)
	if err != nil {
		return
	}

	return r.saveStableApplyIndex(index)
}

func (r *RaftNode) NotifyLeaderChange(ctx context.Context, leader uint64, host string) {
	wg := sync.WaitGroup{}
	for i := range r.appliers {
		wg.Add(1)
		idx := i
		go func() {
			defer wg.Done()
			r.appliers[idx].NotifyLeaderChange(ctx, leader, host)
		}()
	}
	wg.Wait()
}

func (r *RaftNode) ModuleApply(ctx context.Context, module string, operTypes []int32, datas [][]byte, contexts []ProposeContext) error {
	moduleApplies := r.getApplierByModule(module)
	if moduleApplies == nil {
		return errors.New("raft node can't found applies in map")
	}

	err := moduleApplies.Apply(ctx, operTypes, datas, contexts)
	if err != nil {
		return errors.Info(err, "raft statemachine Apply call module method failed").Detail(err)
	}
	return nil
}

func (r *RaftNode) GetLeaderHost() string {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.leaderHost
}

func (r *RaftNode) SetLeaderHost(idx uint64, host string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.leaderHost = r.Nodes[idx]
}

func (r *RaftNode) CreateRaftSnapshot(dbs map[string]SnapshotDB, patchNum int) raftserver.Snapshot {
	applyIndex := r.GetStableApplyIndex()
	items := make([]snapshotItem, 0)
	for dbName := range dbs {
		cfs := dbs[dbName].GetAllCfNames()
		if len(cfs) == 0 {
			snap := dbs[dbName].NewSnapshot()
			iter := dbs[dbName].NewIterator(snap)
			iter.SeekToFirst()
			items = append(items, snapshotItem{DbName: dbName, snap: snap, iter: iter})
			continue
		}
		for i := range cfs {
			snap := dbs[dbName].Table(cfs[i]).NewSnapshot()
			iter := dbs[dbName].Table(cfs[i]).NewIterator(snap)
			iter.SeekToFirst()
			items = append(items, snapshotItem{DbName: dbName, CfName: cfs[i], snap: snap, iter: iter})
		}
	}
	// atomic add openSnapshotsNum
	atomic.AddInt32(&r.openSnapshotsNum, 1)

	return &raftSnapshot{
		name:          "snapshot-" + strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(100000)),
		items:         items,
		dbs:           dbs,
		applyIndex:    applyIndex,
		patchNum:      patchNum,
		closeCallback: r.closeSnapshotCallback,
	}
}

// ApplyRaftSnapshot apply snapshot's data into db
func (r *RaftNode) ApplyRaftSnapshot(ctx context.Context, dbs map[string]SnapshotDB, st raftserver.Snapshot) error {
	var (
		err   error
		data  []byte
		count uint64
	)
	span := trace.SpanFromContextSafe(ctx)

	for data, err = st.Read(); err == nil; data, err = st.Read() {
		reader := bytes.NewBuffer(data)
		for {
			snapData, err := decodeSnapshotData(reader)
			if err != nil {
				if err == io.EOF {
					break
				}
				span.Errorf("ApplyRaftSnapshot decode snapshot data failed, src snapshot data: %v, err: %v", data, err)
				return err
			}
			count++
			dbName := snapData.Header.DbName
			cfName := snapData.Header.CfName

			if snapData.Header.CfName != "" {
				err = dbs[dbName].Table(cfName).Put(kvstore.KV{Key: snapData.Key, Value: snapData.Value})
			} else {
				err = dbs[dbName].Put(kvstore.KV{Key: snapData.Key, Value: snapData.Value})
			}
			if err != nil {
				span.Errorf("ApplyRaftSnapshot put snapshot data failed, snapshot data: %v, err: %v", snapData, err)
				return err
			}
		}
	}
	span.Infof("apply snapshot read data count: %d", count)
	if err != io.EOF {
		span.Errorf("ApplyRaftSnapshot read unexpected error, err: %v", err)
		return err
	}
	// applier LoadData callback
	for _, applier := range r.appliers {
		if err := applier.LoadData(ctx); err != nil {
			span.Errorf("applier[%s] load data failed, err: %s", applier.GetModuleName(), err.Error())
			return err
		}
	}

	return nil
}

func (r *RaftNode) Start() {
	span, ctx := trace.StartSpanFromContextWithTraceID(context.Background(), "", "raft-node-loop")
	ticker := time.NewTicker(time.Duration(defaultFlushCheckIntervalS) * time.Second)
	lastFlushTime := time.Now()
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			now := time.Now()
			current := atomic.LoadUint64(&r.currentApplyIndex)
			stable := r.GetStableApplyIndex()
			if current > stable+r.FlushNumInterval || time.Since(lastFlushTime) > time.Duration(r.FlushTimeIntervalS)*time.Second {
				err := r.flushAll(ctx)
				span.Infof("raft node flush all cost: %d ms", time.Since(now).Milliseconds())
				lastFlushTime = time.Now()
				if err != nil {
					span.Error("raft node loop flush all failed, err: ", err)
					break
				}
				err = r.saveStableApplyIndex(current)
				if err != nil {
					span.Error("raft node loop save stable apply index failed, err: ", err)
					break
				}
				// also, we can try to truncate wal log after stable apply index update
				if stable-r.truncateApplyIndex > r.TruncateNumInterval*2 && atomic.LoadInt32(&r.openSnapshotsNum) == 0 {
					truncatedIndex := stable - r.TruncateNumInterval
					err = r.Truncate(truncatedIndex)
					if err != nil {
						span.Errorf("raft node truncate wal log failed, stable[%d], truncate[%d], err: %s", stable, r.truncateApplyIndex, err.Error())
						break
					}
					r.truncateApplyIndex = truncatedIndex
				}
			}
		case <-r.closeCh:
			ctx.Done()
			return
		}
	}
}

func (r *RaftNode) Stop() {
	// stop background flush checker
	close(r.closeCh)
	time.Sleep(1 * time.Second)
	// stop raft server
	r.RaftServer.Stop()
	r.raftDB.Close()
}

func (r *RaftNode) saveStableApplyIndex(new uint64) error {
	old := atomic.LoadUint64(&r.stableApplyIndex)
	if old >= new {
		return nil
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	// double check
	old = atomic.LoadUint64(&r.stableApplyIndex)
	if old >= new {
		return nil
	}

	indexValue := make([]byte, 8)
	binary.BigEndian.PutUint64(indexValue, new)

	err := r.raftDB.Put(ApplyIndexKey, indexValue)
	if err != nil {
		return errors.Info(err, "put flush apply index failed").Detail(err)
	}
	atomic.StoreUint64(&r.stableApplyIndex, new)

	return nil
}

// FlushAll will call all applier's flush method and record flush_apply_index into persistent storage
func (r *RaftNode) flushAll(ctx context.Context) error {
	wg := sync.WaitGroup{}
	errs := make([]error, len(r.appliers))
	for i := range r.appliers {
		idx := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			errs[idx] = r.appliers[idx].Flush(ctx)
		}()
	}
	wg.Wait()
	for i := range errs {
		if errs[i] != nil {
			return errors.Info(errs[i], fmt.Sprintf("flush applier %d failed", i)).Detail(errs[i])
		}
	}
	return nil
}

func (r *RaftNode) closeSnapshotCallback() {
	atomic.AddInt32(&r.openSnapshotsNum, -1)
}

func (r *RaftNode) getApplierByModule(module string) RaftApplier {
	for _, applier := range r.appliers {
		if applier.GetModuleName() == module {
			return applier
		}
	}
	return nil
}
