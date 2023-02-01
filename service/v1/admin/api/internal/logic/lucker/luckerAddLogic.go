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

type LuckyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyAddLogic {
	return &LuckyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckyAddLogic) LuckyAdd(req *types.LuckyAddReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.AddLucky(l.ctx, &admin_pb.Lucky{
		UserId:      req.UserId,
		PrizePoolId: req.PrizePoolId,
		Status:      req.Status,
		LuckDate:    req.LuckDate,
		Ip:          req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Lucky, err)
}
