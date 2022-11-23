package svc

import (
	"context"
	"mixbench/common/asynqmq"
	"mixbench/service/v1/luckdraw/rpc/internal/config"
	"mixbench/service/v1/model"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	ProjectModel  model.ProjectModel
	ActivityModel model.ActivityModel
	PrizeModel    model.PrizeModel
	UserModel     model.UserModel

	LuckerModel     model.LuckerModel
	PrizesPoolModel model.PrizesPoolModel
	UsersPoolModel  model.UsersPoolModel
	PrizeGradeModel model.PrizeGradeModel

	CommonModel *model.CommonModel

	KqueueGoodluckStatusHandlerClient *kq.Pusher
	AsynqClient                       *asynqmq.AsynqProducer
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	svc := &ServiceContext{
		Config:          c,
		ProjectModel:    model.NewProjectModel(sqlConn, c.CacheRedis),
		ActivityModel:   model.NewActivityModel(sqlConn, c.CacheRedis),
		PrizeModel:      model.NewPrizeModel(sqlConn, c.CacheRedis),
		UserModel:       model.NewUserModel(sqlConn, c.CacheRedis),
		LuckerModel:     model.NewLuckerModel(sqlConn, c.CacheRedis),
		PrizesPoolModel: model.NewPrizesPoolModel(sqlConn, c.CacheRedis),
		UsersPoolModel:  model.NewUsersPoolModel(sqlConn, c.CacheRedis),
		PrizeGradeModel: model.NewPrizeGradeModel(sqlConn, c.CacheRedis),

		CommonModel: nil,

		// KqueueGoodluckStatusHandlerClient: kq.NewPusher(c.KqGoodluckConf.Brokers, c.KqGoodluckConf.Topic),
		KqueueGoodluckStatusHandlerClient: nil, //这里放弃使用kakfa，用asynq来代替
		AsynqClient: asynqmq.NewAsynqProducer(ctx, &asynq.RedisClientOpt{
			Addr:     c.Redis.Host,
			Password: c.Redis.Pass,
			DB:       asynqmq.RedisDB,
		}),
	}
	svc.CommonModel = &model.CommonModel{
		ProjectModel:    svc.ProjectModel,
		ActivityModel:   svc.ActivityModel,
		PrizeModel:      svc.PrizeModel,
		UserModel:       svc.UserModel,
		LuckerModel:     svc.LuckerModel,
		PrizesPoolModel: svc.PrizesPoolModel,
		UsersPoolModel:  svc.UsersPoolModel,
		PrizeGradeModel: svc.PrizeGradeModel,
	}
	return svc
}
