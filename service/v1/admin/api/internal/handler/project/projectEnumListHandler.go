package project

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/service/v1/admin/api/internal/logic/project"
	"mixbench/service/v1/admin/api/internal/svc"
)

func ProjectEnumListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := project.NewProjectEnumListLogic(r.Context(), svcCtx)
		resp, err := l.ProjectEnumList()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
