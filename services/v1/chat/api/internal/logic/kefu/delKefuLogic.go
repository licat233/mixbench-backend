package kefu

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelKefuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelKefuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelKefuLogic {
	return &DelKefuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelKefuLogic) DelKefu(req *types.DelKefuReq) (any, error) {
	if err := l.svcCtx.KefuModel.SoftDelete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to soft delete kefu, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
