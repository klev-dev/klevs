package klevs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sync"

	"golang.org/x/exp/maps"
)

type Logs struct {
	dataDir string

	logs   map[string]*Log
	logsMu sync.RWMutex
}

func New(dataDir string) (*Logs, error) {
	entries, err := os.ReadDir(dataDir)
	if err != nil {
		return nil, err
	}

	logs := make(map[string]*Log, len(entries))
	for _, en := range entries {
		if en.IsDir() {
			logDir := filepath.Join(dataDir, en.Name())
			logs[en.Name()], err = NewLog(logDir)
			if err != nil {
				return nil, err
			}
		}
	}

	return &Logs{
		dataDir: dataDir,
		logs:    logs,
	}, nil
}

func (s *Logs) List(ctx context.Context) ([]string, error) {
	s.logsMu.RLock()
	defer s.logsMu.RUnlock()

	logs := maps.Keys(s.logs)
	slices.Sort(logs)
	return logs, nil
}

func (s *Logs) Create(ctx context.Context, name string) (*Log, error) {
	s.logsMu.Lock()
	defer s.logsMu.Unlock()

	if _, ok := s.logs[name]; ok {
		return nil, fmt.Errorf("log '%s' already exists", name)
	}

	logDir := filepath.Join(s.dataDir, name)
	log, err := NewLog(logDir)
	if err != nil {
		return nil, err
	}

	s.logs[name] = log
	return log, nil
}

func (s *Logs) Get(ctx context.Context, name string) (*Log, error) {
	s.logsMu.RLock()
	defer s.logsMu.RUnlock()

	if log, ok := s.logs[name]; ok {
		return log, nil
	}

	return nil, fmt.Errorf("log '%s' not found", name)
}

func (s *Logs) Delete(ctx context.Context, name string) error {
	log, err := s.Get(ctx, name)
	if err != nil {
		return err
	}

	if err := log.Delete(ctx); err != nil {
		return err
	}

	s.logsMu.Lock()
	defer s.logsMu.Unlock()

	delete(s.logs, name)
	return nil
}
