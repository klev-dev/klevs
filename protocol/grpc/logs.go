package grpc

import (
	"context"

	"go.klev.dev/klevs"
)

type Logs struct {
	UnsafeLogsServer
	*klevs.Logs
}

func (l *Logs) List(ctx context.Context, req *LogsListRequest) (*LogsListResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	logs, err := l.Logs.List(ctx)
	if err != nil {
		return nil, err
	}

	return &LogsListResponse{Logs: logs}, nil
}

func (l *Logs) Create(ctx context.Context, req *LogsCreateRequest) (*LogsCreateResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	_, err := l.Logs.Create(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &LogsCreateResponse{}, nil
}

func (l *Logs) Delete(ctx context.Context, req *LogsDeleteRequest) (*LogsDeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if err := l.Logs.Delete(ctx, req.Name); err != nil {
		return nil, err
	}

	return &LogsDeleteResponse{}, nil
}
