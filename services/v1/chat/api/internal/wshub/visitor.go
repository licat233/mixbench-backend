package wshub

import (
	"context"
	"encoding/json"
	"fmt"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorClient struct {
	*Client
	Visitor *types.Visitor
}

func NewVisitorClient(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request, visitor *types.Visitor) (*VisitorClient, error) {
	conn, err := NewWsConn(w, r)
	if err != nil {
		return nil, fmt.Errorf("newWsConn error: %s", err)
	}
	c := &VisitorClient{
		Client: &Client{
			Logger: logx.WithContext(ctx),
			ctx:    ctx,
			hub:    hub,
			conn:   conn,
			send:   make(chan *types.Message, 100),
			Stop:   make(chan bool),
			Status: true,
			w:      w,
			r:      r,
		},
		Visitor: visitor,
	}
	// 注册
	hub.VisitorClients.Set(c)
	return c, nil
}

func (c *VisitorClient) Run() (err error) {
	var WritePumpError, ReadPumpError error
	var wg sync.WaitGroup
	// 开始收发ws信息
	wg.Add(1)
	go func() {
		defer wg.Done()
		WritePumpError = c.WritePump()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadPumpError = c.ReadPump()
	}()

	//等待关闭信号, 这里可要可不要
	// <-c.Stop
	wg.Wait()
	if WritePumpError != nil {
		err = fmt.Errorf("WritePumpError: %s", WritePumpError)
	}
	if ReadPumpError != nil {
		err = fmt.Errorf("%s + ReadPumpError: %s", err, ReadPumpError)
	}
	return
}

// 读取游客通过ws发送过来的信息
func (c *VisitorClient) ReadPump() error {
	defer func() {
		c.Logout()
	}()
	for {
		select {
		case <-c.Stop:
			return nil
		default:
			_, messageContent, err := c.conn.ReadMessage()
			if err != nil {
				//conn被正常关闭
				return nil
				// if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				// 	//conn被正常关闭
				// 	return nil
				// } else {
				// 	//异常关闭
				// 	c.Logger.Error("websocket conn is unexpected close error: %v", err)
				// 	return fmt.Errorf("websocket conn is unexpected close error: %v", err)
				// }
			}
			// 如果该游客没有客服接待
			if c.Visitor.KefuId <= 0 {
				//尝试分配客服
				kefuCli := c.hub.AssignKefu()
				//如果没有客服
				if kefuCli == nil {
					//只能抛弃这条消息
					continue
				}
				//如果有客服，将其分配给该游客
				c.Visitor.KefuId = kefuCli.Kefu.Id
			}
			// 防止某些游客搞骚操作，更改数据，这里的messageContent仅仅是 content ，不是message的json数据
			messageContent = preprocessMessageContent(messageContent)
			message := &types.Message{
				Id:         0,
				Uuid:       uniqueid.NewUUID(),
				Content:    string(messageContent),
				VisitorId:  c.Visitor.Id,
				KefuId:     c.Visitor.KefuId,
				MsgType:    0,
				Status:     0,
				IsRead:     0,
				IsSend:     0,
				SendTime:   time.Now().Local().Unix(),
				ExpireTime: 0,
			}
			//入栈前，持久化一下，防止消息在通道中意外丢失
			if err := c.hub.SendBeforeHandler(message); err != nil {
				// c.Client.SystemErrorHandler(-1, c.Visitor.Id)
				return fmt.Errorf("SendBeforeHandler error: %s", err)
			}
			c.hub.MessageList <- message
		}
	}
}

// 给游客的ws写入消息
func (c *VisitorClient) WritePump() error {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Logout()
	}()

	// 如果该游客有客服接待
	if c.Visitor.KefuId > 0 {
		// 发送欢迎信息
		if err := c.Client.SendPresetMessage(WelcomeMessageKey, -1, c.Visitor.Id); err != nil {
			c.Logger.Error(err)
			return fmt.Errorf("send welcome message error: %s", err)
		}
	} else {
		// 没有分配接待的客服，该怎样处理？说明没有客服在线，无法分配，此时应该发送留言信息
		// 发送留言信息
		if err := c.Client.SendPresetMessage(LeaveMessageKey, -1, c.Visitor.Id); err != nil {
			c.Logger.Error(err)
			return fmt.Errorf("send leave a message error: %s", err)
		}
	}

	//已经准备好接收信息了，那就把缓存的游客消息全部发送
	mdMsgList, _, err := c.hub.SvcCtx.MessageModel.FindList(c.ctx, -1, -1, "", &model.Message{
		VisitorId: c.Visitor.Id,
		MsgType:   KefuMsgType,
	})
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	msgList := dataconv.MdMessages2ApiMessages(mdMsgList)
	for _, msg := range msgList {
		if err := c.WriteMessage(msg); err != nil {
			return err
		}
	}

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetReadDeadline(time.Now().Add(pongWait))
			if !ok {
				// The hub closed the channel.
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
						//正常关闭
						return nil
					} else {
						//异常关闭
						c.Logger.Error("websocket conn IsUnexpectedCloseError: %v", err)
						return fmt.Errorf("websocket conn IsUnexpectedCloseError: %v", err)
					}
				}
				// 已关闭，正常退出
				return nil
			}
			if err := c.WriteMessage(message); err != nil {
				return err
			}
		case <-c.Stop:
			//关闭了
			return nil
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.Logger.Error(err)
				return fmt.Errorf("conn.WriteMessage error: %s", err)
			}
		}
	}
}

