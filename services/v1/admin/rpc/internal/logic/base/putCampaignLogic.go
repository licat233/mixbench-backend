package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutCampaignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutCampaignLogic {
	return &PutCampaignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新活动
func (l *PutCampaignLogic) PutCampaign(in *admin_pb.PutCampaignReq) (*admin_pb.PutCampaignResp, error) {
	data := &model.Campaign{
		Id:        in.Id,
		Uuid:      in.Uuid,
		Name:      in.Name,
		ProjectId: in.ProjectId,
		Status:    in.Status,
		Pattern:   in.Pattern,
		MaxCount:  in.MaxCount,
		MaxPrizes: in.MaxPrizes,
		StartTime: time.Unix(in.StartTime, 0).Local(),
		EndTime:   time.Unix(in.EndTime, 0).Local(),
	}

	if err := l.svcCtx.CampaignModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.PutCampaignResp{}, nil
}
