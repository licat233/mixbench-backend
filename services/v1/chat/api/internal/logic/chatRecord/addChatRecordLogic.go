package chatRecord

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

type AddChatRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddChatRecordLogic {
	return &AddChatRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddChatRecordLogic) AddChatRecord(req *types.AddChatRecordReq) (any, error) {
	in := &model.ChatRecord{
		Uuid:        uniqueid.NewUUID(), // 这里的uniqueid包，请自己定义
		Content:     req.Content,
		ContentType: req.ContentType,
		VisitorId:   req.VisitorId,
		KefuId:      req.KefuId,
		ProjectId:   req.ProjectId,
		MsgType:     req.MsgType,
		Status:      req.Status,
		IsRead:      req.IsRead,
		IsSend:      req.IsSend,
		SendTime:    time.Unix(req.SendTime, 0).Local(),
	}
	if _, err := l.svcCtx.ChatRecordModel.Insert(l.ctx, in); err != nil {
		l.Logger.Error("failed to insert chatRecord, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
