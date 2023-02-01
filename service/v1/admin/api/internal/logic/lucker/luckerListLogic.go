package lucky

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

type LuckyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyListLogic {
	return &LuckyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckyListLogic) LuckyList(req *types.LuckyListReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.GetLuckyList(l.ctx, &admin_pb.LuckyListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Lucky: &admin_pb.Lucky{
			UserId:      req.UserId,
			PrizePoolId: req.PrizePoolId,
			Status:      req.Status,
			LuckDate:    req.LuckDate,
			Ip:          req.Ip,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
