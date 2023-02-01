/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:39:06
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

type BasePutProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBasePutProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BasePutProjectLogic {
	return &BasePutProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BasePutProjectLogic) BasePutProject(in *luckydraw_pb.PutProjectReq) (*luckydraw_pb.PutProjectResp, error) {
	data := &model.Project{
		Id:     in.Id,
		Uuid:   in.Uuid,
		Name:   in.Name,
		Status: in.Status,
	}

	if err := l.svcCtx.ProjectModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutProjectResp{}, nil
}
