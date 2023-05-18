package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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

func (l *ProjectDetailsLogic) ProjectDetails(req *types.GetProjectListReq) (any, error) {
	resp, err := l.svcCtx.LuckydrawRpc.GetProjectDetails(l.ctx, &luckydraw_pb.GetProjectDetailsReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Project: &luckydraw_pb.Project{
			Name:   req.Name,
			Status: req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	data := []*types.ProjectDtail{}
	for _, v := range resp.ProjectDetails {
		data = append(data, PbProjectDetailToApiProjectDetail(v))
	}
	return respx.ListResp("", data, resp.Total, req.PageSize, req.Page), nil
}
