package logic

import (
	"context"
	"mixbench/service/v1/captcha/api/internal/svc"
	"mixbench/service/v1/captcha/api/internal/types"
	"mixbench/service/v1/captcha/rpc/captcha_pb"
	"net/http"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *EmailCaptchaLogic {
	return &EmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *EmailCaptchaLogic) EmailCaptcha(req *types.EmailCaptchaReq) (any, error) {
	email := strings.TrimSpace(req.Email)
	if req.CaptchaId == "" || req.Solution == "" || email == "" {
		return nil, errorx.ErrInvalidParam
	}

	if !govalidator.IsEmail(email) {
		return nil, errorx.ErrInvalidEmail
	}

	ok, err := RequestCountValid(l.ctx, l.svcCtx.RedisClient, l.r)
	if err != nil {
		logx.Error(err)
		return nil, errorx.InternalError(err)
	}
	if !ok {
		return nil, errorx.ErrTooManyRequests
	}

	//校验图片验证码是否正确
	in := &captcha_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	}
	res, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, in)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errorx.Alert("Captcha verification failed")
	}

	resp, err := l.svcCtx.CaptchaRpc.GetEmailCaptcha(l.ctx, &captcha_pb.GetEmailCaptchaReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	data := &types.EmailCaptchaResp{
		CaptchaId: resp.CaptchaId,
		ExpiresAt: resp.ExpiresAt,
	}
	return respx.NewData(data), nil
}
