package campaign

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/logic/dataconv"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCampaignListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCampaignListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCampaignListLogic {
	return &GetCampaignListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCampaignListLogic) GetCampaignList(req *types.GetCampaignListReq) (any, error) {
	in := &admin_pb.GetCampaignListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Campaign: &admin_pb.Campaign{
			Id:        req.Id,
			Uuid:      req.Uuid,
			Name:      req.Name,
			ProjectId: req.ProjectId,
			Status:    req.Status,
			Pattern:   req.Pattern,
			MaxCount:  req.MaxCount,
			MaxPrizes: req.MaxPrizes,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		},
	}
	rpcResp, err := l.svcCtx.AdminRpc.GetCampaignList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Campaigns
	data := dataconv.PbCampaigns2ApiCampaigns(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
