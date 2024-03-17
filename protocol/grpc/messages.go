package grpc

import (
	"context"

	"go.klev.dev/klevs/server"
)

type Messages struct {
	UnsafeMessagesServer
	*server.Server
}

func (m *Messages) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *Messages) Consume(ctx context.Context, req *ConsumeRequest) (*ConsumeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
