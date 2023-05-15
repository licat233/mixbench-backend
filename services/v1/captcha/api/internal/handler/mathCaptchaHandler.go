package handler

import (
	"mixbench/services/v1/captcha/api/internal/logic"
	"mixbench/services/v1/captcha/api/internal/svc"
	"mixbench/services/v1/captcha/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func mathCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMathCaptchaLogic(r.Context(), svcCtx, r)
		resp, err := l.MathCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
