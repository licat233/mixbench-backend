package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCampaignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCampaignLogic {
	return &DelCampaignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除活动
func (l *DelCampaignLogic) DelCampaign(in *admin_pb.DelCampaignReq) (*admin_pb.DelCampaignResp, error) {
	if err := l.svcCtx.CampaignModel.SoftDelete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.DelCampaignResp{}, nil
}
