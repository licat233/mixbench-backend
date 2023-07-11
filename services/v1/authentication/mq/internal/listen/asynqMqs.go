package listen

import (
	"context"
	"mixbench/services/v1/authentication/mq/internal/mqs/aq/consumer"
	"mixbench/services/v1/authentication/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use asynq
func AsynqMqs(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		consumer.NewAsynqConsumer(ctx, svcCtx),
	}
}
