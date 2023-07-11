/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-02 10:23:23
 * @Description: licat233@gmail.com
 */
package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mixbench/common/emailx"
	"mixbench/services/v1/authentication/mq/internal/svc"
	"mixbench/services/v1/authentication/mq/types"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

/**
 * consumer 从redis中取出task进行处理
 */
type AsynqConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Server *asynq.Server
	logx.Logger
}

func NewAsynqConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqConsumer {
	return &AsynqConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Server: asynq.NewServer(svcCtx.RedisClientOpt, asynq.Config{Concurrency: 10}),
		Logger: logx.WithContext(ctx),
	}
}

var _ service.Service = &AsynqConsumer{}

func (p *AsynqConsumer) Start() {
	if err := p.Server.Run(asynq.HandlerFunc(p.handler)); err != nil {
		log.Fatal(err)
	}
}

func (p *AsynqConsumer) Stop() {
	p.Server.Stop()
}

func (p *AsynqConsumer) handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case types.SendEmailCaptchaKey:
		var req types.SendEmailCaptchaReq
		if err := json.Unmarshal(t.Payload(), &req); err != nil {
			p.Logger.Error(err)
			return err
		}
		if err := p.SendEmailCaptchaLogic(&req); err != nil {
			p.Logger.Error(err)
			return err
		}
	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}

func (p *AsynqConsumer) SendEmailCaptchaLogic(req *types.SendEmailCaptchaReq) error {
	emails := req.EmailAddrees
	message := req.Message
	key := req.RedisKey
	err := emailx.SendCaptcha(p.ctx, "mixbench平台验证码", &message, emails)
	if err != nil {
		p.Logger.Error(err)
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, err = p.svcCtx.RedisClient.DelCtx(p.ctx, key); err != nil {
			p.Logger.Error(err)
			return err
		}
		return err
	}
	return err
}
