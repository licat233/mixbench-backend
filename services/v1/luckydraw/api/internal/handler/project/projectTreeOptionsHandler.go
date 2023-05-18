package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/luckydraw/api/internal/logic/project"
	"mixbench/services/v1/luckydraw/api/internal/svc"
)

func ProjectTreeOptionsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := project.NewProjectTreeOptionsLogic(r.Context(), svcCtx)
		resp, err := l.ProjectTreeOptions()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
