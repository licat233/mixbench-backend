package campaign

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCampaignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutCampaignLogic {
	return &PutCampaignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutCampaignLogic) PutCampaign(req *types.PutCampaignReq) (any, error) {
	in := &admin_pb.PutCampaignReq{
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
	if _, err := l.svcCtx.AdminRpc.PutCampaign(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
