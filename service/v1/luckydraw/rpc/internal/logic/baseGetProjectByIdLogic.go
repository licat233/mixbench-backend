/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:35:07
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

type BaseGetProjectByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetProjectByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetProjectByIdLogic {
	return &BaseGetProjectByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetProjectByIdLogic) BaseGetProjectById(in *luckydraw_pb.GetProjectByIdReq) (*luckydraw_pb.GetProjectByIdResp, error) {
	res, err := l.svcCtx.ProjectModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetProjectByIdResp{
		Project: tools.MdProjectToPbProject(res),
	}, nil
}
