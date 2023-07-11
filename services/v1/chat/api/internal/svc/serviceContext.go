package svc

import (
	"fmt"
	"mixbench/common/utils"
	"mixbench/services/v1/authentication/rpc/authentication"
	"mixbench/services/v1/chat/api/internal/config"
	"mixbench/services/v1/chat/api/internal/middleware"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VisitorAuth rest.Middleware

	AuthenticationRpc authentication.Authentication

	ChatRecordModel    model.ChatRecordModel
	IpblackModel       model.IpblackModel
	KefuModel          model.KefuModel
	MessageModel       model.MessageModel
	MessageConfigModel model.MessageConfigModel
	VisitorModel       model.VisitorModel
	ProjectModel       model.ProjectModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn, err := utils.NewMysql(c.Mysql.DataSource)
	if err != nil {
		panic(fmt.Errorf("连接mysql数据库失败: %s", err))
	}
	svc := &ServiceContext{
		Config:             c,
		VisitorAuth:        middleware.NewVisitorAuthMiddleware().Handle,
		AuthenticationRpc:  authentication.NewAuthentication(zrpc.MustNewClient(c.AuthenticationRpc)),
		ChatRecordModel:    model.NewChatRecordModel(sqlConn, c.CacheRedis),
		IpblackModel:       model.NewIpblackModel(sqlConn, c.CacheRedis),
		KefuModel:          model.NewKefuModel(sqlConn, c.CacheRedis),
		MessageModel:       model.NewMessageModel(sqlConn, c.CacheRedis),
		MessageConfigModel: model.NewMessageConfigModel(sqlConn, c.CacheRedis),
		VisitorModel:       model.NewVisitorModel(sqlConn, c.CacheRedis),
		ProjectModel:       model.NewProjectModel(sqlConn, c.CacheRedis),
	}
	return svc
}
