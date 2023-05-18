package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectActivitiesLogic {
	return &ProjectActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectActivitiesLogic) ProjectActivities(req *types.ProjectActivitiesReq) (any, error) {
	resp, err := l.svcCtx.LuckydrawRpc.GetProjectActivities(l.ctx, &luckydraw_pb.GetProjectActivitiesReq{
		ProjectId: req.ProjectId,
	})
	if err != nil {
		return nil, err
	}
	data := dataconv.PbActivities2ApiActivities(resp.Activities)
	n := int64(len(data))
	return respx.ListResp("", data, n, n, 1), nil
}
