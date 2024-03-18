package klevs

import (
	"context"
	"time"

	"github.com/klev-dev/klevdb"
	"github.com/klev-dev/klevdb/message"
)

type Message = message.Message

type Log struct {
	db klevdb.Log
}

func (l *Log) Publish(ctx context.Context, msgs []Message) (int64, error) {
	return l.db.Publish(msgs)
}

func (l *Log) Consume(ctx context.Context, offset int64, maxCount int64, poll time.Duration) (int64, []Message, error) {
	return l.db.Consume(offset, maxCount)
}
