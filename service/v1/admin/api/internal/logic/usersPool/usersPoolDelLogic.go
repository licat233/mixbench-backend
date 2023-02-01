package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPoolDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPoolDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPoolDelLogic {
	return &UserPoolDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPoolDelLogic) UserPoolDel(req *types.UserPoolDelReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.DelUserPool(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
