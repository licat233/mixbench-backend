package ipblack

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIpblackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIpblackLogic {
	return &DelIpblackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelIpblackLogic) DelIpblack(req *types.DelIpblackReq) (any, error) {
	if err := l.svcCtx.IpblackModel.Delete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to delete ipblack, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
