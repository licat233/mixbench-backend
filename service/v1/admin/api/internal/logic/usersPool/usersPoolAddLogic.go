package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersPoolAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersPoolAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersPoolAddLogic {
	return &UsersPoolAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersPoolAddLogic) UsersPoolAdd(req *types.UsersPoolAddReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.AddUsersPool(l.ctx, &admin_pb.UsersPool{
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Status:     req.Status,
		Score:      req.Score,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.UsersPool, err)
}
