package campaign

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCampaignEnumsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCampaignEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCampaignEnumsLogic {
	return &GetCampaignEnumsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCampaignEnumsLogic) GetCampaignEnums(req *types.GetCampaignEnumsReq) (any, error) {
	rpcResp, err := l.svcCtx.AdminRpc.GetCampaignEnums(l.ctx, &admin_pb.GetCampaignEnumsReq{
		ParentId: req.ParentId,
	})
	if err != nil {
		return nil, err
	}
	data := make([]*types.Enum, 0)
	for _, v := range rpcResp.Enums {
		data = append(data, &types.Enum{
			Label: v.Label,
			Value: v.Value,
		})
	}

	return respx.DefaultSingleResp(data, nil)
}
