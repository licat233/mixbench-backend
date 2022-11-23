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

	LuckerModel     model.LuckerModel
	PrizesPoolModel model.PrizesPoolModel
	UsersPoolModel  model.UsersPoolModel
	PrizeGradeModel model.PrizeGradeModel

	CommonModel *model.CommonModel

	AdminerModel       model.AdminerModel
	JwtBlacklistsModel model.JwtBlacklistsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	svc := &ServiceContext{
		Config: c,
		RedisCli: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Pass = c.Redis.Pass
			r.Type = c.Redis.Type
		}),
		ProjectModel:       model.NewProjectModel(sqlConn, c.CacheRedis),
		ActivityModel:      model.NewActivityModel(sqlConn, c.CacheRedis),
		PrizeModel:         model.NewPrizeModel(sqlConn, c.CacheRedis),
		UserModel:          model.NewUserModel(sqlConn, c.CacheRedis),
		LuckerModel:        model.NewLuckerModel(sqlConn, c.CacheRedis),
		PrizesPoolModel:    model.NewPrizesPoolModel(sqlConn, c.CacheRedis),
		UsersPoolModel:     model.NewUsersPoolModel(sqlConn, c.CacheRedis),
		PrizeGradeModel:    model.NewPrizeGradeModel(sqlConn, c.CacheRedis),
		AdminerModel:       model.NewAdminerModel(sqlConn, c.CacheRedis),
		JwtBlacklistsModel: model.NewJwtBlacklistsModel(sqlConn, c.CacheRedis),
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
