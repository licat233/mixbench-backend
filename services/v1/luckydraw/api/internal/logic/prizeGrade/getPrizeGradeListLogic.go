package prizeGrade

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPrizeGradeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeListLogic {
	return &GetPrizeGradeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPrizeGradeListLogic) GetPrizeGradeList(req *types.GetPrizeGradeListReq) (any, error) {
	in := &luckydraw_pb.GetPrizeGradeListReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		PrizeGrade: &luckydraw_pb.PrizeGrade{
			Id:    req.Id,
			Uuid:  req.Uuid,
			Name:  req.Name,
			IsWin: req.IsWin,
		},
	}
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetPrizeGradeList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.PrizeGrades
	data := dataconv.PbPrizeGrades2ApiPrizeGrades(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
