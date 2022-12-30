package logic

import (
	"context"
	"fmt"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/captcha/api/internal/svc"
	"mixbench/service/v1/captcha/api/internal/types"
	"mixbench/service/v1/captcha/rpc/captcha_pb"
	"net/http"
	"strings"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/utils"

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
		l.Logger.Error(err)
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

	// 检查邮箱是否已注册
	adminerResp, err := l.svcCtx.AdminRpc.FindByEmail(l.ctx, &admin_pb.FindByEmailReq{
		Email: email,
	})
	if err != nil {
		return nil, err
	}

	adminer := adminerResp.Adminer
	datetime := time.Now().Local().Format("2006-01-02 15:04:05")
	ip := l.r.RemoteAddr

	addrDetail := utils.QueryAddress(ip)
	address := fmt.Sprintf("%s-%s-%s", addrDetail.Continent, addrDetail.Country, addrDetail.City)

	resp, err := l.svcCtx.CaptchaRpc.GetEmailCaptcha(l.ctx, &captcha_pb.GetEmailCaptchaReq{
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
