package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UsersPoolInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUsersPoolInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UsersPoolInfoLogic {
	return &UsersPoolInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UsersPoolInfoLogic) UsersPoolInfo(req *types.UsersPoolGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetUsersPool(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.UsersPool, err)
}
