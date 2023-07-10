package wshub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

type VisitorClient struct {
	*Client
	Visitor *model.Visitor
}

// 新建游客client，如果初始化失败，会返回 nil, nil
func NewVisitorClient(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request) (*VisitorClient, error) {
	if hub == nil {
		return nil, errors.New("hub is nil")
	}
	var isOk = false
	conn, err := NewWsConn(w, r)
	if err != nil {
		return nil, fmt.Errorf("newWsConn error: %s", err)
	}
	defer func() {
		if !isOk {
			if conn != nil {
				conn.Close()
			}
			// CloseWsConn(conn)
		}
	}()
	c := &VisitorClient{
		Client: &Client{
			Logger:    logx.WithContext(ctx),
			ctx:       ctx,
			hub:       hub,
			conn:      conn,
			send:      make(chan *model.Message, 100),
			Stop:      make(chan bool),
			Status:    true,
			ProjectId: 0,
			w:         w,
			r:         r,
		},
		Visitor: nil,
	}
	// 鉴权
	ok, id, err := c.Identify()
	if err != nil {
		return nil, fmt.Errorf("token identify error: %v", err)
	}
	if !ok || id <= 0 {
		return nil, nil
	}

	//查看是否存在该游客
	visitor, err := c.hub.SvcCtx.VisitorModel.FindOne(c.ctx, id)
	if err != nil && err != model.ErrNotFound {
		c.Logger.Error(err)
		return nil, err
	}

	if err == model.ErrNotFound {
		return nil, nil
	}

	//查找是否存在该项目
	_, err = c.hub.SvcCtx.ProjectModel.FindOne(c.ctx, visitor.ProjectId)
	if err != nil && err != model.ErrNotFound {
		c.Logger.Error(err)
		return nil, err
	}

	if err == model.ErrNotFound {
		return nil, nil
	}

	c.Visitor = visitor
	c.Client.ProjectId = visitor.ProjectId

	// 注册
	hub.VisitorClients.Set(c)

	//更新数据库中的数据为在线状态
	visitor.Status = 1
	err = c.hub.SvcCtx.VisitorModel.Update(c.ctx, visitor)
	if err != nil {
		c.Logger.Error(err)
		return nil, err
	}

	isOk = true

	return c, nil
}

func (c *VisitorClient) Run() (err error) {
	var WritePumpError, ReadPumpError error
	var wg sync.WaitGroup
	// 开始收发ws信息
	wg.Add(1)
	go func() {
		defer wg.Done()
		ReadPumpError = c.ReadPump()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		WritePumpError = c.WritePump()
	}()

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
				return nil
			}
			// 防止某些游客搞骚操作，更改数据，这里的messageContent仅仅是 content ，不是message的json数据
			messageContent = preprocessMessageContent(messageContent)
			content := string(messageContent)
			if strings.ToLower(content) == "ping" {
				//ping 不用管
				continue
			}

			// 如果该游客没有客服接待
			if c.Visitor.KefuId <= 0 {
				//尝试分配客服
				kefuCli := c.hub.AssignKefu(c.ProjectId)
				//如果没有客服
				if kefuCli == nil {
					//则缓存这条消息，待客服上线后，给与回复

					// continue
				} else {
					//如果有客服，将其分配给该游客
					c.Visitor.KefuId = kefuCli.Kefu.Id
				}
			}

			message := c.NewMessage(content)

			//入栈前，持久化一下，防止消息在通道中意外丢失
			if err := c.hub.SendBeforeHandler(message); err != nil {
				return fmt.Errorf("SendBeforeHandler error: %s", err)
			}
			//有客服接待，则发送
			if c.Visitor.KefuId > 0 {
				c.hub.MessageStore <- message
			}
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

	//已经准备好接收信息了，那就把缓存的游客消息全部发送给该游客
	mdMsgList, err := c.hub.SvcCtx.MessageModel.FindsToVisitorByVisitorId(c.ctx, c.Visitor.Id)
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
			if err := c.WriteMessage(dataconv.MdMessageToApiMessage(message)); err != nil {
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

func (c *VisitorClient) Logout() error {
	if !c.Client.Status {
		return nil
	}
	c.Client.Status = false
	c.hub.VisitorClients.Del(c.Visitor.Id)

	//更新数据库中的数据为离线状态
	c.Visitor.Status = 0
	err := c.hub.SvcCtx.VisitorModel.Update(c.ctx, c.Visitor)
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	return nil
}

func (c *VisitorClient) NewMessage(content string) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     content,
		ContentType: TextContentType,
		VisitorId:   c.Visitor.Id,
		KefuId:      c.Visitor.KefuId,
		MsgType:     VisitorMsgType, //目前仅支持普通文本的消息
		Status:      0,
		IsRead:      0,
		IsSend:      0,
		SendTime:    time.Now().Local(),
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
