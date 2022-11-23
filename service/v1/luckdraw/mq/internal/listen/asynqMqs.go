package listen

import (
	"context"
	"mixbench/service/v1/luckdraw/mq/internal/mqs/aq/consumer"
	"mixbench/service/v1/luckdraw/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use asynq
func AsynqMqs(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		consumer.NewAsynqConsumer(ctx, svcCtx),
	}
}
