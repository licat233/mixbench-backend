package message

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutMessageLogic {
	return &PutMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutMessageLogic) PutMessage(req *types.PutMessageReq) (any, error) {
	in := &model.Message{
		Id:         req.Id,
		Uuid:       req.Uuid,
		Content:    req.Content,
		VisitorId:  req.VisitorId,
		KefuId:     req.KefuId,
		MsgType:    req.MsgType,
		Status:     req.Status,
		IsRead:     req.IsRead,
		IsSend:     req.IsSend,
		SendTime:   time.Unix(req.SendTime, 0).Local(),
		ExpireTime: time.Unix(req.ExpireTime, 0).Local(),
	}
	if err := l.svcCtx.MessageModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update message, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
