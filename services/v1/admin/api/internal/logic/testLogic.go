package logic

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/captcha/rpc/captcha_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *TestLogic {
	return &TestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
		r:      r,
	}
}

func (l *TestLogic) Test(req *types.TestReq) (any, error) {
	ip := utils.RemoteIp(l.r)

	addrDetail := utils.QueryAddress(ip)
	address := fmt.Sprintf("%s-%s-%s", addrDetail.Continent, addrDetail.Country, addrDetail.City)
	resp, err := l.svcCtx.CaptchaRpc.SendEmailCaptcha(l.ctx, &captcha_pb.SendEmailCaptchaReq{
		Email:    "1697470538@qq.com",
		Account:  "licat233",
		Datetime: time.Now().Local().Format("2006-01-02 15:04:05"),
		Address:  address,
		Ip:       ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp, nil)
}
