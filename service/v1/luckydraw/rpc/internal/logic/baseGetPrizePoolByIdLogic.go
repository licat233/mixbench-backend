/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:33:36
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

type BaseGetPrizePoolByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetPrizePoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetPrizePoolByIdLogic {
	return &BaseGetPrizePoolByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetPrizePoolByIdLogic) BaseGetPrizePoolById(in *luckydraw_pb.GetPrizePoolByIdReq) (*luckydraw_pb.GetPrizePoolByIdResp, error) {
	res, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizePoolByIdResp{
		PrizePool: tools.MdPrizePoolToPbPrizePool(res),
	}, nil
}
