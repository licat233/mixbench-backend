package prizesPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetLuckValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetLuckValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetLuckValueLogic {
	return &ResetLuckValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetLuckValueLogic) ResetLuckValue(req *types.ResetLuckValueReq) (any, error) {
	_, err := l.svcCtx.AdminRpc.ResetLuckValue(l.ctx, &admin_pb.ResetLuckValueReq{
		ActivityId:   req.ActivityId,
		ActivityUuid: req.ActivityUuid,
	})
	if err != nil {
		return nil, err
	}
	return respx.StateResp("", true), nil
}
