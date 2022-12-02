package lucker

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LuckerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerListLogic {
	return &LuckerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckerListLogic) LuckerList(req *types.LuckerListReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.GetLuckerList(l.ctx, &admin_pb.LuckerListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Lucker: &admin_pb.Lucker{
			UserId:       req.UserId,
			PrizesPoolId: req.PrizesPoolId,
			Status:       req.Status,
			LuckDate:     req.LuckDate,
			Ip:           req.Ip,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
