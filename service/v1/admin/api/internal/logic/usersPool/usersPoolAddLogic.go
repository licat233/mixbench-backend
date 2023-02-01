package usersPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPoolAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPoolAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPoolAddLogic {
	return &UserPoolAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPoolAddLogic) UserPoolAdd(req *types.UserPoolAddReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.AddUserPool(l.ctx, &admin_pb.UserPool{
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Status:     req.Status,
		Score:      req.Score,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.UserPool, err)
}
