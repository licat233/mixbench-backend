package logic

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"mixbench/common/emailx"
	"mixbench/common/errorx"
	"mixbench/services/v1/captcha/rpc/captcha_pb"
	"mixbench/services/v1/captcha/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
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

// 获取邮箱验证码
func (l *SendEmailCaptchaLogic) SendEmailCaptcha(in *captcha_pb.SendEmailCaptchaReq) (resp *captcha_pb.SendEmailCaptchaResp, err error) {
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

	//防止某个邮箱过于频繁的获取邮件
	emailKey := fmt.Sprintf("sendCount:%s", email)
	var count int
	value, err := l.svcCtx.RedisClient.GetCtx(l.ctx, emailKey)
	if err != nil && err != redis.Nil {
		return nil, errorx.IntRpcErr(err)
	}
	if err == nil {
		count, err = strconv.Atoi(value)
		if err != nil {
			if _, err := l.svcCtx.RedisClient.DelCtx(l.ctx, emailKey); err != nil {
				l.Logger.Error(err)
				return nil, errorx.IntRpcErr(err)
			}
		}
	}

	if count > 5 {
		ttl, err := l.svcCtx.RedisClient.TtlCtx(l.ctx, emailKey)
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if ttl > 0 {
			return nil, errorx.ExtRpcErr(fmt.Errorf("email sent too frequently. Please try again in %d seconds", ttl))
		}
	}

	count++

	rdbExp := int(l.svcCtx.CaptchaExpire.Seconds())

	value = strconv.Itoa(count)

	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, emailKey, value, rdbExp); err != nil {
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

	key := formatEmailCaptchaKey(captchaId)

	//先存入redis中，key为邮箱，value为验证码，过期时间为 expireTime
	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, key, email, rdbExp); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	err = emailx.SendCaptcha(l.ctx, "mixbench验证码", &emailx.CaptchaMessage{
		Captcha:  answer,
		Account:  in.Account,
		Datetime: in.Datetime,
		Address:  in.Address,
		IP:       in.Ip,
	}, []string{email})
	if err != nil {
		l.Logger.Error(err)
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, e := l.svcCtx.RedisClient.DelCtx(l.ctx, key); e != nil {
			l.Logger.Error(e)
			return nil, errorx.IntRpcErr(e)
		}
		return nil, errorx.IntRpcErr(err)
	}

	//发送成功
	expiresAt := time.Now().Add(l.svcCtx.CaptchaExpire).Unix()
	return &captcha_pb.SendEmailCaptchaResp{
		CaptchaId: captchaId,
		ExpiresAt: expiresAt,
	}, nil
}
