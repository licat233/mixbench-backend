package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUsersPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUsersPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUsersPoolLogic {
	return &DelUsersPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUsersPoolLogic) DelUsersPool(in *admin_pb.ItemReq) (*admin_pb.StatusResp, error) {
	if in.Id == -1 {
		return nil, errorx.ExtRpcErr(errors.New("无法删除，请设置禁止状态即可"))
	}
	data, err := l.svcCtx.UsersPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("UsersPool not found"))
	}

	err = l.svcCtx.UsersPoolModel.DeleteSoft(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return newStatusResponse(), nil
}
