package adminer

import (
	"net/http"

	"mixbench/services/v1/admin/api/internal/logic/adminer"
	"mixbench/services/v1/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CurrentInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := adminer.NewCurrentInfoLogic(r.Context(), svcCtx)
		resp, err := l.CurrentInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
