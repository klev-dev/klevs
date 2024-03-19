package grpc

import (
	"context"

	"go.klev.dev/klevs"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Messages struct {
	UnsafeMessagesServer
	*klevs.Logs
}

func (m *Messages) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {
	log, err := m.Logs.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	msgs := make([]klevs.Message, len(req.Messages))
	for i, m := range req.Messages {
		if m.Time != nil {
			msgs[i].Time = m.Time.AsTime()
		}
		msgs[i].Key = m.Key
		msgs[i].Value = m.Value
	}

	nextOffset, err := log.Publish(ctx, msgs)
	if err != nil {
		return nil, err
	}
	return &PublishResponse{NextOffset: nextOffset}, nil
}

func (m *Messages) NextOffset(ctx context.Context, req *NextOffsetRequest) (*NextOffsetResponse, error) {
	log, err := m.Logs.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	nextOffset, err := log.NextOffset(ctx)
	if err != nil {
		return nil, err
	}

	return &NextOffsetResponse{NextOffset: nextOffset}, nil
}

func (m *Messages) Consume(ctx context.Context, req *ConsumeRequest) (*ConsumeResponse, error) {
	log, err := m.Logs.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	nextOffset, msgs, err := log.Consume(ctx, req.Offset, req.MaxCount, 0)
	if err != nil {
		return nil, err
	}

	outMsgs := make([]*ConsumeMessage, len(msgs))
	for i, m := range msgs {
		outMsgs[i] = &ConsumeMessage{
			Offset: m.Offset,
			Time:   timestamppb.New(m.Time),
			Key:    m.Key,
			Value:  m.Value,
		}
	}

	return &ConsumeResponse{
		NextOffset: nextOffset,
		Messages:   outMsgs,
	}, nil
}

func (m *Messages) GetByOffset(ctx context.Context, req *GetByOffsetRequest) (*GetByOffsetResponse, error) {
	log, err := m.Logs.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	msg, err := log.GetByOffset(ctx, req.Offset)
	if err != nil {
		return nil, err
	}

	outMsg := &ConsumeMessage{
		Offset: msg.Offset,
		Time:   timestamppb.New(msg.Time),
		Key:    msg.Key,
		Value:  msg.Value,
	}

	return &GetByOffsetResponse{
		Message: outMsg,
	}, nil
}

func (m *Messages) GetByKey(ctx context.Context, req *GetByKeyRequest) (*GetByKeyResponse, error) {
	log, err := m.Logs.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	msg, err := log.GetByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	outMsg := &ConsumeMessage{
		Offset: msg.Offset,
		Time:   timestamppb.New(msg.Time),
		Key:    msg.Key,
		Value:  msg.Value,
	}

	return &GetByKeyResponse{
		Message: outMsg,
	}, nil
}

func (m *Messages) GetByTime(ctx context.Context, req *GetByTimeRequest) (*GetByTimeResponse, error) {
	log, err := m.Logs.Get(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	msg, err := log.GetByTime(ctx, req.Time.AsTime())
	if err != nil {
		return nil, err
	}

	outMsg := &ConsumeMessage{
		Offset: msg.Offset,
		Time:   timestamppb.New(msg.Time),
		Key:    msg.Key,
		Value:  msg.Value,
	}

	return &GetByTimeResponse{
		Message: outMsg,
	}, nil
}
