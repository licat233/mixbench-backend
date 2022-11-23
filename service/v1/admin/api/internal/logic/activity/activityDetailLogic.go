package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityDetailLogic {
	return &ActivityDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityDetailLogic) ActivityDetail(req *types.ActivityDetailReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetActivityDetail(l.ctx, &admin_pb.GetActivityDetailReq{
		ActivityId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.SingleResp("", resp.Data), nil
}
