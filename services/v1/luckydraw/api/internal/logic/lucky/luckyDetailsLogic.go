package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckyDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyDetailsLogic {
	return &LuckyDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckyDetailsLogic) LuckyDetails(req *types.LuckyDetailsReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return respx.ListResp("", []any{}, 0, req.PageSize, req.Page), nil
	}
	in := &luckydraw_pb.LuckyDetailsReq{
		ListReq: &luckydraw_pb.ListReq{PageSize: req.PageSize, Page: req.Page, Keyword: req.Keyword},
		Lucky: &luckydraw_pb.Lucky{
			Id:          0,
			Uuid:        "",
			UserId:      req.UserId,
			PrizePoolId: req.PrizePoolId,
			Status:      req.Status,
			LuckDate:    req.LuckDate,
			Ip:          req.Ip,
		},
		UserName:     req.UserName,
		ProjectId:    req.ProjectId,
		ActivityId:   req.ActivityId,
		ActivityName: req.ActivityName,
		PrizeId:      req.PeizeId,
		PrizeName:    req.PrizeName,
		PrizeGradeId: req.PeizeGradeId,
	}
	resp, err := l.svcCtx.LuckydrawRpc.GetLuckyDetailsList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.ListResp("", resp.Data, resp.Total, req.PageSize, req.Page), nil
}
