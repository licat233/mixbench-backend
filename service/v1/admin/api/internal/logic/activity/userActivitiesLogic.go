package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserActivitiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserActivitiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserActivitiesLogic {
	return &UserActivitiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserActivitiesLogic) UserActivities(req *types.UserActivitiesReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.UserActivity(l.ctx, &admin_pb.UserActivityReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return respx.NewData(resp.Data), nil
}
