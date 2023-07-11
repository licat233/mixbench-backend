package listen

import (
	"context"
	"mixbench/services/v1/authentication/mq/internal/config"
	"mixbench/services/v1/authentication/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/service"
)

func Mqs(c config.Config) []service.Service {
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	var services []service.Service

	//kq: pub sub
	// services = append(services, KqMqs(c, ctx, svcCtx)...)  //暂时不需要使用 kafka

	//使用asynq代替
	services = append(services, AsynqMqs(ctx, svcCtx)...)
	return services
}
