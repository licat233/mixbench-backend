/*
 * @Author: licat
 * @Date: 2023-01-31 17:50:01
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 15:51:49
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	baseluckydrawlogic "mixbench/services/v1/luckydraw/rpc/internal/logic/base"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPrizesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPrizesLogic {
	return &GetActivityPrizesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某活动的奖品
func (l *GetActivityPrizesLogic) GetActivityPrizes(in *luckydraw_pb.GetActivityPrizesReq) (*luckydraw_pb.GetActivityPrizesResp, error) {
	var activity *model.Activity
	var err error
	var activity_id int64 = in.ActivityId
	//获取该活动信息
	if in.ActivityUuid != "" {
		activity, err = l.svcCtx.ActivityModel.FindByUuid(l.ctx, in.ActivityUuid)
		if err != nil && err != model.ErrNotFound {
			return nil, errorx.IntRpcErr(err)
		}
		if err == model.ErrNotFound {
			return &luckydraw_pb.GetActivityPrizesResp{
				ActivityPrizes: nil,
			}, nil
		}
		activity_id = activity.Id
	}

	//获取活动奖品池
	prizesPoolLogic := NewGetActivityPrizePoolLogic(l.ctx, l.svcCtx)
	prizesPoolListResp, err := prizesPoolLogic.GetActivityPrizePool(&luckydraw_pb.GetActivityPrizePoolReq{
		ActivityId: activity_id,
	})
	if err != nil {
		return nil, err
	}
	//提取奖品id
	var prizesIds []int64
	for _, prizesPool := range prizesPoolListResp.PrizePools {
		prizesIds = append(prizesIds, prizesPool.PrizeId)
	}

	//查找奖品
	prizeList, err := l.svcCtx.PrizeModel.FindsByIds(l.ctx, prizesIds)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}

	//获取奖品等级
	prizeGradeListLogic := baseluckydrawlogic.NewGetPrizeGradeListLogic(l.ctx, l.svcCtx)
	prizeGradeListResp, err := prizeGradeListLogic.GetPrizeGradeList(&luckydraw_pb.GetPrizeGradeListReq{
		ListReq:    nil,
		PrizeGrade: nil,
	})
	if err != nil {
		return nil, err
	}

	//整合数据
	var activityPrizes []*luckydraw_pb.ActivityPrize
	for _, prize := range prizeList {
		for _, prizesPool := range prizesPoolListResp.PrizePools {
			if prize.Id == prizesPool.PrizeId {
				for _, prizeGrade := range prizeGradeListResp.PrizeGrades {
					if prizesPool.GradeId == prizeGrade.Id {
						activityPrizes = append(activityPrizes, &luckydraw_pb.ActivityPrize{
							Prize:      tools.MdPrizeToPbPrize(prize),
							PrizePool:  prizesPool,
							PrizeGrade: prizeGrade,
						})
					}
				}
			}
		}
	}

	return &luckydraw_pb.GetActivityPrizesResp{
		ActivityPrizes: activityPrizes,
	}, nil
}
