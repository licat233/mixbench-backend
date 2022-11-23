package upload

import (
	"net/http"

	"mixbench/service/v1/admin/api/internal/logic/upload"
	"mixbench/service/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadLogoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := upload.NewUploadLogoLogic(r.Context(), svcCtx, r)
		resp, err := l.UploadLogo()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
