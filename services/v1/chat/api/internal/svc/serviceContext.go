package svc

import (
	"mixbench/services/v1/captcha/rpc/captcha"
	"mixbench/services/v1/chat/api/internal/client"
	"mixbench/services/v1/chat/api/internal/config"
	"mixbench/services/v1/chat/api/internal/middleware"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	VisitorAuth rest.Middleware

	CaptchaRpc captcha.Captcha

	ChatRecordModel    model.ChatRecordModel
	IpblackModel       model.IpblackModel
	KefuModel          model.KefuModel
	MessageModel       model.MessageModel
	MessageConfigModel model.MessageConfigModel
	VisitorModel       model.VisitorModel

	Hub *client.Hub
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	svc := &ServiceContext{
		Config: c,

		VisitorAuth:        middleware.NewVisitorAuthMiddleware().Handle,
		CaptchaRpc:         captcha.NewCaptcha(zrpc.MustNewClient(c.CaptchaRpc)),
		ChatRecordModel:    model.NewChatRecordModel(sqlConn, c.CacheRedis),
		IpblackModel:       model.NewIpblackModel(sqlConn, c.CacheRedis),
		KefuModel:          model.NewKefuModel(sqlConn, c.CacheRedis),
		MessageModel:       model.NewMessageModel(sqlConn, c.CacheRedis),
		MessageConfigModel: model.NewMessageConfigModel(sqlConn, c.CacheRedis),
		VisitorModel:       model.NewVisitorModel(sqlConn, c.CacheRedis),

		Hub: nil,
	}
	svc.Hub = client.NewHub(svc.MessageModel, svc.ChatRecordModel)
	return svc
}
