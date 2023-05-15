package listen

import (
	"context"
	"mixbench/services/v1/luckydraw/mq/internal/config"
	mqKq "mixbench/services/v1/luckydraw/mq/internal/mqs/kq"
	"mixbench/services/v1/luckydraw/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// pub sub use kq (kafka)
func KqMqs(c config.Config, ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(c.GoodluckConf, mqKq.NewGoodluckStatusMq(ctx, svcCtx)),
	}
}
