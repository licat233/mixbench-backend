/*
 * @Author: licat
 * @Date: 2023-02-17 17:24:28
 * @LastEditors: licat
 * @LastEditTime: 2023-02-17 18:11:28
 * @Description: licat233@gmail.com
 */
package svc

import (
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	RedisCli          *redis.Redis
	AdminerModel      model.AdminerModel
	JwtBlacklistModel model.JwtBlacklistModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,
		RedisCli: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
		AdminerModel:      model.NewAdminerModel(sqlConn, c.CacheRedis),
		JwtBlacklistModel: model.NewJwtBlacklistModel(sqlConn, c.CacheRedis),
	}
}
