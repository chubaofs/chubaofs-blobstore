// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/blobstore/api/blobnode (interfaces: StorageAPI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	io "io"
	reflect "reflect"

	blobnode "github.com/cubefs/blobstore/api/blobnode"
	proto "github.com/cubefs/blobstore/common/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockStorageAPI is a mock of StorageAPI interface.
type MockStorageAPI struct {
	ctrl     *gomock.Controller
	recorder *MockStorageAPIMockRecorder
}

// MockStorageAPIMockRecorder is the mock recorder for MockStorageAPI.
type MockStorageAPIMockRecorder struct {
	mock *MockStorageAPI
}

// NewMockStorageAPI creates a new mock instance.
func NewMockStorageAPI(ctrl *gomock.Controller) *MockStorageAPI {
	mock := &MockStorageAPI{ctrl: ctrl}
	mock.recorder = &MockStorageAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageAPI) EXPECT() *MockStorageAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStorageAPI) Close(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageAPIMockRecorder) Close(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorageAPI)(nil).Close), arg0, arg1)
}

// CreateChunk mocks base method.
func (m *MockStorageAPI) CreateChunk(arg0 context.Context, arg1 string, arg2 *blobnode.CreateChunkArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChunk", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateChunk indicates an expected call of CreateChunk.
func (mr *MockStorageAPIMockRecorder) CreateChunk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChunk", reflect.TypeOf((*MockStorageAPI)(nil).CreateChunk), arg0, arg1, arg2)
}

// DeleteShard mocks base method.
func (m *MockStorageAPI) DeleteShard(arg0 context.Context, arg1 string, arg2 *blobnode.DeleteShardArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteShard indicates an expected call of DeleteShard.
func (mr *MockStorageAPIMockRecorder) DeleteShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShard", reflect.TypeOf((*MockStorageAPI)(nil).DeleteShard), arg0, arg1, arg2)
}

