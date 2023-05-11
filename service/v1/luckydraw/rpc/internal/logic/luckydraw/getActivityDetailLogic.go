/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:00:20
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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

// 获取活动详细
func (l *GetActivityDetailLogic) GetActivityDetail(in *luckydraw_pb.GetActivityDetailReq) (*luckydraw_pb.GetActivityDetailResp, error) {
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("cannot find activity with the specified id"))
	}
	activityPrizeLogic := NewGetActivityPrizesLogic(l.ctx, l.svcCtx)
	resp, e := activityPrizeLogic.GetActivityPrizes(&luckydraw_pb.GetActivityPrizesReq{
		ActivityId: activity.Id,
	})
	if e != nil {
		return nil, e
	}

	return &luckydraw_pb.GetActivityDetailResp{
		ActivityDetail: &luckydraw_pb.ActivityDetail{
			Activity:       tools.MdActivityToPbActivity(activity),
			ActivityPrizes: resp.ActivityPrizes,
		},
	}, nil
}
