package logic

import (
	"context"
	"image/color"
	"mixbench/service/v1/captcha/rpc/captcha"
	"mixbench/service/v1/captcha/rpc/captcha_pb"
	"mixbench/service/v1/captcha/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/mojocn/base64Captcha"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMathCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMathCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMathCaptchaLogic {
	return &GetMathCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMathCaptchaLogic) GetMathCaptcha(in *captcha_pb.GetCaptchaReq) (*captcha_pb.GetCaptchaResp, error) {
	captchaConfig := &CaptchaConfig{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			Height:          int(in.Height),
			Width:           int(in.Width),
			NoiseCount:      0,
			ShowLineOptions: 2,
			BgColor:         &color.RGBA{R: uint8(in.BgColor.R), G: uint8(in.BgColor.G), B: uint8(in.BgColor.B), A: uint8(in.BgColor.A)},
			Fonts:           []string{"actionj.ttf", "wqy-microhei.ttc", "chromohv.ttf"},
		},
	}

	driver := captchaConfig.DriverMath.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, l.svcCtx.CaptchaStore)
	captchaId, base64Image, err := c.Generate()
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &captcha.GetCaptchaResp{
		CaptchaId:   captchaId,
		Base64Image: base64Image,
	}, nil
}
