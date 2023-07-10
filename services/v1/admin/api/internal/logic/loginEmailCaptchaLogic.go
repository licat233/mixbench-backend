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
	"fmt"
	"net/http"
	"strings"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"
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
	r      *http.Request
}

func NewLoginEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LoginEmailCaptchaLogic {
	return &LoginEmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
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

	// 检查邮箱是否已注册
	adminerResp, err := l.svcCtx.AdminRpc.FindByEmail(l.ctx, &admin_pb.FindByEmailReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}
	if adminerResp.Adminer == nil {
		return nil, errorx.Alert("This email is not registered")
	}

	ip := utils.RemoteIp(l.r)

	addrDetail := utils.QueryAddress(ip)
	address := fmt.Sprintf("%s-%s-%s", addrDetail.Continent, addrDetail.Country, addrDetail.City)

	//开始发送邮件
	result, err := l.svcCtx.CaptchaRpc.SendEmailCaptcha(l.ctx, &captcha_pb.SendEmailCaptchaReq{
		Email:    email,
		Account:  adminerResp.Adminer.Passport,
		Datetime: time.Now().Local().Format("2006-01-02 15:04:05"),
		Address:  address,
		Ip:       ip,
	})
	if err != nil {
		return nil, err
	}

	return respx.NewData(&types.CaptchaResp{
		CaptchaId: result.CaptchaId,
		ExpiresAt: result.ExpiresAt,
	}), nil
}
