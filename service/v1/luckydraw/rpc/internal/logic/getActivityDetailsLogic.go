/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 16:27:18
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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

// 获取活动详细列表
func (l *GetActivityDetailsLogic) GetActivityDetails(in *luckydraw_pb.GetActivityDetailsReq) (*luckydraw_pb.GetActivityDetailsResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.ActivityModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbActivityToMdActivity(in.Activity))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*luckydraw_pb.ActivityDetail{}
	activityPrizeLogic := NewGetActivityPrizesLogic(l.ctx, l.svcCtx)
	for _, activity := range list {
		resp, e := activityPrizeLogic.GetActivityPrizes(&luckydraw_pb.GetActivityPrizesReq{
			ActivityId: activity.Id,
		})
		if e != nil {
			return nil, e
		}
		data = append(data, &luckydraw_pb.ActivityDetail{
			Activity:       tools.MdActivityToPbActivity(activity),
			ActivityPrizes: resp.ActivityPrizes,
		})
	}

	return &luckydraw_pb.GetActivityDetailsResp{
		ActivityDetails: data,
		Total:           total,
	}, nil
}
