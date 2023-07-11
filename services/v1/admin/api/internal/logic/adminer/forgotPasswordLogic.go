package adminer

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type ForgotPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgotPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgotPasswordLogic {
	return &ForgotPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgotPasswordLogic) ForgotPassword(req *types.ForgotPasswordReq) (any, error) {
	//邮箱格式验证
	if !govalidator.IsEmail(req.Email) {
		return nil, errorx.ErrInvalidEmail
	}

	//验证码校验
	resp1, err := l.svcCtx.AuthenticationRpc.VerifyCaptcha(l.ctx, &authentication_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !resp1.Success {
		return nil, errorx.ExternalError(nil, errors.New("验证码错误"))
	}

	//主逻辑
	resp, err := l.svcCtx.AdminRpc.ForgotPassword(l.ctx, &admin_pb.ForgotPasswordReq{
		Email: req.Email,
	})
	if err != nil {
		return nil, err
	}
	return respx.StatusResp("", resp.Status), nil
}
