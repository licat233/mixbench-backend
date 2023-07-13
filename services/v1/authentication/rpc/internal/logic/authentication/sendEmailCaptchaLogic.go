package authenticationlogic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"mixbench/common/emailx"
	"mixbench/common/errorx"
	"mixbench/common/redisx"
	"mixbench/services/v1/authentication/mq/asynqmq"
	"mixbench/services/v1/authentication/mq/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailCaptchaLogic {
	return &SendEmailCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送邮箱验证码
func (l *SendEmailCaptchaLogic) SendEmailCaptcha(in *authentication_pb.SendEmailCaptchaReq) (resp *authentication_pb.SendEmailCaptchaResp, err error) {
	defer func() {
		e := recover()
		if e != nil {
			resp = nil
			l.Logger.Error(fmt.Errorf("has panic: %+v", e))
			err = errorx.IntRpcErr(fmt.Errorf("has panic: %+v", e))
		}
	}()
	email := strings.TrimSpace(in.Email)
	//校验邮箱格式是否正确
	if !govalidator.IsEmail(email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}

	ip := in.Ip
	//校验ip是否正确
	if !govalidator.IsIP(ip) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidIp)
	}

	//防止某个ip的疯狂请求行为, 限制10分钟100次
	key := formatIpKey(ip)
	isExtErr, err := redisx.KeyValueValidator(l.ctx, l.svcCtx.RedisClient, key, 600, 100)
	if err != nil {
		if isExtErr {
			return nil, errorx.ExtRpcErr(err)
		}
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//放行ip

	//可能会存在多个不同的ip给一个email发送邮件，所以不应该限制email

	//加以限制，同ip同email，1分钟只能发送1次
	key = formatIpEmailKey(email, ip)
	isExtErr, err = redisx.KeyValueValidator(l.ctx, l.svcCtx.RedisClient, key, 60, 1)
	if err != nil {
		if isExtErr {
			return nil, errorx.ExtRpcErr(err)
		}
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//生成captcha_id
	captchaId, err := generateDigitCaptcha(l.svcCtx.CaptchaStore)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//获取captcha_id的answer
	answer := l.svcCtx.CaptchaStore.Get(captchaId, false)

	key = formatCaptchaKey(captchaId)
	rdbExp := int(l.svcCtx.CaptchaExpire.Seconds())

	//先存入redis中
	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, key, email, rdbExp); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	message := &emailx.CaptchaMessage{
		Captcha:  answer,
		Account:  in.Account,
		Datetime: in.Datetime,
		Address:  in.Address,
		IP:       in.Ip,
	}
	emails := []string{email}
	//避免响应超时，使用消息队列进行发送
	err = l.asynqSendEmailCaptcha(&types.SendEmailCaptchaReq{
		EmailAddrees: emails,
		RedisKey:     key,
		Message:      message,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//发送成功
	expiresAt := time.Now().Add(l.svcCtx.CaptchaExpire).Unix()
	return &authentication_pb.SendEmailCaptchaResp{
		CaptchaId: captchaId,
		ExpiresAt: expiresAt,
	}, nil
}

func (l *SendEmailCaptchaLogic) asynqSendEmailCaptcha(msg *types.SendEmailCaptchaReq) error {
	if l.svcCtx.AsynqClient == nil {
		return errors.New("asynq client未初始化")
	}
	if msg == nil {
		return errors.New("msg is nil")
	}
	payload, err := json.Marshal(*msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}
	task, err := asynqmq.NewSendEmailCaptchaTask(payload)
	if err != nil {
		return err
	}
	return l.svcCtx.AsynqClient.AddTask(task)
}

func (l *SendEmailCaptchaLogic) KqSendEmailCaptcha(msg *types.SendEmailCaptchaReq) error {
	if l.svcCtx.KqPusherClient == nil {
		err := errors.New("kafka client 未初始化")
		return err
	}
	if msg == nil {
		return errors.New("msg is nil")
	}
	body, err := json.Marshal(*msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}
	return l.svcCtx.KqPusherClient.Push(string(body))
}
