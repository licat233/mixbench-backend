package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPatternEnumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityPatternEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPatternEnumListLogic {
	return &ActivityPatternEnumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityPatternEnumListLogic) ActivityPatternEnumList() (any, error) {
	resp, err := l.svcCtx.LuckydrawRpc.GetActivityPatternEnums(l.ctx, &luckydraw_pb.NilReq{})
	if err != nil {
		return nil, err
	}
	data := dataconv.PbEnumsToApiEnums(resp.Enums)
	return respx.DefaultSingleResp(data, err)
}
