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

package client

import (
	"context"

	"github.com/cubefs/blobstore/api/blobnode"
	"github.com/cubefs/blobstore/common/proto"
)

// BlobnodeAPI interface of blobnode client deleter api
type BlobnodeAPI interface {
	MarkDelete(ctx context.Context, location proto.VunitLocation, bid proto.BlobID) error
	Delete(ctx context.Context, location proto.VunitLocation, bid proto.BlobID) error
}

type blobnodeClient struct {
	client blobnode.StorageAPI
}

// NewBlobnodeClient returns blobnode client
func NewBlobnodeClient(cfg *blobnode.Config) BlobnodeAPI {
	return &blobnodeClient{blobnode.New(cfg)}
}

// MarkDelete mark delete blob
func (c *blobnodeClient) MarkDelete(ctx context.Context, location proto.VunitLocation, bid proto.BlobID) error {
	return c.client.MarkDeleteShard(ctx, location.Host, &blobnode.DeleteShardArgs{
		DiskID: location.DiskID,
		Vuid:   location.Vuid,
		Bid:    bid,
	})
}

// Delete delete blob
func (c *blobnodeClient) Delete(ctx context.Context, location proto.VunitLocation, bid proto.BlobID) error {
	return c.client.DeleteShard(ctx, location.Host, &blobnode.DeleteShardArgs{
		DiskID: location.DiskID,
		Vuid:   location.Vuid,
		Bid:    bid,
	})
}
