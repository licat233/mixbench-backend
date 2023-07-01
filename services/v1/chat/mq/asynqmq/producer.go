package asynqmq

import (
	"context"
	"encoding/json"
	"mixbench/services/v1/chat/model"
	"mixbench/services/v1/chat/mq/types"

	"github.com/hibiken/asynq"
)

/**
 * producer 负责封装任务创建，向redis为容器的queue中放入task
 */
type AsynqProducer struct {
	ctx    context.Context
	Client *asynq.Client
}

func NewAsynqProducer(ctx context.Context, redisClientOpt *asynq.RedisClientOpt) *AsynqProducer {
	return &AsynqProducer{
		ctx:    ctx,
		Client: asynq.NewClient(redisClientOpt),
	}
}

func NewMessageTask(msg model.Message) (*asynq.Task, error) {
	payload, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(types.MessageTypename, payload), nil
}

func (a *AsynqProducer) AddTask(msg model.Message) error {
	task, err := NewMessageTask(msg)
	if err != nil {
		return err
	}
	_, err = a.Client.EnqueueContext(a.ctx, task)
	if err != nil {
		return err
	}
	return nil
}
