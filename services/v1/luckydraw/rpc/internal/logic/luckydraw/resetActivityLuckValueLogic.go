/*
 * @Author: licat
 * @Date: 2023-02-01 18:36:29
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 18:40:51
 * @Description: licat233@gmail.com
 */
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

type ResetActivityLuckValueLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetActivityLuckValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetActivityLuckValueLogic {
	return &ResetActivityLuckValueLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 重置某活动奖品池的奖品运气值，传入一个即可，由于是危险操作，所以需要同时传入 uuid 和 id
func (l *ResetActivityLuckValueLogic) ResetActivityLuckValue(in *luckydraw_pb.ResetActivityLuckValueReq) (*luckydraw_pb.StatusResp, error) {
	//先验证活动信息是否正确
	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("the activity is not found"))
	}
	if activity.Uuid != in.ActivityUuid {
		return nil, errorx.ExtRpcErr(errors.New("params invalid: ID and UUID do not match"))
	}
	prizesPools, _, err := l.svcCtx.PrizePoolModel.FindList(l.ctx, -1, -1, "", &model.PrizePool{
		ActivityId: in.ActivityId,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	for _, prizesPool := range prizesPools {
		prizesPool.LuckValue = 100
		if err = l.svcCtx.PrizePoolModel.Update(l.ctx, prizesPool); err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	}
	return &luckydraw_pb.StatusResp{}, nil
}
