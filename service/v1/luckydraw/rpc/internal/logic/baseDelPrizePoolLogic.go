/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:26:59
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseDelPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseDelPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseDelPrizePoolLogic {
	return &BaseDelPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseDelPrizePoolLogic) BaseDelPrizePool(in *luckydraw_pb.DelPrizePoolReq) (*luckydraw_pb.DelPrizePoolResp, error) {
	data, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("prizesPool not found"))
	}

	err = l.svcCtx.PrizePoolModel.DeleteSoft(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.DelPrizePoolResp{}, nil
}
