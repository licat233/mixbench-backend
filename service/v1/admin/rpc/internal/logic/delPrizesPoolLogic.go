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

type DelPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPrizePoolLogic {
	return &DelPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPrizePoolLogic) DelPrizePool(in *admin_pb.ItemReq) (*admin_pb.StatusResp, error) {
	data, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("prizesPool not found"))
	}

	err = l.svcCtx.PrizePoolModel.DeleteSoft(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return newStatusResponse(), nil
}
