/*
 * @Author: licat
 * @Date: 2023-02-17 17:24:28
 * @LastEditors: licat
 * @LastEditTime: 2023-02-17 18:11:28
 * @Description: licat233@gmail.com
 */
package svc

import (
	"fmt"
	"log"
	"mixbench/common/mysqlx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config        config.Config
	RedisCli      *redis.Redis
	AdminerModel  model.AdminerModel
	ProjectModel  model.ProjectModel
	CampaignModel model.CampaignModel
	WebModel      model.WebModel
	CompanyModel  model.CompanyModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := mysqlx.NewMysql(c.Mysql.DataSource)
	if err != nil {
		log.Fatalln(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	return &ServiceContext{
		Config: c,
		RedisCli: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
		AdminerModel:  model.NewAdminerModel(sqlConn, c.CacheRedis),
		ProjectModel:  model.NewProjectModel(sqlConn, c.CacheRedis),
		CampaignModel: model.NewCampaignModel(sqlConn, c.CacheRedis),
		WebModel:      model.NewWebModel(sqlConn, c.CacheRedis),
		CompanyModel:  model.NewCompanyModel(sqlConn, c.CacheRedis),
	}
}
