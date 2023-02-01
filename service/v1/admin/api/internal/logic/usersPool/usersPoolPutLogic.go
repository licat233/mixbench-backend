package usersPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPoolPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPoolPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPoolPutLogic {
	return &UserPoolPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPoolPutLogic) UserPoolPut(req *types.UserPoolPutReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.PutUserPool(l.ctx, &admin_pb.UserPool{
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
