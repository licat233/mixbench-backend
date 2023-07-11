package captcha

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewSendEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *SendEmailCaptchaLogic {
	return &SendEmailCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *SendEmailCaptchaLogic) SendEmailCaptcha(req *types.EmailCaptchaReq) (any, error) {
	email := strings.TrimSpace(req.Email)
	if req.CaptchaId == "" || req.Solution == "" || email == "" {
		return nil, errorx.ErrInvalidParam
	}

	if !govalidator.IsEmail(email) {
		return nil, errorx.ErrInvalidEmail
	}

	//校验图片验证码是否正确
	in := &authentication_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	}
	res, err := l.svcCtx.AuthenticationRpc.VerifyCaptcha(l.ctx, in)
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errorx.Alert("Captcha verification failed")
	}

	// 检查邮箱是否已注册
	adminerResp, err := l.svcCtx.AdminRpc.FindByEmail(l.ctx, &admin_pb.FindByEmailReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	adminer := adminerResp.Adminer
	datetime := time.Now().Local().Format("2006-01-02 15:04:05")
	ip := utils.RemoteIp(l.r)

	addrDetail := utils.QueryAddress(ip)
	address := fmt.Sprintf("%s-%s-%s", addrDetail.Continent, addrDetail.Country, addrDetail.City)

	resp, err := l.svcCtx.AuthenticationRpc.SendEmailCaptcha(l.ctx, &authentication_pb.SendEmailCaptchaReq{
		Email:    email,
		Account:  adminer.Passport,
		Datetime: datetime,
		Address:  address,
		Ip:       ip,
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
