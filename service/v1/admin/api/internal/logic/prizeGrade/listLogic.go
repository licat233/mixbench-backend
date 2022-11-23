package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

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

func (l *ListLogic) List(req *types.PrizeGradeListReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetPrizeGradeList(l.ctx, &admin_pb.PrizeGradeListReq{
		ListReq: &admin_pb.ListReq{PageSize: req.PageSize, Page: req.Page, Keyword: req.Keyword},
		PrizeGrade: &admin_pb.PrizeGrade{
			Name:  req.Name,
			IsWin: req.IsWin,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
