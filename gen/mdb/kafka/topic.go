// Code generated by sdkgen. DO NOT EDIT.

//nolint
package kafka

import (
	"context"

	"google.golang.org/grpc"

	kafka "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/kafka/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// TopicServiceClient is a kafka.TopicServiceClient with
// lazy GRPC connection initialization.
type TopicServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// Create implements kafka.TopicServiceClient
func (c *TopicServiceClient) Create(ctx context.Context, in *kafka.CreateTopicRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewTopicServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements kafka.TopicServiceClient
func (c *TopicServiceClient) Delete(ctx context.Context, in *kafka.DeleteTopicRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewTopicServiceClient(conn).Delete(ctx, in, opts...)
}

// Get implements kafka.TopicServiceClient
func (c *TopicServiceClient) Get(ctx context.Context, in *kafka.GetTopicRequest, opts ...grpc.CallOption) (*kafka.Topic, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewTopicServiceClient(conn).Get(ctx, in, opts...)
}

// List implements kafka.TopicServiceClient
func (c *TopicServiceClient) List(ctx context.Context, in *kafka.ListTopicsRequest, opts ...grpc.CallOption) (*kafka.ListTopicsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewTopicServiceClient(conn).List(ctx, in, opts...)
}

type TopicIterator struct {
	ctx  context.Context
	opts []grpc.CallOption

	err     error
	started bool

	client  *TopicServiceClient
	request *kafka.ListTopicsRequest

	items []*kafka.Topic
}

func (c *TopicServiceClient) TopicIterator(ctx context.Context, clusterId string, opts ...grpc.CallOption) *TopicIterator {
	return &TopicIterator{
		ctx:    ctx,
		opts:   opts,
		client: c,
		request: &kafka.ListTopicsRequest{
			ClusterId: clusterId,
			PageSize:  1000,
		},
	}
}

func (it *TopicIterator) Next() bool {
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

	it.items = response.Topics
	it.request.PageToken = response.NextPageToken
	return len(it.items) > 0
}

func (it *TopicIterator) Value() *kafka.Topic {
	if len(it.items) == 0 {
		panic("calling Value on empty iterator")
	}
	return it.items[0]
}

func (it *TopicIterator) Error() error {
	return it.err
}

// Update implements kafka.TopicServiceClient
func (c *TopicServiceClient) Update(ctx context.Context, in *kafka.UpdateTopicRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return kafka.NewTopicServiceClient(conn).Update(ctx, in, opts...)
}
