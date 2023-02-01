package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolLogic {
	return &GetUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserPoolLogic) GetUserPool(in *admin_pb.ItemReq) (*admin_pb.GetUserPoolResp, error) {
	res, err := l.svcCtx.UserPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}
	return &admin_pb.GetUserPoolResp{
		UserPool: mdUserPoolToPbUserPool(res),
	}, nil
}
