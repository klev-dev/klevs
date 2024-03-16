package server

import (
	"context"

	"go.klev.dev/klevs/api"
)

type Logs struct {
	api.UnsafeLogsServer
	DataDir string
}

func (l *Logs) List(ctx context.Context, req *api.LogsListRequest) (*api.LogsListResponse, error) {
	return nil, nil
}

func (l *Logs) Create(ctx context.Context, req *api.LogsCreateRequest) (*api.LogsCreateResponse, error) {
	return nil, nil
}

func (l *Logs) Delete(ctx context.Context, req *api.LogsDeleteRequest) (*api.LogsDeleteResponse, error) {
	return nil, nil
}
