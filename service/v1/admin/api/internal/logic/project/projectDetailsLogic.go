package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectDetailsLogic {
	return &ProjectDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectDetailsLogic) ProjectDetails(req *types.ProjectListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetProjectDetails(l.ctx, &admin_pb.ProjectListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Project: &admin_pb.Project{
			Name:   req.Name,
			Status: req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.ListResp("", resp.Data, resp.Total, req.PageSize, req.Page), nil
}
