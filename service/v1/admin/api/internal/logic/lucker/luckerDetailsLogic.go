package lucker

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckerDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerDetailsLogic {
	return &LuckerDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckerDetailsLogic) LuckerDetails(req *types.LuckerDetailsReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return respx.ListResp("", []any{}, 0, req.PageSize, req.Page), nil
	}
	in := &admin_pb.LuckerDetailsReq{
		ListReq:      &admin_pb.ListReq{PageSize: req.PageSize, Page: req.Page, Keyword: req.Keyword},
		Lucker:       &admin_pb.Lucker{UserId: req.UserId, PrizesPoolId: req.PrizesPoolId, Status: req.Status, LuckDate: req.LuckDate, Ip: req.Ip},
		UserName:     req.UserName,
		ProjectId:    req.ProjectId,
		ActivityId:   req.ActivityId,
		ActivityName: req.ActivityName,
		PrizeId:      req.PeizeId,
		PrizeName:    req.PrizeName,
		PrizeGradeId: req.PeizeGradeId,
	}
	resp, err := l.svcCtx.AdminRpc.GetLuckerDetailsList(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.ListResp("", resp.Data, resp.Total, req.PageSize, req.Page), nil
}
