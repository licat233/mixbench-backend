package producer

import (
	"context"
	"mixbench/services/v1/luckydraw/asynqmq"

	"github.com/hibiken/asynq"
)

func NewAsynqProducer(ctx context.Context, redisClientOpt *asynq.RedisClientOpt) *asynqmq.AsynqProducer {
	return asynqmq.NewAsynqProducer(ctx, redisClientOpt)
}
