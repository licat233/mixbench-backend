package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

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
	in := &admin_pb.LuckyDetailsReq{
		ListReq:      &admin_pb.ListReq{PageSize: req.PageSize, Page: req.Page, Keyword: req.Keyword},
		Lucky:        &admin_pb.Lucky{UserId: req.UserId, PrizePoolId: req.PrizePoolId, Status: req.Status, LuckDate: req.LuckDate, Ip: req.Ip},
		UserName:     req.UserName,
		ProjectId:    req.ProjectId,
		ActivityId:   req.ActivityId,
		ActivityName: req.ActivityName,
		PrizeId:      req.PeizeId,
		PrizeName:    req.PrizeName,
		PrizeGradeId: req.PeizeGradeId,
	}
	resp, err := l.svcCtx.AdminRpc.GetLuckyDetailsList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.ListResp("", resp.Data, resp.Total, req.PageSize, req.Page), nil
}
