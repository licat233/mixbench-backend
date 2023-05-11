package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyListLogic {
	return &GetLuckyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLuckyListLogic) GetLuckyList(req *types.GetLuckyListReq) (any, error) {
	in := &luckydraw_pb.GetLuckyListReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Lucky: &luckydraw_pb.Lucky{
			Id:          req.Id,
			Uuid:        req.Uuid,
			UserId:      req.UserId,
			PrizePoolId: req.PrizePoolId,
			Status:      req.Status,
			LuckDate:    req.LuckDate,
			Ip:          req.Ip,
		},
	}
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetLuckyList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Luckies
	data := dataconv.PbLuckies2ApiLuckies(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
