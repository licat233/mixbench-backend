package captcha

import (
	"net/http"

	"mixbench/services/v1/authentication/api/internal/logic/captcha"
	"mixbench/services/v1/authentication/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ShowCaptchaImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := captcha.NewShowCaptchaImageLogic(r.Context(), svcCtx, w, r)
		err := l.ShowCaptchaImage()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
