/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:27:35
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

type BaseDelUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseDelUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseDelUserPoolLogic {
	return &BaseDelUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseDelUserPoolLogic) BaseDelUserPool(in *luckydraw_pb.DelUserPoolReq) (*luckydraw_pb.DelUserPoolResp, error) {
	if in.Id == -1 {
		return nil, errorx.ExtRpcErr(errors.New("无法删除，请设置禁止状态即可"))
	}
	data, err := l.svcCtx.UserPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("UserPool not found"))
	}

	err = l.svcCtx.UserPoolModel.DeleteSoft(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.DelUserPoolResp{}, nil
}