// DiskInfo mocks base method.
func (m *MockStorageAPI) DiskInfo(arg0 context.Context, arg1 string, arg2 *blobnode.DiskStatArgs) (*blobnode.DiskInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiskInfo", arg0, arg1, arg2)
	ret0, _ := ret[0].(*blobnode.DiskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DiskInfo indicates an expected call of DiskInfo.
func (mr *MockStorageAPIMockRecorder) DiskInfo(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiskInfo", reflect.TypeOf((*MockStorageAPI)(nil).DiskInfo), arg0, arg1, arg2)
}

// GetShard mocks base method.
func (m *MockStorageAPI) GetShard(arg0 context.Context, arg1 string, arg2 *blobnode.GetShardArgs) (io.ReadCloser, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetShard indicates an expected call of GetShard.
func (mr *MockStorageAPIMockRecorder) GetShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShard", reflect.TypeOf((*MockStorageAPI)(nil).GetShard), arg0, arg1, arg2)
}

// GetShards mocks base method.
func (m *MockStorageAPI) GetShards(arg0 context.Context, arg1 string, arg2 *blobnode.GetShardsArgs) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShards", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShards indicates an expected call of GetShards.
func (mr *MockStorageAPIMockRecorder) GetShards(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShards", reflect.TypeOf((*MockStorageAPI)(nil).GetShards), arg0, arg1, arg2)
}

// IsOnline mocks base method.
func (m *MockStorageAPI) IsOnline(arg0 context.Context, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOnline", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsOnline indicates an expected call of IsOnline.
func (mr *MockStorageAPIMockRecorder) IsOnline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnline", reflect.TypeOf((*MockStorageAPI)(nil).IsOnline), arg0, arg1)
}

// ListChunks mocks base method.
func (m *MockStorageAPI) ListChunks(arg0 context.Context, arg1 string, arg2 *blobnode.ListChunkArgs) ([]*blobnode.ChunkInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListChunks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*blobnode.ChunkInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListChunks indicates an expected call of ListChunks.
func (mr *MockStorageAPIMockRecorder) ListChunks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListChunks", reflect.TypeOf((*MockStorageAPI)(nil).ListChunks), arg0, arg1, arg2)
}

// ListShards mocks base method.
func (m *MockStorageAPI) ListShards(arg0 context.Context, arg1 string, arg2 *blobnode.ListShardsArgs) ([]*blobnode.ShardInfo, proto.BlobID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShards", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*blobnode.ShardInfo)
	ret1, _ := ret[1].(proto.BlobID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListShards indicates an expected call of ListShards.
func (mr *MockStorageAPIMockRecorder) ListShards(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShards", reflect.TypeOf((*MockStorageAPI)(nil).ListShards), arg0, arg1, arg2)
}

// MarkDeleteShard mocks base method.
func (m *MockStorageAPI) MarkDeleteShard(arg0 context.Context, arg1 string, arg2 *blobnode.DeleteShardArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkDeleteShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDeleteShard indicates an expected call of MarkDeleteShard.
func (mr *MockStorageAPIMockRecorder) MarkDeleteShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDeleteShard", reflect.TypeOf((*MockStorageAPI)(nil).MarkDeleteShard), arg0, arg1, arg2)
}

// PutShard mocks base method.
func (m *MockStorageAPI) PutShard(arg0 context.Context, arg1 string, arg2 *blobnode.PutShardArgs) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutShard indicates an expected call of PutShard.
func (mr *MockStorageAPIMockRecorder) PutShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutShard", reflect.TypeOf((*MockStorageAPI)(nil).PutShard), arg0, arg1, arg2)
}

// RangeGetShard mocks base method.
func (m *MockStorageAPI) RangeGetShard(arg0 context.Context, arg1 string, arg2 *blobnode.RangeGetShardArgs) (io.ReadCloser, uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RangeGetShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(uint32)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RangeGetShard indicates an expected call of RangeGetShard.
func (mr *MockStorageAPIMockRecorder) RangeGetShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RangeGetShard", reflect.TypeOf((*MockStorageAPI)(nil).RangeGetShard), arg0, arg1, arg2)
}

// ReleaseChunk mocks base method.
func (m *MockStorageAPI) ReleaseChunk(arg0 context.Context, arg1 string, arg2 *blobnode.ChangeChunkStatusArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseChunk", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseChunk indicates an expected call of ReleaseChunk.
func (mr *MockStorageAPIMockRecorder) ReleaseChunk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseChunk", reflect.TypeOf((*MockStorageAPI)(nil).ReleaseChunk), arg0, arg1, arg2)
}

// SetChunkReadonly mocks base method.
func (m *MockStorageAPI) SetChunkReadonly(arg0 context.Context, arg1 string, arg2 *blobnode.ChangeChunkStatusArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChunkReadonly", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChunkReadonly indicates an expected call of SetChunkReadonly.
func (mr *MockStorageAPIMockRecorder) SetChunkReadonly(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChunkReadonly", reflect.TypeOf((*MockStorageAPI)(nil).SetChunkReadonly), arg0, arg1, arg2)
}

// SetChunkReadwrite mocks base method.
func (m *MockStorageAPI) SetChunkReadwrite(arg0 context.Context, arg1 string, arg2 *blobnode.ChangeChunkStatusArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChunkReadwrite", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChunkReadwrite indicates an expected call of SetChunkReadwrite.
func (mr *MockStorageAPIMockRecorder) SetChunkReadwrite(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChunkReadwrite", reflect.TypeOf((*MockStorageAPI)(nil).SetChunkReadwrite), arg0, arg1, arg2)
}

// Stat mocks base method.
func (m *MockStorageAPI) Stat(arg0 context.Context, arg1 string) ([]*blobnode.DiskInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", arg0, arg1)
	ret0, _ := ret[0].([]*blobnode.DiskInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockStorageAPIMockRecorder) Stat(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockStorageAPI)(nil).Stat), arg0, arg1)
}

// StatChunk mocks base method.
func (m *MockStorageAPI) StatChunk(arg0 context.Context, arg1 string, arg2 *blobnode.StatChunkArgs) (*blobnode.ChunkInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatChunk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*blobnode.ChunkInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatChunk indicates an expected call of StatChunk.
func (mr *MockStorageAPIMockRecorder) StatChunk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatChunk", reflect.TypeOf((*MockStorageAPI)(nil).StatChunk), arg0, arg1, arg2)
}

// StatShard mocks base method.
func (m *MockStorageAPI) StatShard(arg0 context.Context, arg1 string, arg2 *blobnode.StatShardArgs) (*blobnode.ShardInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(*blobnode.ShardInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatShard indicates an expected call of StatShard.
func (mr *MockStorageAPIMockRecorder) StatShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatShard", reflect.TypeOf((*MockStorageAPI)(nil).StatShard), arg0, arg1, arg2)
}

// String mocks base method.
func (m *MockStorageAPI) String(arg0 context.Context, arg1 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String", arg0, arg1)
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockStorageAPIMockRecorder) String(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockStorageAPI)(nil).String), arg0, arg1)
}
