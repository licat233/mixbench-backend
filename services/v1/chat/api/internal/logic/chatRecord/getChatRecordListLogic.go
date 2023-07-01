package chatRecord

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

type GetChatRecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetChatRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatRecordListLogic {
	return &GetChatRecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatRecordListLogic) GetChatRecordList(req *types.GetChatRecordListReq) (any, error) {
	in := &model.ChatRecord{
		Id:        req.Id,
		Uuid:      req.Uuid,
		VisitorId: req.VisitorId,
		KefuId:    req.KefuId,
		Status:    req.Status,
		IsRead:    req.IsRead,
		IsSend:    req.IsSend,
		SendTime:  time.Unix(req.SendTime, 0).Local(),
	}
	mdList, total, err := l.svcCtx.ChatRecordModel.FindList(l.ctx, req.PageSize, req.Page, req.Keyword, in)
	if err != nil {
		l.Logger.Error("failed to query chatRecord list, error: ", err)
		return nil, errorx.InternalError(err)
	}
	data := dataconv.MdChatRecords2ApiChatRecords(mdList)

	return respx.DefaultListResp(data, total, req.PageSize, req.Page, nil)
}
