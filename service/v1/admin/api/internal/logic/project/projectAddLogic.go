package project

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectAddLogic {
	return &ProjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectAddLogic) ProjectAdd(req *types.ProjectAddReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.AddProject(l.ctx, &admin_pb.Project{
		Name:   req.Name,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Project, err)
}
