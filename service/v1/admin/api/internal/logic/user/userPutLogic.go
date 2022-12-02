package user

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPutLogic {
	return &UserPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPutLogic) UserPut(req *types.UserPutReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutUser(l.ctx, &admin_pb.User{
		Id:       req.Id,
		Uuid:     req.Uuid,
		Passport: req.Passport,
		Email:    req.Email,
		Phone:    req.Phone,
		Name:     req.Name,
		Address:  req.Address,
		Score:    req.Score,
		Status:   req.Status,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
