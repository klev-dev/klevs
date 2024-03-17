package server

import (
	"context"
	"io/fs"
	"path/filepath"
	"strings"

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

	paths := map[string]struct{}{}
	filepath.WalkDir(l.DataDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(d.Name(), ".log") {
			paths[path] = struct{}{}
		}
		return nil
	})

	logs := make([]*api.LogName, 0, len(paths))
	for p := range paths {
		rel, err := filepath.Rel(l.DataDir, p)
		if err != nil {
			return nil, err
		}

		dir, name := filepath.Split(rel)
		log := &api.LogName{Name: name}
		if len(dir) > 0 {
			log.Path = strings.Split(dir, string(filepath.Separator))
		}
		logs = append(logs, log)
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
