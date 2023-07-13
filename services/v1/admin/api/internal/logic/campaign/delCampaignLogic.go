package campaign

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/api/internal/types"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCampaignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCampaignLogic {
	return &DelCampaignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCampaignLogic) DelCampaign(req *types.DelCampaignReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.DelCampaign(l.ctx, &admin_pb.DelCampaignReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
