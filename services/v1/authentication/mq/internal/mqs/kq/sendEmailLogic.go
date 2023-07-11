/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-02 10:39:36
 * @Description: licat233@gmail.com
 */
package kq

import (
	"context"
	"encoding/json"
	"mixbench/common/emailx"
	"mixbench/services/v1/authentication/mq/internal/svc"
	"mixbench/services/v1/authentication/mq/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Consume 方法，用于实现kq.ConsumeHandler接口的方法Consume(key, value string) error
func (l *SendEmailLogic) Consume(key, val string) error {
	var req types.SendEmailCaptchaReq
	if err := json.Unmarshal([]byte(val), &req); err != nil {
		l.Logger.Errorf("Consume unmarshal error, err: %+v", err)
		return err
	}
	if err := l.SendEmailCaptchaLogic(&req); err != nil {
		l.Logger.Error(err)
		return err
	}
	// switch key {
	// case types.SendEmailCaptchaKey:
	// 	var req types.SendEmailCaptchaReq
	// 	if err := json.Unmarshal([]byte(val), &req); err != nil {
	// 		l.Logger.Errorf("Consume unmarshal error, err: %+v", err)
	// 		return err
	// 	}
	// 	if err := l.SendEmailCaptchaLogic(&req); err != nil {
	// 		l.Logger.Error(err)
	// 		return err
	// 	}
	// default:
	// 	return fmt.Errorf("unexpected task type: %s", key)
	// }
	return nil
}

func (l *SendEmailLogic) SendEmailCaptchaLogic(req *types.SendEmailCaptchaReq) error {
	emails := req.EmailAddrees
	message := req.Message
	key := req.RedisKey
	err := emailx.SendCaptcha(l.ctx, "mixbench平台验证码", &message, emails)
	if err != nil {
		l.Logger.Error(err)
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, err = l.svcCtx.RedisClient.DelCtx(l.ctx, key); err != nil {
			l.Logger.Error(err)
			return err
		}
		return err
	}
	return err
}
