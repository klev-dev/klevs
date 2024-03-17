package server

import (
	"context"
	"os"

	"go.klev.dev/klevs/api"
)

type Logs struct {
	api.UnsafeLogsServer
	DataDir string
}

func (l *Logs) List(ctx context.Context, req *api.LogsListRequest) (*api.LogsListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(l.DataDir)
	if err != nil {
		return nil, err
	}

	logs := make([]string, 0, len(entries))
	for _, en := range entries {
		if en.IsDir() {
			logs = append(logs, en.Name())
		}
	}

	return &api.LogsListResponse{Logs: logs}, nil
}

func (l *Logs) Create(ctx context.Context, req *api.LogsCreateRequest) (*api.LogsCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *Logs) Delete(ctx context.Context, req *api.LogsDeleteRequest) (*api.LogsDeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
