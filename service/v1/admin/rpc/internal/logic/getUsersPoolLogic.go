package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersPoolLogic {
	return &GetUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUsersPoolLogic) GetUsersPool(in *admin_pb.ItemReq) (*admin_pb.GetUsersPoolResp, error) {
	res, err := l.svcCtx.UsersPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}
	return &admin_pb.GetUsersPoolResp{
		UsersPool: mdUsersPoolToPbUsersPool(res),
	}, nil
}
