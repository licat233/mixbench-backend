package lucky

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckyByPrizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutLuckyByPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckyByPrizeLogic {
	return &PutLuckyByPrizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutLuckyByPrizeLogic) PutLuckyByPrize(req *types.LuckyPutByPrizeReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.LuckydrawRpc.PutLuckyByPrize(l.ctx, &luckydraw_pb.PutLuckyByPrizeReq{
		Lucky: &luckydraw_pb.Lucky{
			Id:          req.Id,
			Uuid:        req.Uuid,
			UserId:      req.UserId,
			PrizePoolId: req.PrizeId,
			Status:      req.Status,
			LuckDate:    req.LuckDate,
			Ip:          req.Ip,
		},
		PrizeId:    req.PrizeId,
		ActivityId: req.ActivityId,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStatusResp(err)
}
