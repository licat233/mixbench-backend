/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:37:16
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

type BaseGetUserPoolByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetUserPoolByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetUserPoolByIdLogic {
	return &BaseGetUserPoolByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetUserPoolByIdLogic) BaseGetUserPoolById(in *luckydraw_pb.GetUserPoolByIdReq) (*luckydraw_pb.GetUserPoolByIdResp, error) {
	res, err := l.svcCtx.UserPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetUserPoolByIdResp{
		UserPool: tools.MdUserPoolToPbUserPool(res),
	}, nil
}
