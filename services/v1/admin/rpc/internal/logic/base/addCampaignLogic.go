package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCampaignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCampaignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCampaignLogic {
	return &AddCampaignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加活动
func (l *AddCampaignLogic) AddCampaign(in *admin_pb.AddCampaignReq) (*admin_pb.AddCampaignResp, error) {
	data := &model.Campaign{
		Uuid:      uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Name:      in.Name,
		ProjectId: in.ProjectId,
		Status:    in.Status,
		Pattern:   in.Pattern,
		MaxCount:  in.MaxCount,
		MaxPrizes: in.MaxPrizes,
		StartTime: time.Unix(in.StartTime, 0).Local(),
		EndTime:   time.Unix(in.EndTime, 0).Local(),
	}
	if _, err := l.svcCtx.CampaignModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddCampaignResp{}, nil
}
