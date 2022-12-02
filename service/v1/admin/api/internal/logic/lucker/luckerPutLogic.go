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

type LuckerPutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckerPutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckerPutLogic {
	return &LuckerPutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckerPutLogic) LuckerPut(req *types.LuckerPutReq) (any, error) {
	if req.Ip != "" && !govalidator.IsIP(req.Ip) {
		return nil, errorx.ErrInvalidParam
	}
	_, err := l.svcCtx.AdminRpc.PutLucker(l.ctx, &admin_pb.LuckerPutReq{
		Id:           req.Id,
		Uuid:         req.Uuid,
		UserId:       req.Id,
		PrizesPoolId: req.PrizesPoolId,
		Status:       req.Status,
		LuckDate:     req.LuckDate,
		Ip:           req.Ip,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultStateResp(err)
}
