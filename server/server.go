package server

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/klev-dev/klevdb"
	"golang.org/x/exp/maps"
)

type Server struct {
	dataDir string

	logs   map[string]klevdb.Log
	logsMu sync.RWMutex
}

func New(dataDir string) (*Server, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}

	logs := make(map[string]klevdb.Log, len(entries))
	for _, en := range entries {
		if en.IsDir() {
			logDir := filepath.Join(dataDir, en.Name())
			log, err := klevdb.Open(logDir, klevdb.Options{})
			if err != nil {
				return nil, err
			}
			logs[en.Name()] = log
		}
	}

	return &Server{
		dataDir: dataDir,
		logs:    logs,
	}, nil
}

func (s *Server) List(ctx context.Context) ([]string, error) {
	s.logsMu.RLock()
	defer s.logsMu.RUnlock()

	return maps.Keys(s.logs), nil
}

func (s *Server) Create(ctx context.Context, name string) (klevdb.Log, error) {
	s.logsMu.Lock()
	defer s.logsMu.Unlock()

	if _, ok := s.logs[name]; ok {
		return nil, fmt.Errorf("log '%s' already exists", name)
	}

	logDir := filepath.Join(s.dataDir, name)
	log, err := klevdb.Open(logDir, klevdb.Options{
		CreateDirs: true,
	})
	if err != nil {
		return nil, err
	}
	s.logs[name] = log
	return log, nil
}

func (s *Server) Delete(ctx context.Context, name string) error {
	return nil
}
