package lucker

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutByPrizeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutByPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutByPrizeLogic {
	return &PutByPrizeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutByPrizeLogic) PutByPrize(req *types.LuckerPutByPrizeReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutLuckerByPrize(l.ctx, &admin_pb.LuckerPutByPrizeReq{
		Id:         req.Id,
		Uuid:       req.Uuid,
		UserId:     req.Id,
		PrizeId:    req.PrizeId,
		ActivityId: req.ActivityId,
		Status:     req.Status,
		LuckDate:   req.LuckDate,
		Ip:         req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
