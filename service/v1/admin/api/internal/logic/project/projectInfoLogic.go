package project

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectInfoLogic {
	return &ProjectInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectInfoLogic) ProjectInfo(req *types.ProjectGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetProject(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Project, err)
}
