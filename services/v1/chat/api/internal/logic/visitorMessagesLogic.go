package logic

import (
	"context"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVisitorMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VisitorMessagesLogic {
	return &VisitorMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VisitorMessagesLogic) VisitorMessages() (any, error) {
	id := ctxdata.GetInt64FromCtx(l.ctx, jwtx.JwtUidKey)
	list, err := l.svcCtx.MessageModel.FindsByVisitorId(l.ctx, id)
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessages2ApiMessages(list)
	return respx.DefaultSingleResp(data, nil)
}
