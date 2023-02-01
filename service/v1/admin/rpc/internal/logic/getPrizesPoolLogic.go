package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolLogic {
	return &GetPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizePoolLogic) GetPrizePool(in *admin_pb.ItemReq) (*admin_pb.GetPrizePoolResp, error) {
	res, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &admin_pb.GetPrizePoolResp{
		PrizePool: mdPrizePoolToPbPrizePool(res),
	}, nil
}
