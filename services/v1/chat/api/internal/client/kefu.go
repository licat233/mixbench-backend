package client

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

type KefuClient struct {
	*Client
	Kefu *types.Kefu
}

func NewKefuClient(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request, kefu *types.Kefu) (*KefuClient, error) {
	conn, err := NewWsConn(w, r)
	if err != nil {
		return nil, fmt.Errorf("newWsConn error: %s", err)
	}
	c := &KefuClient{
		Client: &Client{
			Logger: logx.WithContext(ctx),
			ctx:    ctx,
			hub:    hub,
			conn:   conn,
			send:   make(chan *types.Message, 20),
			Stop:   make(chan bool),
			w:      w,
			r:      r,
		},
		Kefu: kefu,
	}
	// 注册
	hub.KefuClients.Set(c)
	return c, nil
}

func (c *KefuClient) Run() (err error) {
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

// 读取客服通过ws发送过来的信息
func (c *KefuClient) ReadPump() error {
	defer func() {
		c.Logout()
	}()
	for {
		select {
		case <-c.Stop:
			return nil
		default:
			_, messageBody, err := c.conn.ReadMessage()
			if err != nil {
				//异常关闭
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					c.Logger.Error("websocket conn IsUnexpectedCloseError: %v", err)
					return fmt.Errorf("websocket conn IsUnexpectedCloseError: %v", err)
				}
				//conn被正常关闭
				return nil
			}
			// messageBody 为一个完整的 message 的 json 数据
			var message types.Message
			if err := json.Unmarshal(messageBody, &message); err != nil {
				// 不符合规范的 messageBody，丢弃
				continue
			}
			if message.VisitorId <= 1 {
				c.Logger.Error("unvalidtor visitor id: %d", message.VisitorId)
				continue
			}

			message.Uuid = uniqueid.NewUUID()
			message.KefuId = c.Kefu.Id
			message.MsgType = 1

			//入栈前，持久化一下，防止消息在通道中意外丢失
			if err := c.hub.SendBeforeHandler(&message); err != nil {
				c.Client.SystemErrorHandler()
				return fmt.Errorf("SendBeforeHandler error: %s", err)
			}
			c.hub.MessageList <- &message
		}
	}
}

// 给客服的ws写入消息
func (c *KefuClient) WritePump() error {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Logout()
	}()

	//发送公告信息
	if err := c.Client.SendPresetMessage(PublicMessageKey); err != nil {
		c.Logger.Error(err)
		return fmt.Errorf("send public message error: %s", err)
	}

	//已经准备好接收信息了，那就把缓存的游客消息全部发送
	mdMsgList, _, err := c.hub.MessageModel.FindList(c.ctx, -1, -1, "", &model.Message{
		KefuId:  c.Kefu.Id,
		MsgType: VisitorMsgType,
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
			if !ok {
				// The hub closed the channel.
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					c.Logger.Error(err)
					return fmt.Errorf("conn.WriteMessage error: %s", err)
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

func (c *KefuClient) WriteMessage(message *types.Message) error {
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

func (c *KefuClient) Logout() {
	c.hub.KefuClients.Del(c.Kefu.Id)
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
