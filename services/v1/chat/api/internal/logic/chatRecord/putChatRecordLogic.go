package chatRecord

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

type PutChatRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutChatRecordLogic {
	return &PutChatRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutChatRecordLogic) PutChatRecord(req *types.PutChatRecordReq) (any, error) {
	in := &model.ChatRecord{
		Id:          req.Id,
		Uuid:        req.Uuid,
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
	if err := l.svcCtx.ChatRecordModel.Update(l.ctx, in); err != nil {
		l.Logger.Error("failed to update chatRecord, error: ", err)
		return nil, errorx.InternalError(err)
	}

	return respx.DefaultStatusResp(nil)
}
