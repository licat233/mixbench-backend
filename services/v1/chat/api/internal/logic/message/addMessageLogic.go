package message

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMessageLogic {
	return &AddMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMessageLogic) AddMessage(req *types.AddMessageReq) (any, error) {
	in := &model.Message{
		Uuid:       uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Content:    req.Content,
		VisitorId:  req.VisitorId,
		KefuId:     req.KefuId,
		MsgType:    req.MsgType,
		Status:     req.Status,
		IsRead:     req.IsRead,
		IsSend:     req.IsSend,
		SendTime:   time.Unix(req.SendTime, 0).Local(),
		ExpireTime: time.Now().Local().Add(3 * 24 * time.Hour),
	}
	if _, err := l.svcCtx.MessageModel.Insert(l.ctx, in); err != nil {
		l.Logger.Error("failed to insert message, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
