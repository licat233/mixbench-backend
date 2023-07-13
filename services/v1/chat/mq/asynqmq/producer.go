package asynqmq

import (
	"context"
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

func (a *AsynqProducer) AddTask(task *asynq.Task) error {
	_, err := a.Client.EnqueueContext(a.ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func NewSendMessageTask(payload []byte) (*asynq.Task, error) {
	return asynq.NewTask(types.TaskTypeSendMessage, payload), nil
}
