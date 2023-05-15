package project

import (
	"net/http"

	"mixbench/services/v1/luckydraw/api/internal/logic/project"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetProjectEnumsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetProjectEnumsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := project.NewGetProjectEnumsLogic(r.Context(), svcCtx)
		resp, err := l.GetProjectEnums(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
