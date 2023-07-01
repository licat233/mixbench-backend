package client

import (
	"context"
	"encoding/json"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/types"
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	c := &VisitorClient{
		Client: &Client{
			Logger: logx.WithContext(ctx),
			ctx:    ctx,
			hub:    hub,
			conn:   conn,
			send:   make(chan *types.Message, 100),
			Stop:   make(chan bool),
		},
		Visitor: visitor,
	}
	// 注册
	hub.VisitorClients.Set(c)
	return c, nil
}

func (c *VisitorClient) Run() {
	// 开始收发ws信息
	go c.WritePump()
	go c.ReadPump()
	//等待关闭信号
	<-c.Stop
}

// 读取游客通过ws发送过来的信息
func (c *VisitorClient) ReadPump() {
	defer func() {
		c.hub.VisitorClients.Del(c.Visitor.Id)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, messageContent, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger.Error("websocket conn.close error: %v", err)
			}
			c.Client.SystemErrorHandler()
			return
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
			SendTime:   0,
			ExpireTime: 0,
		}
		//入栈前，持久化一下，防止消息在通道中意外丢失
		c.hub.SendBeforeHandler(message)
		c.hub.MessageList <- message
	}
}

// 给游客的ws写入消息
func (c *VisitorClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.hub.VisitorClients.Del(c.Visitor.Id)
	}()

	// 发送欢迎信息
	if c.Client.SendWelcomeMessage() != nil {
		return
	}

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					c.Logger.Error(err)
				}
				c.Client.SystemErrorHandler()
				return
			}
			body, _ := json.Marshal(*message)
			if err := c.conn.WriteMessage(websocket.TextMessage, body); err != nil {
				c.Logger.Error(err)
				c.Client.SystemErrorHandler()
				return
			}
			_ = c.hub.SendAfterHandler(message)
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.Logger.Error(err)
				c.Client.SystemErrorHandler()
				return
			}
		}
	}
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

// func (l *visitorClientMap) Lock() {
// 	l.RwMutex.Lock()
// }

// func (l *visitorClientMap) Unlock() {
// 	l.RwMutex.Unlock()
// }

// func (l *visitorClientMap) RLock() {
// 	l.RwMutex.RLock()
// }

// func (l *visitorClientMap) RUnlock() {
// 	l.RwMutex.RUnlock()
// }
