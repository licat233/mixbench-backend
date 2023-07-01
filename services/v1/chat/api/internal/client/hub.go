package client

import (
	"context"
	"fmt"
	"mixbench/common/emailx"
	"mixbench/common/timer"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type Hub struct {
	logx.Logger
	ctx context.Context

	Timer  timer.Timer
	StopCh chan bool //监听停止信号的通道

	Upgrader       websocket.Upgrader
	KefuClients    *kefuClientMap
	VisitorClients *visitorClientMap
	MessageList    chan *types.Message

	MessageModel       model.MessageModel
	ChatRecordModel    model.ChatRecordModel
	MessageConfigModel model.MessageConfigModel

	hasSendBeforeHandlerErr bool
	hasSendAfterHandlerErr  bool
	hasDbErr                bool
}

func NewHub(messageModel model.MessageModel, chatRecordModel model.ChatRecordModel, messageConfigModel model.MessageConfigModel) *Hub {
	ctx := context.Background()
	return &Hub{
		Logger:                  logx.WithContext(ctx),
		ctx:                     ctx,
		Timer:                   timer.NewTimerTask(),
		StopCh:                  make(chan bool),
		Upgrader:                upgrader,
		KefuClients:             newKefuClientMap(),
		VisitorClients:          newVisitorClientMap(),
		MessageList:             make(chan *types.Message, 4000),
		MessageModel:            messageModel,
		ChatRecordModel:         chatRecordModel,
		MessageConfigModel:      messageConfigModel,
		hasSendBeforeHandlerErr: false,
		hasSendAfterHandlerErr:  false,
		hasDbErr:                false,
	}
}

func (h *Hub) WithContext(ctx context.Context) *Hub {
	h.ctx = ctx
	h.Logger = logx.WithContext(ctx)
	return h
}

func (h *Hub) Run() {
	h.DeleteExpiredMessages()
	go h.AssignChanMessage()
	go h.AssignDBMessage()
	h.Timer.Start()
	// 锁住服务，防止服务退出
	h.Block()
}

func (h *Hub) Block() {
	<-h.StopCh
}

func (h *Hub) Stop() {
	h.StopCh <- true
	h.Timer.Close()
	close(h.MessageList)
	close(h.StopCh)
}

// 分配客服
func (h *Hub) AssignKefu() *KefuClient {
	// 前提是有客服在线
	if h.KefuClients.Len() == 0 {
		// 没有客服在线，表示未分配状态，应该执行消息自动回复机制
		return nil
	}
	// 采用最小负载算法，先查询每个客服的负载量
	// 统计容器
	kefuStatisticMap := make(map[int64]int)
	h.VisitorClients.RwMutex.RLock()
	h.KefuClients.RwMutex.RLock()
	for i := range h.VisitorClients.Map {
		kefu_id := h.VisitorClients.Map[i].Visitor.KefuId
		// 确保该客服在线
		if _, ok := h.KefuClients.Map[kefu_id]; ok {
			kefuStatisticMap[kefu_id]++
		}
	}
	h.KefuClients.RwMutex.RUnlock()
	h.VisitorClients.RwMutex.RUnlock()
	// 获取统计容器中的最小值
	var minCount int
	var kefuId int64
	for id, count := range kefuStatisticMap {
		if minCount >= count {
			minCount = count
			kefuId = id
		}
	}
	return h.KefuClients.Get(kefuId)
}

// 通道消息的接收与分发
func (h *Hub) AssignChanMessage() {
	for {
		select {
		case <-h.StopCh:
			// 当接收到 stop 信号时，停止循环任务
			return
		case message := <-h.MessageList:
			h.AssignMessage(message)
		}
	}
}

// 数据库消息的接收与分发
func (h *Hub) AssignDBMessage() {
	ticker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-h.StopCh:
			// 当接收到 stop 信号时，停止循环任务
			return
		case <-ticker.C:
			// 每次 ticker 触发时执行的操作
			// 查找在线的客服
			kefuIds := h.KefuClients.Ids()
			// 查找需要发给客服的消息
			kefuMessages, err := h.MessageModel.FindsToKefu(h.ctx, kefuIds)
			if err != nil {
				h.Logger.Error(err)
				// 出现内部错误，邮件通知管理员
				msg := fmt.Sprintf("读取数据库message记录失败：%s", err)
				emailx.SendNotification(h.ctx, "chat系统异常", msg, []string{emailx.AdminerEmail})
				return
			}
			// 查找在线的游客
			visitorIds := h.VisitorClients.Ids()
			// 查找需要发给游客的消息
			visitorMessages, err := h.MessageModel.FindsToVisitor(h.ctx, visitorIds)
			if err != nil {
				h.Logger.Error(err)
				// 出现内部错误，邮件通知管理员
				msg := fmt.Sprintf("读取数据库message记录失败：%s", err)
				emailx.SendNotification(h.ctx, "chat系统异常", msg, []string{emailx.AdminerEmail})
				return
			}
			mdMessages := append(kefuMessages, visitorMessages...)
			messages := dataconv.MdMessages2ApiMessages(mdMessages)
			for i := range messages {
				select {
				case h.MessageList <- messages[i]:
					// 发送成功
				default:
					//发送失败
				}
			}
			//
		}
	}
}

