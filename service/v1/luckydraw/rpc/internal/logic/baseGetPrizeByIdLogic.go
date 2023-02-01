/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:31:12
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

type BaseGetPrizeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetPrizeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetPrizeByIdLogic {
	return &BaseGetPrizeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetPrizeByIdLogic) BaseGetPrizeById(in *luckydraw_pb.GetPrizeByIdReq) (*luckydraw_pb.GetPrizeByIdResp, error) {
	data, err := l.svcCtx.PrizeModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	// if err == model.ErrNotFound {
	// 	return nil, errorx.ExtRpcErr(errors.New("the prize is not found"))
	// }
	return &luckydraw_pb.GetPrizeByIdResp{
		Prize: tools.MdPrizeToPbPrize(data),
	}, nil
}
