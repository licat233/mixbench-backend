package captcha

import (
	"net/http"

	"mixbench/services/v1/authentication/api/internal/logic/captcha"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyEmailCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmailCaptchaVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := captcha.NewVerifyEmailCaptchaLogic(r.Context(), svcCtx, r)
		resp, err := l.VerifyEmailCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
