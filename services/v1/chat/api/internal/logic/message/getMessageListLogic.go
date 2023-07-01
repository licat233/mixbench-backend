package message

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageListLogic) GetMessageList(req *types.GetMessageListReq) (any, error) {
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
	mdList, total, err := l.svcCtx.MessageModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("failed to query message list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdMessages2ApiMessages(mdList)

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
