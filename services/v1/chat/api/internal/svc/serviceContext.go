package svc

import (
	"context"
	"errors"
	"fmt"
	"log"
	"mixbench/common/utils"
	adminClient "mixbench/services/v1/admin/rpc/client"
	authenticationClient "mixbench/services/v1/authentication/rpc/client"
	"mixbench/services/v1/chat/api/internal/config"
	"mixbench/services/v1/chat/api/internal/middleware"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	goredis "github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	ctx    context.Context
	Config config.Config

	VisitorAuth rest.Middleware

	AuthenticationRpc *authenticationClient.AuthenticationRpc
	AdminRpc          *adminClient.AdminRpc

	// ChatRecordModel    model.ChatRecordModel
	// IpblackModel       model.IpblackModel
	// KefuModel          model.KefuModel
	MessageModel       model.MessageModel
	MessageConfigModel model.MessageConfigModel
	VisitorModel       model.VisitorModel
	// ProjectModel       model.ProjectModel

	// KqPusherClient *kq.Pusher
	// AsynqClient    *asynqmq.AsynqProducer

	Redis    *redis.Redis
	RedisCli *goredis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()
	sqlConn, err := utils.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatal(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	redisCli, err := redis.NewRedis(c.Redis.RedisConf)
	if err != nil {
		log.Fatal(fmt.Errorf("NewRedis失败: %s", err))
	}
	if !redisCli.Ping() {
		log.Fatal(errors.New("连接redis数据库失败"))
	}

	goredisCli := goredis.NewClient(&goredis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.RedisDB,
	})
	_, err = goredisCli.Ping(ctx).Result()
	if err != nil {
		log.Fatal(fmt.Errorf("连接redis数据库失败: %s", err))
	}

	svc := &ServiceContext{
		ctx:                ctx,
		Config:             c,
		VisitorAuth:        middleware.NewVisitorAuthMiddleware().Handle,
		AuthenticationRpc:  authenticationClient.NewAuthenticationRpc(zrpc.MustNewClient(c.AuthenticationRpc)),
		AdminRpc:           adminClient.NewAdminRpc(zrpc.MustNewClient(c.AdminRpc)),
		MessageModel:       model.NewMessageModel(sqlConn, c.CacheRedis),
		MessageConfigModel: model.NewMessageConfigModel(sqlConn, c.CacheRedis),
		VisitorModel:       model.NewVisitorModel(sqlConn, c.CacheRedis),
		Redis:              redisCli,
		RedisCli:           goredisCli,
	}
	return svc
}
