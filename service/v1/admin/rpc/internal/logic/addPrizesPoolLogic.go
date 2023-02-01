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

type AddPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizePoolLogic {
	return &AddPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPrizePoolLogic) AddPrizePool(in *admin_pb.PrizePool) (*admin_pb.AddPrizePoolResp, error) {
	if in == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}

	data := &model.PrizePool{
		Uuid:       uniqueid.NewUUID(),
		ActivityId: in.ActivityId,
		PrizeId:    in.PrizeId,
		GradeId:    in.GradeId,
		LuckValue:  in.LuckValue,
		Status:     in.Status,
		Total:      in.Total,
		Remainder:  in.Remainder,
	}
	result, err := l.svcCtx.PrizePoolModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	prizesPool, err := l.svcCtx.PrizePoolModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.AddPrizePoolResp{
		PrizePool: mdPrizePoolToPbPrizePool(prizesPool),
	}, nil
}
