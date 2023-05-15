/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 12:06:13
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"strings"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/captcha/rpc/captcha_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginEmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginEmailCaptchaLogic {
	return &LoginEmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginEmailCaptchaLogic) LoginEmailCaptcha(req *types.LoginEmailCaptchaReq) (any, error) {
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
	resp, err := l.svcCtx.AdminRpc.ExistsEmail(l.ctx, &admin_pb.ExistsEmailReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	if !resp.Exists {
		return nil, errorx.Alert("This email is not registered")
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
