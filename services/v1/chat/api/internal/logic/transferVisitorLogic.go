package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/respx"
	"mixbench/services/v1/chat/api/internal/client"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferVisitorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferVisitorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferVisitorLogic {
	return &TransferVisitorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferVisitorLogic) TransferVisitor(req *types.TransferVisitorReq) (any, error) {
	//获取当前客服id
	// currentKefuId := ctxdata.GetInt64FromCtx(l.ctx, "id")

	visitorId := req.VisitorId
	if visitorId <= 0 {
		return nil, errorx.ExternalError("visitorId <=0", "visitorId must be greater than 0")
	}
	kefuId := req.KefuId
	if kefuId <= 0 {
		return nil, errorx.ExternalError("kefuId <=0", "kefuId must be greater than 0")
	}

	//先判断该游客是否存在
	visitor, err := l.svcCtx.VisitorModel.FindOne(l.ctx, visitorId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	//如果不存在，则返回错误
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError("The visitorId does not exist")
	}

	//先判断该客服是否存在
	_, err = l.svcCtx.KefuModel.FindOne(l.ctx, kefuId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	//如果不存在，则返回错误
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError("The kefuId does not exist")
	}

	//修改该游客的接待客服id
	visitor.KefuId = kefuId
	//更新游客信息
	if err := l.svcCtx.VisitorModel.Update(l.ctx, visitor); err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}

	//同时还要更新 clients 中心的游客信息
	//判断是否还在线
	visitorCli := l.svcCtx.Hub.VisitorClients.Get(visitorId)
	if visitorCli != nil {
		//如果还在线，则修改
		visitorCli.Visitor.KefuId = kefuId
	}
	//如果不在线，就不用修改

	//再更新message缓存的数据
	//查找出游客发出去的信息
	mdMsgList, _, err := l.svcCtx.MessageModel.FindList(l.ctx, -1, -1, "", &model.Message{
		VisitorId: visitorId,
		MsgType:   client.VisitorMsgType,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	// 必须一条一条的更新，才能更新redis缓存
	for _, msg := range mdMsgList {
		msg.KefuId = kefuId
		if err := l.svcCtx.MessageModel.Update(l.ctx, msg); err != nil {
			l.Logger.Error(err)
			return nil, errorx.InternalError(err)
		}
	}

	//更新 chat_record 数据
	mdRecordList, _, err := l.svcCtx.ChatRecordModel.FindList(l.ctx, -1, -1, "", &model.ChatRecord{
		VisitorId: visitorId,
		MsgType:   client.VisitorMsgType,
	})
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	// 必须一条一条的更新，才能更新redis缓存
	for _, record := range mdRecordList {
		record.KefuId = kefuId
		if err := l.svcCtx.ChatRecordModel.Update(l.ctx, record); err != nil {
			l.Logger.Error(err)
			return nil, errorx.InternalError(err)
		}
	}

	//还需要修改 send 通道中的消息
	//需要再 send 前做校验逻辑，代码不在这里写

	return respx.DefaultStatusResp(nil)
}
