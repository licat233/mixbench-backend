package handler

import (
	"mixbench/service/v1/captcha/api/internal/logic"
	"mixbench/service/v1/captcha/api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func showCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewShowCaptchaLogic(r.Context(), svcCtx, w)
		err := l.ShowCaptcha()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
