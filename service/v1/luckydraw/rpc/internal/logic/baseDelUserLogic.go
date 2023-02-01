/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:27:23
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

type BaseDelUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseDelUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseDelUserLogic {
	return &BaseDelUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseDelUserLogic) BaseDelUser(in *luckydraw_pb.DelUserReq) (*luckydraw_pb.DelUserResp, error) {
	data, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("user not found"))
	}

	err = l.svcCtx.UserModel.DeleteSoft(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.DelUserResp{}, nil
}
