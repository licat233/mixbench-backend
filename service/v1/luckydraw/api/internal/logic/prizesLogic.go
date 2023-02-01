package logic

import (
	"context"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PrizesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizesLogic {
	return &PrizesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizesLogic) Prizes(req *types.PrizesReq) (resp any, err error) {
	if govalidator.IsUUID(req.ActivityUUID) == false {
		return nil, errorx.ErrInvalidUuid
	}
	in := &luckydraw_pb.ActivePrizesReq{
		ActivityUuid: req.ActivityUUID,
	}
	res, err := l.svcCtx.LuckdrawRpc.GetActivePrizes(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res.Prizes, err)
}
