package campaign

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mixbench/services/v1/admin/api/internal/logic/campaign"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
)

func DelCampaignHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelCampaignReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := campaign.NewDelCampaignLogic(r.Context(), svcCtx)
		resp, err := l.DelCampaign(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
