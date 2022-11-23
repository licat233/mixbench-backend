package prizesPool

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.PrizesPoolListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizesPoolList(l.ctx, &admin_pb.PrizesPoolListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		PrizesPool: &admin_pb.PrizesPool{
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
