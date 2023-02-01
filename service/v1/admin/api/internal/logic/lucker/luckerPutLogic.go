package lucky

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

type LuckyPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyPutLogic {
	return &LuckyPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckyPutLogic) LuckyPut(req *types.LuckyPutReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutLucky(l.ctx, &admin_pb.LuckyPutReq{
		Id:          req.Id,
		Uuid:        req.Uuid,
		UserId:      req.Id,
		PrizePoolId: req.PrizePoolId,
		Status:      req.Status,
		LuckDate:    req.LuckDate,
		Ip:          req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
