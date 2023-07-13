package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCampaignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCampaignLogic {
	return &GetCampaignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动
func (l *GetCampaignLogic) GetCampaign(in *admin_pb.GetCampaignReq) (*admin_pb.GetCampaignResp, error) {
	res, err := l.svcCtx.CampaignModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &admin_pb.GetCampaignResp{
		Campaign: dataconv.MdCampaignToPbCampaign(res),
	}, nil
}
