package logic

import (
	"context"
	"errors"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/redisx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyEmailCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEmailCaptchaLogic {
	return &VerifyEmailCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 校验邮箱验证码
func (l *VerifyEmailCaptchaLogic) VerifyEmailCaptcha(in *authentication_pb.VerifyEmailCaptchaReq) (*authentication_pb.VerifyEmailCaptchaResp, error) {
	//检查邮箱格式
	email := strings.TrimSpace(in.Email)

	//校验邮箱格式是否正确
	if !govalidator.IsEmail(email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}
	//处理captchaId
	captchaId := strings.TrimSpace(in.CaptchaId)
	if len(captchaId) == 0 {
		return nil, errorx.ExtRpcErr(errors.New("captcha_id is nil"))
	}

	//redis中查找该id对应的email
	key := formatCaptchaKey(captchaId)
	emailV, err := l.svcCtx.RedisClient.GetCtx(l.ctx, key)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	if emailV == "" {
		//不存在
		return &authentication_pb.VerifyEmailCaptchaResp{}, nil
	}

	//检查邮箱与captcha_id是否对应
	if emailV != email {
		//不相等，说明该captcha_id不属于该email
		return &authentication_pb.VerifyEmailCaptchaResp{}, nil
	}

	//校验captcha与solution
	solution := strings.TrimSpace(in.Solution)
	ok := l.svcCtx.CaptchaStore.Verify(captchaId, solution, false)
	if !ok {
		//防止硬破解，疯狂尝试，不限制ip是因为考虑局域网
		key = formatEmailCaptchaKey(email, captchaId)
		//2分钟可以有10次数错机会
		isExtErr, err := redisx.KeyValueValidator(l.ctx, l.svcCtx.RedisClient, key, 120, 10)
		if err != nil {
			if isExtErr {
				return nil, errorx.ExtRpcErr(err)
			}
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		return &authentication_pb.VerifyEmailCaptchaResp{}, nil
	}

	//如果校验通过了，则清除验证码
	_ = l.svcCtx.CaptchaStore.Get(captchaId, true)

	return &authentication_pb.VerifyEmailCaptchaResp{
		Success: true,
	}, nil
}
