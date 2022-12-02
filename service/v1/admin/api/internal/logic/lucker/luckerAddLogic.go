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

type LuckerAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerAddLogic {
	return &LuckerAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckerAddLogic) LuckerAdd(req *types.LuckerAddReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.AddLucker(l.ctx, &admin_pb.Lucker{
		UserId:       req.UserId,
		PrizesPoolId: req.PrizesPoolId,
		Status:       req.Status,
		LuckDate:     req.LuckDate,
		Ip:           req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Lucker, err)
}
