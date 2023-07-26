package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorLastMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVisitorLastMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VisitorLastMessageLogic {
	return &VisitorLastMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VisitorLastMessageLogic) VisitorLastMessage(req *types.VisitorLastMessageReq) (any, error) {
	if req.VisitorId < 1 {
		return nil, errorx.ExternalError("visitor id < 1")
	}
	mdMessage, err := l.svcCtx.MessageModel.FindLastByVisitorId(l.ctx, req.VisitorId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessageToApiMessage(mdMessage)
	return respx.DefaultSingleResp(data, nil)
}
