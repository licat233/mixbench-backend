/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:27:11
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

type BaseDelProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseDelProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseDelProjectLogic {
	return &BaseDelProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseDelProjectLogic) BaseDelProject(in *luckydraw_pb.DelProjectReq) (*luckydraw_pb.DelProjectResp, error) {
	data, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("project not found"))
	}

	err = l.svcCtx.ProjectModel.DeleteSoft(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &luckydraw_pb.DelProjectResp{}, nil
}
