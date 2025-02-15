// Code generated by sdkgen. DO NOT EDIT.

//nolint
package ydb

import (
	"context"

	"google.golang.org/grpc"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	ydb "github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1"
)

//revive:disable

// DatabaseServiceClient is a ydb.DatabaseServiceClient with
// lazy GRPC connection initialization.
type DatabaseServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Backup implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Backup(ctx context.Context, in *ydb.BackupDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Backup(ctx, in, opts...)
}

// Create implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Create(ctx context.Context, in *ydb.CreateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Delete(ctx context.Context, in *ydb.DeleteDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Get(ctx context.Context, in *ydb.GetDatabaseRequest, opts ...grpc.CallOption) (*ydb.Database, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Get(ctx, in, opts...)
}

// List implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) List(ctx context.Context, in *ydb.ListDatabasesRequest, opts ...grpc.CallOption) (*ydb.ListDatabasesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).List(ctx, in, opts...)
}

type DatabaseIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DatabaseServiceClient
	request *ydb.ListDatabasesRequest

	items []*ydb.Database
}

func (c *DatabaseServiceClient) DatabaseIterator(ctx context.Context, req *ydb.ListDatabasesRequest, opts ...grpc.CallOption) *DatabaseIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DatabaseIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *DatabaseIterator) Next() bool {
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

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.List(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.Databases
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DatabaseIterator) Take(size int64) ([]*ydb.Database, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*ydb.Database

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DatabaseIterator) TakeAll() ([]*ydb.Database, error) {
	return it.Take(0)
}

func (it *DatabaseIterator) Value() *ydb.Database {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DatabaseIterator) Error() error {
	return it.err
}

// ListAccessBindings implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).ListAccessBindings(ctx, in, opts...)
}

type DatabaseAccessBindingsIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err           error
	started       bool
	requestedSize int64
	pageSize      int64

	client  *DatabaseServiceClient
	request *access.ListAccessBindingsRequest

	items []*access.AccessBinding
}

func (c *DatabaseServiceClient) DatabaseAccessBindingsIterator(ctx context.Context, req *access.ListAccessBindingsRequest, opts ...grpc.CallOption) *DatabaseAccessBindingsIterator {
	var pageSize int64
	const defaultPageSize = 1000
	pageSize = req.PageSize
	if pageSize == 0 {
		pageSize = defaultPageSize
	}
	return &DatabaseAccessBindingsIterator{
		ctx:      ctx,
		opts:     opts,
		client:   c,
		request:  req,
		pageSize: pageSize,
	}
}

func (it *DatabaseAccessBindingsIterator) Next() bool {
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

	if it.requestedSize == 0 || it.requestedSize > it.pageSize {
		it.request.PageSize = it.pageSize
	} else {
		it.request.PageSize = it.requestedSize
	}

	response, err := it.client.ListAccessBindings(it.ctx, it.request, it.opts...)
	it.err = err
	if err != nil {
		return false
	}

	it.items = response.AccessBindings
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *DatabaseAccessBindingsIterator) Take(size int64) ([]*access.AccessBinding, error) {
	if it.err != nil {
		return nil, it.err
	}

	if size == 0 {
		size = 1 << 32 // something insanely large
	}
	it.requestedSize = size
	defer func() {
		// reset iterator for future calls.
		it.requestedSize = 0
	}()

	var result []*access.AccessBinding

	for it.requestedSize > 0 && it.Next() {
		it.requestedSize--
		result = append(result, it.Value())
	}

	if it.err != nil {
		return nil, it.err
	}

	return result, nil
}

func (it *DatabaseAccessBindingsIterator) TakeAll() ([]*access.AccessBinding, error) {
	return it.Take(0)
}

func (it *DatabaseAccessBindingsIterator) Value() *access.AccessBinding {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *DatabaseAccessBindingsIterator) Error() error {
	return it.err
}

// Restore implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Restore(ctx context.Context, in *ydb.RestoreBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Restore(ctx, in, opts...)
}

// SetAccessBindings implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).SetAccessBindings(ctx, in, opts...)
}

// Start implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Start(ctx context.Context, in *ydb.StartDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Start(ctx, in, opts...)
}

// Stop implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Stop(ctx context.Context, in *ydb.StopDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) Update(ctx context.Context, in *ydb.UpdateDatabaseRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).Update(ctx, in, opts...)
}

// UpdateAccessBindings implements ydb.DatabaseServiceClient
func (c *DatabaseServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return ydb.NewDatabaseServiceClient(conn).UpdateAccessBindings(ctx, in, opts...)
}
