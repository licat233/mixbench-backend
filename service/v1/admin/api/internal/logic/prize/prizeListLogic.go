package prize

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeListLogic {
	return &PrizeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeListLogic) PrizeList(req *types.PrizeListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizeList(l.ctx, &admin_pb.PrizeListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Prize: &admin_pb.Prize{
			Name:   req.Name,
			Price:  req.Price,
			Status: req.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
