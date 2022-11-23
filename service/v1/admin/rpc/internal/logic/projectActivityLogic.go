package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProjectActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProjectActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectActivityLogic {
	return &ProjectActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProjectActivityLogic) ProjectActivity(in *admin_pb.ProjectActivityReq) (*admin_pb.ProjectActivityResp, error) {
	list, err := l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, in.ProjectId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.ProjectActivityResp{
		Data: mdActivityListToPbActivityList(list),
	}, nil
}
