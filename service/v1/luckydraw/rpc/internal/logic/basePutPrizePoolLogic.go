/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:38:55
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BasePutPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutPrizePoolLogic {
	return &BasePutPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutPrizePoolLogic) BasePutPrizePool(in *luckydraw_pb.PutPrizePoolReq) (*luckydraw_pb.PutPrizePoolResp, error) {
	data := &model.PrizePool{
		Id:         in.Id,
		Uuid:       in.Uuid,
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		GradeId:    in.GradeId,
		LuckValue:  in.LuckValue,
		Status:     in.Status,
		Total:      in.Total,
		Remainder:  in.Remainder,
	}

	if err := l.svcCtx.PrizePoolModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutPrizePoolResp{}, nil
}
