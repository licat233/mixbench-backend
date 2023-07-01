/*
 * @Author: licat
 * @Date: 2023-01-30 16:38:35
 * @LastEditors: licat
 * @LastEditTime: 2023-01-30 18:11:08
 * @Description: licat233@gmail.com
 */
package svc

import (
	"context"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/mq/asynqmq"
	"mixbench/services/v1/luckydraw/rpc/internal/config"

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

	LuckyModel      model.LuckyModel
	PrizePoolModel  model.PrizePoolModel
	UserPoolModel   model.UserPoolModel
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
		LuckyModel:      model.NewLuckyModel(sqlConn, c.CacheRedis),
		PrizePoolModel:  model.NewPrizePoolModel(sqlConn, c.CacheRedis),
		UserPoolModel:   model.NewUserPoolModel(sqlConn, c.CacheRedis),
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
		LuckyModel:      svc.LuckyModel,
		PrizePoolModel:  svc.PrizePoolModel,
		UserPoolModel:   svc.UserPoolModel,
		PrizeGradeModel: svc.PrizeGradeModel,
	}
	return svc
}
