package prizesPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizePoolListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizePoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizePoolListLogic {
	return &PrizePoolListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizePoolListLogic) PrizePoolList(req *types.PrizePoolListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizePoolList(l.ctx, &admin_pb.PrizePoolListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		PrizePool: &admin_pb.PrizePool{
			ActivityId: req.ActivityId,
			PrizeId:    req.PrizeId,
			GradeId:    req.GradeId,
			LuckValue:  req.LuckValue,
			Status:     req.Status,
			Total:      req.Total,
			Remainder:  req.Remainder,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
