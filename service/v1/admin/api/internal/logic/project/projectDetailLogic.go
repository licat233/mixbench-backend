package project

import (
	"context"

	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectDetailLogic {
	return &ProjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectDetailLogic) ProjectDetail(req *types.ProjectDetailReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetProjectDetail(l.ctx, &admin_pb.GetProjectDetailReq{
		ProjectId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}
