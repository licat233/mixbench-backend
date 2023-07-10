package luckydrawlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetLuckValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetLuckValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetLuckValueLogic {
	return &ResetLuckValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置奖品幸运值
func (l *ResetLuckValueLogic) ResetLuckValue(in *luckydraw_pb.ResetLuckValueReq) (*luckydraw_pb.StatusResp, error) {
	//先验证活动信息是否正确
	activity, err := l.svcCtx.ActivityModel.FindById(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("not found activity"))
	}
	if activity.Uuid != in.ActivityUuid {
		return nil, errorx.ExtRpcErr(errors.New("params invalid"))
	}
	prizePools, err := l.svcCtx.PrizePoolModel.FindsByActivityIds(l.ctx, []int64{in.ActivityId})
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	for _, prizePool := range prizePools {
		prizePool.LuckValue = 100
		if err = l.svcCtx.PrizePoolModel.Update(l.ctx, prizePool); err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	}
	return &luckydraw_pb.StatusResp{}, nil
}
