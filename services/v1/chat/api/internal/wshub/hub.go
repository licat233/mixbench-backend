package wshub

import (
	"context"
	"fmt"
	"mixbench/common/emailx"
	"mixbench/common/errorx"
	"mixbench/common/timer"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

var HubSvc *Hub

type Hub struct {
	logx.Logger
	ctx context.Context

	SvcCtx *svc.ServiceContext

	Timer  timer.Timer
	StopCh chan bool //监听停止信号的通道

	Upgrader       websocket.Upgrader
	KefuClients    *kefuClientMap
	VisitorClients *visitorClientMap
	MessageStore   chan *model.Message

	hasSendBeforeHandlerErr bool
	hasSendAfterHandlerErr  bool
	hasDbErr                bool
}

func NewHub(svcCtx *svc.ServiceContext) *Hub {
	ctx := context.Background()
	HubSvc = &Hub{
		Logger:                  logx.WithContext(ctx),
		ctx:                     ctx,
		SvcCtx:                  svcCtx,
		Timer:                   timer.NewTimerTask(),
		StopCh:                  make(chan bool),
		Upgrader:                upgrader,
		KefuClients:             newKefuClientMap(),
		VisitorClients:          newVisitorClientMap(),
		MessageStore:            make(chan *model.Message, 4000),
		hasSendBeforeHandlerErr: false,
		hasSendAfterHandlerErr:  false,
		hasDbErr:                false,
	}
	return HubSvc
}

func (h *Hub) WithContext(ctx context.Context) *Hub {
	h.ctx = ctx
	h.Logger = logx.WithContext(ctx)
	return h
}

func (h *Hub) Run() (err error) {
	h.DeleteExpiredMessages()
	go func() {
		err = h.AssignChanMessage()
	}()
	go func() {
		h.AssignDBMessage()
	}()
	//开始定时任务
	h.Timer.Start()
	fmt.Println("Starting websocket server...")
	// 锁住服务，防止服务退出
	h.Block()
	return nil
}

func (h *Hub) Block() {
	<-h.StopCh
}

func (h *Hub) Stop() {
	h.StopCh <- true
	h.Timer.Close()
	close(h.MessageStore)
	close(h.StopCh)
}

// 分配客服
func (h *Hub) AssignKefu(projectId int64) *KefuClient {
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
		if kefu, ok := h.KefuClients.Map[kefu_id]; ok {
			//确保该客服是指定项目的
			if kefu.ProjectId == projectId {
				kefuStatisticMap[kefu_id]++
			}
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
func (h *Hub) AssignChanMessage() error {
	for {
		select {
		case <-h.StopCh:
			// 当接收到 stop 信号时，停止循环任务
			return nil
		case message := <-h.MessageStore:
			if err := h.AssignMessage(message); err != nil {
				return err
			}
		}
	}
	// return nil
}

// 数据库消息的接收与分发
// 避免重复发送，当send中的方法发送后，如果来不及移除message库中的message，则会二次发送
// 解决方案：最新的message将不会发送，只发送2分钟以后的message
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
			// 如果没有客服在线
			if len(kefuIds) == 0 {
				continue
			}
			// 查找需要发给客服的消息
			kefuMessages, err := h.SvcCtx.MessageModel.FindsToKefuByKefuIds(h.ctx, kefuIds)
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
			visitorMessages, err := h.SvcCtx.MessageModel.FindsToVisitorByVisitorIds(h.ctx, visitorIds)
			if err != nil {
				h.Logger.Error(err)
				// 出现内部错误，邮件通知管理员
				msg := fmt.Sprintf("读取数据库message记录失败：%s", err)
				emailx.SendNotification(h.ctx, "chat系统异常", msg, []string{emailx.AdminerEmail})
				return
			}
			messages := append(kefuMessages, visitorMessages...)
			for i := range messages {
				//如果不是客服或者游客发送的消息，则跳过
				if messages[i].MsgType != KefuMsgType && messages[i].MsgType != VisitorMsgType {
					continue
				}
				//时间约束，如果是2分钟内的，则忽略
				if time.Since(time.Unix(messages[i].SendTime.Unix(), 0)).Minutes() < 2 {
					continue
				}
				select {
				case h.MessageStore <- messages[i]:
					// 发送成功
				default:
					//发送失败
				}
			}
			//
		}
	}
}

