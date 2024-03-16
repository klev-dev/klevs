package server

import (
	"context"

	"go.klev.dev/klevs/api"
)

type Messages struct {
	api.UnsafeMessagesServer
	DataDir string
}

func (m *Messages) Publish(ctx context.Context, req *api.PublishRequest) (*api.PublishResponse, error) {
	return nil, nil
}

func (m *Messages) Consume(ctx context.Context, req *api.ConsumeRequest) (*api.ConsumeResponse, error) {
	return nil, nil
}