func (c *VisitorClient) WriteMessage(message *types.Message) error {
	body, _ := json.Marshal(*message)
	if err := c.conn.WriteMessage(websocket.TextMessage, body); err != nil {
		c.Logger.Error(err)
		return fmt.Errorf("conn.WriteMessage error: %s", err)
	}
	if err := c.hub.SendAfterHandler(message); err != nil {
		c.Logger.Error(err)
		return fmt.Errorf("SendAfterHandler error: %s", err)
	}
	return nil
}

func (c *VisitorClient) Logout() {
	if !c.Client.Status {
		return
	}
	c.Client.Status = false
	c.hub.VisitorClients.Del(c.Visitor.Id)
}

type visitorClientMap struct {
	Map     map[int64]*VisitorClient
	RwMutex *sync.RWMutex
}

func newVisitorClientMap() *visitorClientMap {
	return &visitorClientMap{
		Map:     map[int64]*VisitorClient{},
		RwMutex: &sync.RWMutex{},
	}
}

func (l *visitorClientMap) Ids() []int64 {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	var res []int64
	for i := range l.Map {
		res = append(res, l.Map[i].Visitor.Id)
	}
	return res
}

func (l *visitorClientMap) Set(vc *VisitorClient) {
	if vc == nil {
		return
	}
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	l.Map[vc.Visitor.Id] = vc
}

func (l *visitorClientMap) Del(id int64) {
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	if v, ok := l.Map[id]; ok {
		v.Close()
		delete(l.Map, id)
	}
}

func (l *visitorClientMap) Get(id int64) *VisitorClient {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	return l.Map[id]
}

func (l *visitorClientMap) Len() int {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	return len(l.Map)
}

// // GetWaitVisitors 获取正在等待客服接待的游客
// func (l *visitorClientMap) GetWaitVisitors() []*types.Visitor {
// 	l.RwMutex.RLock()
// 	defer l.RwMutex.RUnlock()
// 	var res []*types.Visitor
// 	for i := range l.Map {
// 		if l.Map[i].Visitor.KefuId <= 0 {
// 			res = append(res, l.Map[i].Visitor)
// 		}
// 	}
// 	return res
// }

// GetByKefuId 通过kefuid查找
func (l *visitorClientMap) GetByKefuId(kefuId int64) []*VisitorClient {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	var res []*VisitorClient
	for i := range l.Map {
		if l.Map[i].Visitor.KefuId == kefuId || l.Map[i].Visitor.KefuId <= 0 {
			res = append(res, l.Map[i])
		}
	}
	return res
}
