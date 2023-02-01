/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-01-14 16:08:22
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/service/v1/admin/rpc/internal/config"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	RedisCli *redis.Redis

	ProjectModel  model.ProjectModel
	ActivityModel model.ActivityModel
	PrizeModel    model.PrizeModel
	UserModel     model.UserModel

	LuckyModel      model.LuckyModel
	PrizePoolModel  model.PrizePoolModel
	UserPoolModel   model.UserPoolModel
	PrizeGradeModel model.PrizeGradeModel

	CommonModel *model.CommonModel

	AdminerModel      model.AdminerModel
	JwtBlacklistModel model.JwtBlacklistModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	svc := &ServiceContext{
		Config: c,
		RedisCli: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
		ProjectModel:      model.NewProjectModel(sqlConn, c.CacheRedis),
		ActivityModel:     model.NewActivityModel(sqlConn, c.CacheRedis),
		PrizeModel:        model.NewPrizeModel(sqlConn, c.CacheRedis),
		UserModel:         model.NewUserModel(sqlConn, c.CacheRedis),
		LuckyModel:        model.NewLuckyModel(sqlConn, c.CacheRedis),
		PrizePoolModel:    model.NewPrizePoolModel(sqlConn, c.CacheRedis),
		UserPoolModel:     model.NewUserPoolModel(sqlConn, c.CacheRedis),
		PrizeGradeModel:   model.NewPrizeGradeModel(sqlConn, c.CacheRedis),
		AdminerModel:      model.NewAdminerModel(sqlConn, c.CacheRedis),
		JwtBlacklistModel: model.NewJwtBlacklistModel(sqlConn, c.CacheRedis),
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