func (h *Hub) AssignMessage(message *model.Message) error {
	switch message.MsgType {
	case VisitorMsgType:
		// 游客消息，转发给客服
		if cli := h.KefuClients.Get(message.KefuId); cli != nil {
			//可能会存在数据库查找错误
			if err := cli.Send(message); err != nil {
				return err
			}
		}
		//如果客服不在线，因为已经持久化过了，所以这里直接抛弃即可，不用管
	case KefuMsgType:
		// 客服消息，公告消息，转发给游客
		if cli := h.VisitorClients.Get(message.VisitorId); cli != nil {
			//可能会存在数据库查找错误
			if err := cli.Send(message); err != nil {
				return err
			}
		}
	case PublicMsgType, SystemMsgType:
		//公告信息，可能发给客服，或者游客
		//依据kefuId和visitorId来做判断
		if message.KefuId > 0 {
			//发给客服
			if cli := h.KefuClients.Get(message.KefuId); cli != nil {
				//可能会存在数据库查找错误
				if err := cli.Send(message); err != nil {
					return err
				}
			}
		}
		if message.VisitorId > 0 {
			//发给游客
			if cli := h.VisitorClients.Get(message.VisitorId); cli != nil {
				//可能会存在数据库查找错误
				if err := cli.Send(message); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// SendBeforeHandler 发送之前的处理
// 持久化处理，保存到message表中
func (h *Hub) SendBeforeHandler(message *model.Message) error {
	// 如果不是客服或者游客的消息，不用保存
	if message.MsgType != KefuMsgType && message.MsgType != VisitorMsgType {
		return nil
	}
	if h.hasSendBeforeHandlerErr {
		return nil
	}
	// 过期时间为3天后
	exp := time.Now().Local().Add(3 * 24 * time.Hour)
	data := &model.Message{
		Id:          message.Id,
		Uuid:        uniqueid.NewUUID(),
		Content:     message.Content,
		ContentType: message.ContentType,
		VisitorId:   message.VisitorId,
		KefuId:      message.KefuId,
		ProjectId:   message.ProjectId,
		MsgType:     message.MsgType,
		Status:      message.Status,
		IsRead:      message.IsRead,
		IsSend:      message.IsSend,
		SendTime:    time.Now().Local(),
		ExpireTime:  exp,
	}
	// 存入数据库
	if _, err := h.SvcCtx.MessageModel.Insert(h.ctx, data); err != nil {
		h.hasSendBeforeHandlerErr = true
		h.Logger.Error(err)
		// 出现内部错误，邮件通知管理员
		msg := fmt.Sprintf("聊天消息保存到数据库失败：%s", err)
		_ = emailx.SendNotification(h.ctx, "chat系统异常", msg, []string{emailx.AdminerEmail})
		return err
	}
	return nil
}

// SendAfterHandler 消息发送成功后的回调，确保已经发送给前端了之后，再调用
func (h *Hub) SendAfterHandler(message *types.Message) error {
	// 如果不是客服或者游客的消息，不用保存
	if message.MsgType != 0 && message.MsgType != 1 {
		return nil
	}
	if h.hasSendAfterHandlerErr {
		return nil
	}
	// 将 message 表中的记录删除，并添加到 chat_record 表中
	err := h.SvcCtx.MessageModel.Delete(h.ctx, message.Id)
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
	_, err = h.SvcCtx.ChatRecordModel.Insert(h.ctx, data)
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
		err := h.SvcCtx.MessageModel.DeleteExpiredRecords(h.ctx)
		if err != nil {
			h.Logger.Error(err)
			h.Timer.StopTask("DeleteExpiredMessages")
		}
	})
	return err
}

// TransferVisitor 转移有游客给其它客服
func (l *Hub) TransferVisitor(visitorId, kefuId int64) (err error) {
	if visitorId <= 0 {
		return errorx.ExternalError("visitorId <=0", "visitorId must be greater than 0")
	}

	//更新 clients 中心的游客信息
	//判断是否还在线
	visitorCli := l.VisitorClients.Get(visitorId)
	if visitorCli != nil {
		//如果还在线，则修改
		visitorCli.Visitor.KefuId = kefuId
	}
	//如果不在线，就不用修改

	//先判断该游客是否存在
	visitor, err := l.SvcCtx.VisitorModel.FindOne(l.ctx, visitorId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}
	//如果不存在，则返回错误
	if err == model.ErrNotFound {
		return errorx.ExternalError("the visitorId does not exist")
	}

	//先判断该客服是否存在
	_, err = l.SvcCtx.KefuModel.FindOne(l.ctx, kefuId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}
	//如果不存在，则返回错误
	if err == model.ErrNotFound {
		return errorx.ExternalError("the kefuId does not exist")
	}

	//修改该游客的接待客服id
	visitor.KefuId = kefuId
	//更新游客信息
	if err := l.SvcCtx.VisitorModel.Update(l.ctx, visitor); err != nil {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}

	//再更新message缓存的数据
	//查找出游客发出去的信息
	mdMsgList, _, err := l.SvcCtx.MessageModel.FindList(l.ctx, -1, -1, "", &model.Message{
		VisitorId: visitorId,
		MsgType:   VisitorMsgType,
	})
	if err != nil {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}
	// 必须一条一条的更新，才能更新redis缓存
	for _, msg := range mdMsgList {
		msg.KefuId = kefuId
		if err := l.SvcCtx.MessageModel.Update(l.ctx, msg); err != nil {
			l.Logger.Error(err)
			return errorx.InternalError(err)
		}
	}

	//更新 chat_record 数据
	mdRecordList, _, err := l.SvcCtx.ChatRecordModel.FindList(l.ctx, -1, -1, "", &model.ChatRecord{
		VisitorId: visitorId,
		MsgType:   VisitorMsgType,
	})
	if err != nil {
		l.Logger.Error(err)
		return errorx.InternalError(err)
	}
	// 必须一条一条的更新，才能更新redis缓存
	for _, record := range mdRecordList {
		record.KefuId = kefuId
		if err := l.SvcCtx.ChatRecordModel.Update(l.ctx, record); err != nil {
			l.Logger.Error(err)
			return errorx.InternalError(err)
		}
	}

	//还需要修改 send 通道中的消息
	//需要再 send 前做校验逻辑，代码不在这里写
	// 已在 client.go 内函数 (c *Client) Send(message *types.Message) error 中实现

	return nil
}
