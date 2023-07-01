package message

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMessageLogic {
	return &DelMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelMessageLogic) DelMessage(req *types.DelMessageReq) (any, error) {
	if err := l.svcCtx.MessageModel.SoftDelete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to soft delete message, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
