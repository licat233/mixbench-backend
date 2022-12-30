package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityDetailLogic {
	return &GetActivityDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityDetailLogic) GetActivityDetail(in *admin_pb.GetActivityDetailReq) (*admin_pb.GetActivityDetailResp, error) {
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("cannot find activity with the specified id"))
	}
	activityPrizeLogic := NewActivityPrizesLogic(l.ctx, l.svcCtx)
	resp, e := activityPrizeLogic.ActivityPrizes(&admin_pb.ActivityPrizesReq{
		ActivityId: activity.Id,
	})
	if e != nil {
		return nil, e
	}

	return &admin_pb.GetActivityDetailResp{
		Data: &admin_pb.ActivityDetails{
			Activity:       mdActivityToPbActivity(activity),
			ActivityPrizes: resp.Data,
		},
	}, nil
}
