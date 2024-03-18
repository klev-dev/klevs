package klevs

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/klev-dev/klevdb"
	"github.com/klev-dev/klevdb/message"
)

type Message = message.Message

type Log struct {
	logDir string
	db     klevdb.Log

	closed   bool
	deleted  bool
	deleteMu sync.RWMutex
}

func NewLog(logDir string) (*Log, error) {
	db, err := klevdb.Open(logDir, klevdb.Options{})
	if err != nil {
		return nil, err
	}
	return &Log{
		logDir: logDir,
		db:     db,
	}, nil
}

func (l *Log) Publish(ctx context.Context, msgs []Message) (int64, error) {
	l.deleteMu.RLock()
	defer l.deleteMu.RUnlock()

	switch {
	case l.deleted:
		return message.OffsetInvalid, fmt.Errorf("topic already deleted")
	case l.closed:
		return message.OffsetInvalid, fmt.Errorf("topic already closed")
	}

	return l.db.Publish(msgs)
}

func (l *Log) Consume(ctx context.Context, offset int64, maxCount int64, poll time.Duration) (int64, []Message, error) {
	l.deleteMu.RLock()
	defer l.deleteMu.RUnlock()

	switch {
	case l.deleted:
		return message.OffsetInvalid, nil, fmt.Errorf("topic already deleted")
	case l.closed:
		return message.OffsetInvalid, nil, fmt.Errorf("topic already closed")
	}

	return l.db.Consume(offset, maxCount)
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
