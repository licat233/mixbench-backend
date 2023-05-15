package logic

import (
	"context"
	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/captcha/api/internal/svc"
	"mixbench/services/v1/captcha/api/internal/types"
	"mixbench/services/v1/captcha/rpc/captcha_pb"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type MathCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *MathCaptchaLogic {
	return &MathCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *MathCaptchaLogic) MathCaptcha(req *types.CaptchaReq) (any, error) {
	ok, err := RequestCountValid(l.ctx, l.svcCtx.RedisClient, l.r)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	if !ok {
		return nil, errorx.ErrTooManyRequests
	}

	in := &captcha_pb.GetCaptchaReq{
		Height: int32(req.Height),
		Width:  int32(req.Width),
		BgColor: &captcha_pb.CaptchaBgColor{
			R: uint32(req.BgColor.R),
			G: uint32(req.BgColor.G),
			B: uint32(req.BgColor.B),
			A: uint32(req.BgColor.A),
		},
	}
	resp, err := l.svcCtx.CaptchaRpc.GetMathCaptcha(l.ctx, in)
	if err != nil {
		return nil, err
	}
	data := &types.CaptchaInfo{
		CaptchaId:   resp.CaptchaId,
		Base64Image: resp.Base64Image,
	}
	return respx.DefaultSingleResp(data, nil)
}
