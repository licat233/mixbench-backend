package project

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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
	resp, err := l.svcCtx.LuckydrawRpc.GetProjectDetail(l.ctx, &luckydraw_pb.GetProjectDetailReq{
		ProjectId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	data := PbProjectDetailToApiProjectDetail(resp.ProjectDetail)

	return respx.NewData(data), nil
}
