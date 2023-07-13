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

type GetCampaignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCampaignLogic {
	return &GetCampaignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCampaignLogic) GetCampaign(req *types.GetCampaignReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetCampaign(l.ctx, &admin_pb.GetCampaignReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbCampaign := rpcResp.Campaign
	data := dataconv.PbCampaignToApiCampaign(pbCampaign)

	return respx.DefaultSingleResp(data, nil)
}
