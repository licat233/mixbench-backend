/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:28:58
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

type BaseGetLuckyByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetLuckyByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetLuckyByIdLogic {
	return &BaseGetLuckyByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetLuckyByIdLogic) BaseGetLuckyById(in *luckydraw_pb.GetLuckyByIdReq) (*luckydraw_pb.GetLuckyByIdResp, error) {
	data, err := l.svcCtx.LuckyModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	// if err == model.ErrNotFound {
	// 	return nil, errorx.ExtRpcErr(errors.New("the lucky is not found"))
	// }
	return &luckydraw_pb.GetLuckyByIdResp{
		Lucky: tools.MdLuckyToPbLucky(data),
	}, nil
}
