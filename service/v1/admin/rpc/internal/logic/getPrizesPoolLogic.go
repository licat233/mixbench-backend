package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizesPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizesPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizesPoolLogic {
	return &GetPrizesPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizesPoolLogic) GetPrizesPool(in *admin_pb.ItemReq) (*admin_pb.GetPrizesPoolResp, error) {
	res, err := l.svcCtx.PrizesPoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &admin_pb.GetPrizesPoolResp{
		PrizesPool: mdPrizesPoolToPbPrizesPool(res),
	}, nil
}
