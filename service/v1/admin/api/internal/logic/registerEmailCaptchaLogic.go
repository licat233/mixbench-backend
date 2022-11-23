package logic

import (
	"context"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/captcha/rpc/captcha_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterEmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterEmailCaptchaLogic {
	return &RegisterEmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterEmailCaptchaLogic) RegisterEmailCaptcha(req *types.RegisterEmailCaptchaReq) (any, error) {
	if req.Solution == "" || req.CaptchaId == "" || req.Email == "" {
		return nil, errorx.ErrInvalidParam
	}
	email := strings.TrimSpace(req.Email)
	//验证邮箱
	if !govalidator.IsEmail(email) {
		return nil, errorx.ErrInvalidEmail
	}

	//校验图片验证码
	res, err := l.svcCtx.CaptchaRpc.VerifyCaptcha(l.ctx, &captcha_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errorx.Alert("captcha auth failed")
	}

	//检查该邮箱是否存在
	resp, err := l.svcCtx.AdminRpc.ExsitEmail(l.ctx, &admin_pb.ExsitEmailReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	if resp.Exsit {
		return nil, errorx.Alert("This email already exists")
	}

	//开始发送邮件
	result, err := l.svcCtx.CaptchaRpc.GetEmailCaptcha(l.ctx, &captcha_pb.GetEmailCaptchaReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	return respx.NewData(&types.CaptchaResp{
		CaptchaId: result.CaptchaId,
		ExpiresAt: result.ExpiresAt,
	}), nil
}
