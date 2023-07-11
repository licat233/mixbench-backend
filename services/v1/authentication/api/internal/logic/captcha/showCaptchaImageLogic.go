package captcha

import (
	"context"
	"encoding/base64"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowCaptchaImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
	r      *http.Request
}

func NewShowCaptchaImageLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *ShowCaptchaImageLogic {
	return &ShowCaptchaImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowCaptchaImageLogic) ShowCaptchaImage() error {
	in := &authentication_pb.GetCaptchaReq{
		Height: 60,
		Width:  200,
		BgColor: &authentication_pb.CaptchaBgColor{
			R: 254,
			G: 254,
			B: 254,
			A: 254,
		},
	}
	captchaResp, err := l.svcCtx.AuthenticationRpc.GetMathCaptcha(l.ctx, in)
	if err != nil {
		return err
	}
	b64s := captchaResp.Base64Image
	decoded, err := base64.StdEncoding.DecodeString(b64s)
	if err != nil {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}

	body := base64.StdEncoding.EncodeToString(decoded)

	// mark := "data:image/png;base64,"
	// m := strings.Index(b64s, mark)
	// if m == -1 {
	// 	err = errors.New("showCaptchaLogic.go strings.Index(b64s, mark) not found mark")
	// 	l.Logger.Error(err)
	// 	return errorx.InternalError(err)
	// }
	// n := m + len(mark)
	// base := string([]byte(b64s)[n:])
	// p, err := base64.StdEncoding.DecodeString(base)
	// if err != nil {
	// 	l.Logger.Error(err)
	// 	return errorx.InternalError(err)
	// }
	l.w.Header().Set("cache-control", "no-cache")
	l.w.Header().Set("content-type", "image/png")
	l.w.Header().Add("captcha-type", "math")
	l.w.Header().Add("captcha-id", captchaResp.CaptchaId)
	_, err = l.w.Write([]byte(body))
	if err != nil {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}
	return nil
}
