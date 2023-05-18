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
	resp, err := l.svcCtx.LuckydrawRpc.GetUserActivities(l.ctx, &luckydraw_pb.GetUserActivitiesReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	data := []*types.UserActivity{}
	for _, v := range resp.UserActivities {
		item := &types.UserActivity{
			Activity: *dataconv.PbActivityToApiActivity(v.Activity),
			UserPool: *dataconv.PbUserPoolToApiUserPool(v.UserPool),
		}
		data = append(data, item)
	}

	return respx.NewData(data), nil
}
