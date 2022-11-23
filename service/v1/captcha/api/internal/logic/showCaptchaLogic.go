package logic

import (
	"context"
	"encoding/base64"
	"errors"
	"mixbench/service/v1/captcha/api/internal/svc"
	"mixbench/service/v1/captcha/rpc/captcha_pb"
	"net/http"
	"strings"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	w      http.ResponseWriter
}

func NewShowCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter) *ShowCaptchaLogic {
	return &ShowCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		w:      w,
	}
}

func (l *ShowCaptchaLogic) ShowCaptcha() error {
	in := &captcha_pb.GetCaptchaReq{
		Height: 60,
		Width:  200,
		BgColor: &captcha_pb.CaptchaBgColor{
			R: 254,
			G: 254,
			B: 254,
			A: 254,
		},
	}
	captchaResp, err := l.svcCtx.CaptchaRpc.GetMathCaptcha(l.ctx, in)
	if err != nil {
		return err
	}
	b64s := captchaResp.Base64Image
	mark := "data:image/png;base64,"
	m := strings.Index(b64s, mark)
	if m == -1 {
		err = errors.New("showCaptchaLogic.go strings.Index(b64s, mark) not found mark")
		logx.Error(err)
		return errorx.InternalError(err)
	}
	n := m + len(mark)
	base := string([]byte(b64s)[n:])
	p, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		logx.Error(err)
		return errorx.InternalError(err)
	}
	l.w.Header().Set("cache-control", "no-cache")
	l.w.Header().Set("content-type", "image/png")
	l.w.Header().Add("captcha-type", "math")
	l.w.Header().Add("captcha-id", captchaResp.CaptchaId)
	_, err = l.w.Write(p)
	if err != nil {
		logx.Error(err)
		return errorx.InternalError(err)
	}
	return nil
}
