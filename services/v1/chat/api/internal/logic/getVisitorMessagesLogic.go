package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVisitorMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetVisitorMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVisitorMessagesLogic {
	return &GetVisitorMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetVisitorMessagesLogic) GetVisitorMessages(req *types.GetVisitorMessagesReq) (any, error) {
	if req.VisitorId < 1 {
		return nil, errorx.ExternalError("visitor id < 1")
	}
	mdMessages, err := l.svcCtx.MessageModel.FindsByVisitorId(l.ctx, req.VisitorId)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessages2ApiMessages(mdMessages)
	return respx.DefaultSingleResp(data, nil)
}
