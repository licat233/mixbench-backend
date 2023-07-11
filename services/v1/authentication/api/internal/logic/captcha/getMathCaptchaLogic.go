package captcha

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMathCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMathCaptchaLogic {
	return &GetMathCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMathCaptchaLogic) GetMathCaptcha(req *types.CaptchaReq) (any, error) {
	in := &authentication_pb.GetCaptchaReq{
		Height: int32(req.Height),
		Width:  int32(req.Width),
		BgColor: &authentication_pb.CaptchaBgColor{
			R: uint32(req.BgColor.R),
			G: uint32(req.BgColor.G),
			B: uint32(req.BgColor.B),
			A: uint32(req.BgColor.A),
		},
	}
	resp, err := l.svcCtx.AuthenticationRpc.GetMathCaptcha(l.ctx, in)
	if err != nil {
		return nil, err
	}
	data := &types.CaptchaInfo{
		CaptchaId:   resp.CaptchaId,
		Base64Image: resp.Base64Image,
	}
	return respx.DefaultSingleResp(data, nil)
}
