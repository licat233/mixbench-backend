package usersPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersPoolPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersPoolPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersPoolPutLogic {
	return &UsersPoolPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersPoolPutLogic) UsersPoolPut(req *types.UsersPoolPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutUsersPool(l.ctx, &admin_pb.UsersPool{
		Id:         req.Id,
		Uuid:       req.Uuid,
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Status:     req.Status,
		Score:      req.Score,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
