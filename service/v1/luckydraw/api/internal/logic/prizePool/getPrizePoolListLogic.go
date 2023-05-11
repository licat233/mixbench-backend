package prizePool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizePoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolListLogic {
	return &GetPrizePoolListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizePoolListLogic) GetPrizePoolList(req *types.GetPrizePoolListReq) (any, error) {
	in := &luckydraw_pb.GetPrizePoolListReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		PrizePool: &luckydraw_pb.PrizePool{
			Id:         req.Id,
			Uuid:       req.Uuid,
			ActivityId: req.ActivityId,
			PrizeId:    req.PrizeId,
			GradeId:    req.GradeId,
			LuckValue:  req.LuckValue,
			Status:     req.Status,
			Total:      req.Total,
			Remainder:  req.Remainder,
		},
	}
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrizePoolList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.PrizePools
	data := dataconv.PbPrizePools2ApiPrizePools(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
