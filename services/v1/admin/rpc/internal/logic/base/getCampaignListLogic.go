package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCampaignListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCampaignListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCampaignListLogic {
	return &GetCampaignListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动列表
func (l *GetCampaignListLogic) GetCampaignList(in *admin_pb.GetCampaignListReq) (*admin_pb.GetCampaignListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.CampaignModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbCampaignToMdCampaign(in.Campaign))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetCampaignListResp{
		Campaigns: dataconv.MdCampaigns2PbCampaigns(list),
		Total:     total,
	}, nil
}
