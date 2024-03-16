package server

import (
	"context"

	"go.klev.dev/klevs/api"
)

type Logs struct {
	api.UnimplementedLogsServer
}

func (l *Logs) List(ctx context.Context, req *api.LogsListReq) (*api.LogsListResp, error) {
	return nil, nil
}
