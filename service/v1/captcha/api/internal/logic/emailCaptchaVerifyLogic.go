package logic

import (
	"context"
	"mixbench/service/v1/captcha/api/internal/svc"
	"mixbench/service/v1/captcha/api/internal/types"
	"mixbench/service/v1/captcha/rpc/captcha_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCaptchaVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCaptchaVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCaptchaVerifyLogic {
	return &EmailCaptchaVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCaptchaVerifyLogic) EmailCaptchaVerify(req *types.EmailCaptchaVerifyReq) (any, error) {
	if !govalidator.IsEmail(req.Email) {
		return nil, errorx.ErrInvalidEmail
	}

	in := &captcha_pb.AuthEmailCaptchaReq{
		Email:    req.Email,
		Solution: req.Solution,
	}

	res, err := l.svcCtx.CaptchaRpc.AuthEmailCaptcha(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.SingleResp("请求成功", res.Success), nil
}
