package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectActivityLogic {
	return &ProjectActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectActivityLogic) ProjectActivity(req *types.ProjectActivityReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.ProjectActivity(l.ctx, &admin_pb.ProjectActivityReq{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	n := int64(len(resp.Data))
	return respx.ListResp("", resp.Data, n, n, 1), nil
}
