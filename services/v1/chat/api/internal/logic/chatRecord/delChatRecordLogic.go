package chatRecord

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelChatRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelChatRecordLogic {
	return &DelChatRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelChatRecordLogic) DelChatRecord(req *types.DelChatRecordReq) (any, error) {
	if err := l.svcCtx.ChatRecordModel.SoftDelete(l.ctx, req.Id); err != nil {
		l.Logger.Error("failed to soft delete chatRecord, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
