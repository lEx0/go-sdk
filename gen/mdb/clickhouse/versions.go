// Code generated by sdkgen. DO NOT EDIT.

//nolint
package clickhouse

import (
	"context"

	"google.golang.org/grpc"

	clickhouse "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1"
)

//revive:disable

// VersionsServiceClient is a clickhouse.VersionsServiceClient with
// lazy GRPC connection initialization.
type VersionsServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// List implements clickhouse.VersionsServiceClient
func (c *VersionsServiceClient) List(ctx context.Context, in *clickhouse.ListVersionsRequest, opts ...grpc.CallOption) (*clickhouse.ListVersionsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return clickhouse.NewVersionsServiceClient(conn).List(ctx, in, opts...)
}

type VersionsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *VersionsServiceClient
	request *clickhouse.ListVersionsRequest

	items []*clickhouse.Version
}

func (c *VersionsServiceClient) VersionsIterator(ctx context.Context, opts ...grpc.CallOption) *VersionsIterator {
	return &VersionsIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &clickhouse.ListVersionsRequest{
			PageSize: 1000,
		},
	}
}

func (it *VersionsIterator) Next() bool {
	if it.err != nil {
		return false
	}
	if len(it.items) > 1 {
		it.items[0] = nil
		it.items = it.items[1:]
		return true
	}
	it.items = nil // consume last item, if any

	if it.started && it.request.PageToken == "" {
		return false
	}
	it.started = true

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Version
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *VersionsIterator) Value() *clickhouse.Version {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *VersionsIterator) Error() error {
	return it.err
}
