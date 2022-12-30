package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityDetailsLogic {
	return &GetActivityDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActivityDetailsLogic) GetActivityDetails(in *admin_pb.ActivityListReq) (*admin_pb.GetActivityDetailsListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.ActivityModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbActivityToMdActivity(in.Activity))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*admin_pb.ActivityDetails{}
	activityPrizeLogic := NewActivityPrizesLogic(l.ctx, l.svcCtx)
	for _, activity := range list {
		resp, e := activityPrizeLogic.ActivityPrizes(&admin_pb.ActivityPrizesReq{
			ActivityId: activity.Id,
		})
		if e != nil {
			return nil, e
		}
		data = append(data, &admin_pb.ActivityDetails{
			Activity:       mdActivityToPbActivity(activity),
			ActivityPrizes: resp.Data,
		})
	}

	return &admin_pb.GetActivityDetailsListResp{
		Data:  data,
		Total: total,
	}, nil
}
