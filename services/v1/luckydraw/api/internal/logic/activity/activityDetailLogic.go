package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

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
	resp, err := l.svcCtx.LuckydrawRpc.GetActivityDetail(l.ctx, &luckydraw_pb.GetActivityDetailReq{
		ActivityId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	data := resp.ActivityDetail
	res := PbActivityDetailToApiActivityDetail(data)
	return respx.SingleResp("", res), nil
}