func (h *Hub) AssignMessage(message *types.Message) {
	switch message.MsgType {
	case 0:
		// 游客消息，转发给客服
		if cli := h.KefuClients.Get(message.KefuId); cli != nil {
			cli.Send(message)
		}
	case 1, 2:
		// 客服消息,公告消息，转发给游客
		if cli := h.VisitorClients.Get(message.VisitorId); cli != nil {
			cli.Send(message)
		}
	}
}

func (h *Hub) SendBeforeHandler(message *types.Message) error {
	// 如果不是客服或者游客的消息，不用保存
	if message.MsgType != 0 && message.MsgType != 1 {
		return nil
	}
	if h.hasSendBeforeHandlerErr {
		return nil
	}
	// 过期时间为3天后
	exp := time.Now().Local().Add(3 * 24 * time.Hour)
	data := &model.Message{
		Id:         0,
		Uuid:       uniqueid.NewUUID(),
		Content:    message.Content,
		VisitorId:  message.VisitorId,
		KefuId:     message.KefuId,
		MsgType:    message.MsgType,
		Status:     message.Status,
		IsRead:     0,
		IsSend:     0,
		SendTime:   time.Now().Local(),
		ExpireTime: exp,
	}
	// 存入数据库
	if _, err := h.MessageModel.Insert(h.ctx, data); err != nil {
		h.hasSendBeforeHandlerErr = true
		h.Logger.Error(err)
		// 出现内部错误，邮件通知管理员
		msg := fmt.Sprintf("聊天消息保存到数据库失败：%s", err)
		emailx.SendNotification(h.ctx, "chat系统异常", msg, []string{emailx.AdminerEmail})
		return err
	}
	return nil
}

// 消息发送成功后的回调，确保已经发送给前端了之后，再调用
func (h *Hub) SendAfterHandler(message *types.Message) error {
	// 如果不是客服或者游客的消息，不用保存
	if message.MsgType != 0 && message.MsgType != 1 {
		return nil
	}
	if h.hasSendAfterHandlerErr {
		return nil
	}
	// 将 message 表中的记录删除，并添加到 chat_record 表中
	err := h.MessageModel.Delete(h.ctx, message.Id)
	if err != nil {
		h.hasSendAfterHandlerErr = true
		h.Logger.Error(err)
		return err
	}
	data := &model.ChatRecord{
		Uuid:      message.Uuid,
		VisitorId: message.VisitorId,
		KefuId:    message.KefuId,
		Status:    message.Status,
		IsRead:    message.IsRead,
		IsSend:    1,
		SendTime:  time.Unix(message.SendTime, 0).Local(),
	}
	_, err = h.ChatRecordModel.Insert(h.ctx, data)
	if err != nil {
		h.hasSendAfterHandlerErr = true
		h.Logger.Error(err)
		return err
	}
	return nil
}

// 每天检查一遍 message 表中过期的数据，并删除
func (h *Hub) DeleteExpiredMessages() error {
	// 每天凌晨2点，删除过期的消息
	_, err := h.Timer.AddTaskByFunc("DeleteExpiredMessages", "0 2 * * *", func() {
		err := h.MessageModel.DeleteExpiredRecords(h.ctx)
		if err != nil {
			h.Logger.Error(err)
			h.Timer.StopTask("DeleteExpiredMessages")
		}
	})
	return err
}
