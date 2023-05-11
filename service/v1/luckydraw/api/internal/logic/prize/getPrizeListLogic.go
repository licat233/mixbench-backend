package prize

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeListLogic {
	return &GetPrizeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizeListLogic) GetPrizeList(req *types.GetPrizeListReq) (any, error) {
	in := &luckydraw_pb.GetPrizeListReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Prize: &luckydraw_pb.Prize{
			Id:     req.Id,
			Uuid:   req.Uuid,
			Name:   req.Name,
			Image:  req.Image,
			Price:  req.Price,
			Status: req.Status,
		},
	}
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrizeList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Prizes
	data := dataconv.PbPrizes2ApiPrizes(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
