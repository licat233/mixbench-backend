package logic

import (
	"context"
	"net/http"

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
	r      *http.Request
}

func NewVisitorMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *VisitorMessagesLogic {
	return &VisitorMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *VisitorMessagesLogic) VisitorMessages() (any, error) {
	id, err := jwtx.GetIdByRequest(l.r, l.svcCtx.Config.VisitorAuthSecret)
	if err != nil {
		return nil, errorx.AccessError(err)
	}
	list, err := l.svcCtx.MessageModel.FindsByVisitorId(l.ctx, id)
	if err != nil {
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessages2ApiMessages(list)
	return respx.DefaultSingleResp(data, nil)
}
