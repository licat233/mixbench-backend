/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:24:40
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddPrizePoolLogic {
	return &BaseAddPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseAddPrizePoolLogic) BaseAddPrizePool(in *luckydraw_pb.AddPrizePoolReq) (*luckydraw_pb.AddPrizePoolResp, error) {
	data := &model.PrizePool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		GradeId:    in.GradeId,
		LuckValue:  in.LuckValue,
		Status:     in.Status,
		Total:      in.Total,
		Remainder:  in.Remainder,
	}
	result, err := l.svcCtx.PrizePoolModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	prizesPool, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddPrizePoolResp{
		PrizePool: tools.MdPrizePoolToPbPrizePool(prizesPool),
	}, nil
}
