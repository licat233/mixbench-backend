package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPoolInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPoolInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPoolInfoLogic {
	return &UserPoolInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPoolInfoLogic) UserPoolInfo(req *types.UserPoolGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetUserPool(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.UserPool, err)
}
