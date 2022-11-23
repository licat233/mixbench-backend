package handler

import (
	"mixbench/service/v1/captcha/api/internal/logic"
	"mixbench/service/v1/captcha/api/internal/svc"
	"mixbench/service/v1/captcha/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func processCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProcessCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProcessCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.ProcessCaptcha(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
