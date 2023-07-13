package authenticationlogic

import (
	"context"
	"image/color"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

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

// 获取数学验证码
func (l *GetMathCaptchaLogic) GetMathCaptcha(in *authentication_pb.GetCaptchaReq) (*authentication_pb.GetCaptchaResp, error) {
	captchaConfig := &CaptchaConfig{
		CaptchaType: "math",
		DriverMath: &base64Captcha.DriverMath{
			Height: int(in.Height),
			Width:  int(in.Width),
			// NoiseCount:      10,
			// ShowLineOptions: base64Captcha.OptionShowSlimeLine,
			BgColor: &color.RGBA{R: uint8(in.BgColor.R), G: uint8(in.BgColor.G), B: uint8(in.BgColor.B), A: uint8(in.BgColor.A)},
			Fonts:   []string{"actionj.ttf", "wqy-microhei.ttc", "chromohv.ttf"},
		},
	}

	driver := captchaConfig.DriverMath.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, l.svcCtx.CaptchaStore)
	captchaId, base64Image, err := c.Generate()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GetCaptchaResp{
		CaptchaId:   captchaId,
		Base64Image: base64Image,
	}, nil
}
