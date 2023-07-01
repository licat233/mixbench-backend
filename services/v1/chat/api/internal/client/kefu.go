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

type KefuClient struct {
	*Client
	Kefu *types.Kefu
}

func NewKefuClient(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request, kefu *types.Kefu) (*KefuClient, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	c := &KefuClient{
		Client: &Client{
			Logger: logx.WithContext(ctx),
			ctx:    ctx,
			hub:    hub,
			conn:   conn,
			send:   make(chan *types.Message, 20),
			Stop:   make(chan bool),
		},
		Kefu: kefu,
	}
	// 注册
	hub.KefuClients.Set(c)
	return c, nil
}

func (c *KefuClient) Run() {
	// 开始收发ws信息
	go c.WritePump()
	go c.ReadPump()
	//等待关闭信号
	<-c.Stop
}

// 读取客服通过ws发送过来的信息
func (c *KefuClient) ReadPump() {
	defer func() {
		c.hub.KefuClients.Del(c.Kefu.Id)
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, messageBody, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger.Error("websocket conn.close error: %v", err)
			}
			c.Client.SystemErrorHandler()
			return
		}
		// messageBody 为一个完整的 message 的 json 数据
		var message types.Message
		json.Unmarshal(messageBody, &message)
		if message.VisitorId <= 1 {
			c.Logger.Error("unvalidtor visitor id: %d", message.VisitorId)
			continue
		}

		message.Uuid = uniqueid.NewUUID()
		message.KefuId = c.Kefu.Id
		message.MsgType = 1

		//入栈前，持久化一下，防止消息在通道中意外丢失
		c.hub.SendBeforeHandler(&message)
		c.hub.MessageList <- &message
	}
}

// 给客服的ws写入消息
func (c *KefuClient) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.hub.KefuClients.Del(c.Kefu.Id)
	}()

	//发送公告信息
	if c.Client.SendPublicMessage() != nil {
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
				return
			}
			body, _ := json.Marshal(*message)
			if err := c.conn.WriteMessage(websocket.TextMessage, body); err != nil {
				c.Logger.Error(err)
				return
			}
			_ = c.hub.SendAfterHandler(message)
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.Logger.Error(err)
				return
			}
		}
	}
}

type kefuClientMap struct {
	Map     map[int64]*KefuClient
	RwMutex *sync.RWMutex
}

func newKefuClientMap() *kefuClientMap {
	return &kefuClientMap{
		Map:     map[int64]*KefuClient{},
		RwMutex: &sync.RWMutex{},
	}
}

func (l *kefuClientMap) Ids() []int64 {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	var res []int64
	for i := range l.Map {
		res = append(res, l.Map[i].Kefu.Id)
	}
	return res
}

func (l *kefuClientMap) Set(vc *KefuClient) {
	if vc == nil {
		return
	}
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	l.Map[vc.Kefu.Id] = vc
}

func (l *kefuClientMap) Del(id int64) {
	l.RwMutex.Lock()
	defer l.RwMutex.Unlock()
	if v, ok := l.Map[id]; ok {
		v.Close()
		delete(l.Map, id)
	}
}

func (l *kefuClientMap) Get(id int64) *KefuClient {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	return l.Map[id]
}

func (l *kefuClientMap) Len() int {
	l.RwMutex.RLock()
	defer l.RwMutex.RUnlock()
	return len(l.Map)
}

// func (l *kefuClientMap) Lock() {
// 	l.RwMutex.Lock()
// }

// func (l *kefuClientMap) Unlock() {
// 	l.RwMutex.Unlock()
// }

// func (l *kefuClientMap) RLock() {
// 	l.RwMutex.RLock()
// }

// func (l *kefuClientMap) RUnlock() {
// 	l.RwMutex.RUnlock()
// }
