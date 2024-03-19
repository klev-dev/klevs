package klevs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/klev-dev/klevdb"
)

type Message = klevdb.Message
type Stats = klevdb.Stats

type Log struct {
	logDir string
	db     klevdb.Log

	closed   bool
	deleted  bool
	deleteMu sync.RWMutex
}

func NewLog(logDir string) (*Log, error) {
	db, err := klevdb.Open(logDir, klevdb.Options{
		CreateDirs: true,
	})
	if err != nil {
		return nil, err
	}
	return &Log{
		logDir: logDir,
		db:     db,
	}, nil
}

func (l *Log) rlockAndValidate() error {
	l.deleteMu.RLock()

	switch {
	case l.deleted:
		l.deleteMu.RUnlock()
		return fmt.Errorf("topic already deleted")
	case l.closed:
		l.deleteMu.RUnlock()
		return fmt.Errorf("topic already closed")
	}

	return nil
}

func (l *Log) Publish(ctx context.Context, msgs []Message) (int64, error) {
	if err := l.rlockAndValidate(); err != nil {
		return klevdb.OffsetInvalid, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.Publish(msgs)
}

func (l *Log) NextOffset(ctx context.Context) (int64, error) {
	if err := l.rlockAndValidate(); err != nil {
		return klevdb.OffsetInvalid, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.NextOffset()
}

func (l *Log) Consume(ctx context.Context, offset int64, maxCount int64, poll time.Duration) (int64, []Message, error) {
	if err := l.rlockAndValidate(); err != nil {
		return klevdb.OffsetInvalid, nil, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.Consume(offset, maxCount)
}

func (l *Log) GetByOffset(ctx context.Context, offset int64) (Message, error) {
	if err := l.rlockAndValidate(); err != nil {
		return klevdb.InvalidMessage, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.Get(offset)
}

func (l *Log) GetByKey(ctx context.Context, key []byte) (Message, error) {
	if err := l.rlockAndValidate(); err != nil {
		return klevdb.InvalidMessage, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.GetByKey(key)
}

func (l *Log) GetByTime(ctx context.Context, start time.Time) (Message, error) {
	if err := l.rlockAndValidate(); err != nil {
		return klevdb.InvalidMessage, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.GetByTime(start)
}

func (l *Log) Stat(ctx context.Context) (Stats, error) {
	if err := l.rlockAndValidate(); err != nil {
		return Stats{}, err
	}
	defer l.deleteMu.RUnlock()

	return l.db.Stat()
}

func (l *Log) Delete(ctx context.Context) error {
	l.deleteMu.Lock()
	defer l.deleteMu.Unlock()

	if l.deleted {
		return fmt.Errorf("topic already deleted")
	}

	if !l.closed {
		if err := l.db.Close(); err != nil {
			return err
		}
		l.closed = true
	}

	if err := os.RemoveAll(l.logDir); err != nil {
		return err
	}
	l.deleted = true
	return nil
}
