package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserPoolLogic {
	return &AddUserPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserPoolLogic) AddUserPool(in *admin_pb.UserPool) (*admin_pb.AddUserPoolResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	data := &model.UserPool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     in.Status,
		Score:      in.Score,
	}

	result, err := l.svcCtx.UserPoolModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	usersPool, err := l.svcCtx.UserPoolModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddUserPoolResp{
		UserPool: mdUserPoolToPbUserPool(usersPool),
	}, nil
}
