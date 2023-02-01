/*
 * @Author: licat
 * @Date: 2023-01-31 17:59:05
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:04:15
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPrizePoolLogic {
	return &GetActivityPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某活动的奖品池
func (l *GetActivityPrizePoolLogic) GetActivityPrizePool(in *luckydraw_pb.GetActivityPrizePoolReq) (*luckydraw_pb.GetActivityPrizePoolResp, error) {
	list, err := l.svcCtx.PrizePoolModel.FindsByActivityId(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetActivityPrizePoolResp{
		PrizePools: tools.MdPrizePoolsToPbPrizePools(list),
	}, nil
}
