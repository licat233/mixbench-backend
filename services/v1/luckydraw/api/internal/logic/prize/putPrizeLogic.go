package prize

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizeLogic {
	return &PutPrizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutPrizeLogic) PutPrize(req *types.PutPrizeReq) (any, error) {
	in := &luckydraw_pb.PutPrizeReq{
		Prize: &luckydraw_pb.Prize{
			Id:     req.Id,
			Uuid:   req.Uuid,
			Name:   req.Name,
			Image:  req.Image,
			Price:  req.Price,
			Status: req.Status,
		},
	}
	if _, err := l.svcCtx.LuckydrawRpc.PutPrize(l.ctx, in); err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStatusResp(nil)
}
