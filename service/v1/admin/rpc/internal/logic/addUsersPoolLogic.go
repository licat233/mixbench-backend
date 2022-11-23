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

type AddUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUsersPoolLogic {
	return &AddUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUsersPoolLogic) AddUsersPool(in *admin_pb.UsersPool) (*admin_pb.AddUsersPoolResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	data := &model.UsersPool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		UserId:     in.UserId,
		Status:     in.Status,
		Score:      in.Score,
	}

	result, err := l.svcCtx.UsersPoolModel.Insert(l.ctx, data)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	usersPool, err := l.svcCtx.UsersPoolModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddUsersPoolResp{
		UsersPool: mdUsersPoolToPbUsersPool(usersPool),
	}, nil
}
