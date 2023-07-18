package ws

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/chat/api/internal/logic/dataconv"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/model"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
)

// client的使用者接口
type user interface {
	GetId() int64
	GetUuid() string
	GetProjectId() int64
	IsKefu() bool
	IsVisitor() bool
	// NewMessage(content string, receiverId int64) *model.Message
	// GetReceiverId() int64
	// SetReceiverId(receiverId int64)
}

type client struct {
	Logger logx.Logger
	ctx    context.Context

	svcCtx *svc.ServiceContext

	conn *websocket.Conn

	once sync.Once //用于执行一次性操作，避免重复操作，比如close操作
	stop chan bool //关闭信号

	w http.ResponseWriter
	r *http.Request

	pubsub *goredis.PubSub

	wg sync.WaitGroup

	user user
}

func newClient(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request, u user) (*client, error) {
	if !utils.IsWebsocketRequest(r) {
		// 不是 WebSocket 请求
		return nil, errors.New("the current request is not a WebSocket request")
	}
	if u == nil {
		return nil, errorx.New("user == nil")
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err != nil && conn != nil {
			conn.Close()
		}
	}()

	conn.SetWriteDeadline(time.Now().Add(writeWait))
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	channls := []string{u.GetUuid()}
	if u.IsKefu() {
		//如果是客服，上线的时候，应该即使接待没有客服的游客
		channls = append(channls, WaiteKefuChannelValue)
	}
	pubsub := svcCtx.RedisCli.Subscribe(ctx, channls...)

	return &client{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		conn:   conn,
		once:   sync.Once{},
		stop:   make(chan bool),
		w:      w,
		r:      r,
		pubsub: pubsub,
		wg:     sync.WaitGroup{},
		user:   u,
	}, nil
}

func (c *client) Close() {
	c.once.Do(func() {
		c.Stop()
		c.pubsub.Close()
		// 关闭ws.conn
		c.conn.Close()
		if c.user.IsKefu() {
			OnlineKefu.Del(c.user.GetId())
		}
		if c.user.IsVisitor() {
			OnlineVisitor.Del(c.user.GetId())
		}

	})
}

func (c *client) Stop() {
	select {
	case _, ok := <-c.stop:
		if ok {
			close(c.stop)
		}
	default:
		// chan 已经关闭
	}
}

func (c *client) NewPublicMessage(content string, kefuId, visitorId int64) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     content,
		ContentType: TextContentType,
		VisitorId:   visitorId,
		KefuId:      kefuId,
		ProjectId:   c.user.GetProjectId(),
		MsgType:     PublicMsgType,
		Status:      0,
		IsRead:      0,
		IsSend:      0,
	}
}

func (c *client) assignMessage(message *model.Message) error {
	if message == nil {
		return nil
	}
	message.Content = preprocessMessageContent(message.Content)
	if message.Content == "" {
		//空消息
		return nil
	}
	switch message.MsgType {
	case VisitorMsgType:
		// 游客消息，转发给客服
		return c.sendMessageToKefu(message)
	case KefuMsgType:
		// 客服消息，公告消息，转发给游客
		return c.sendMessageToVisitor(message)
	case PublicMsgType:
		//公告信息，可能发给客服，或者游客
		//根据当前client的uuid来指定接收者，即发给当前client的user
		//直接发给user
		return c.publishMessage(c.user.GetUuid(), message)
	case SystemMsgType:
		//系统消息，只能发给客服
		return c.sendMessageToKefu(message)
	}
	return nil
}

func (c *client) sendMessageToKefu(message *model.Message) error {
	if message == nil {
		return nil
	}
	// TODO: 查找出客服账号，取出uuid，发布至redis
	kefuId := message.KefuId
	//如果没有接收者，默认的uuid
	kefuUuid := WaiteKefuChannelValue
	if kefuId > 0 {
		resp, err := c.svcCtx.AdminRpc.GetAdminer(c.ctx, &admin_pb.GetAdminerReq{Id: kefuId})
		if err != nil {
			return err
		}
		if resp.Adminer == nil {
			return errors.New("adminer(kefu) not found")
		}
		kefuUuid = resp.Adminer.Uuid
	}

	return c.publishMessage(kefuUuid, message)
}

func (c *client) sendMessageToVisitor(message *model.Message) error {
	if message == nil {
		return nil
	}
	// TODO: 查找出游客账号，取出uuid，发布至redis
	visitorId := message.VisitorId
	//如果游客id不存在，则抛弃
	if visitorId < 1 {
		return nil
	}
	visitor, err := c.svcCtx.VisitorModel.FindById(c.ctx, visitorId)
	if err != nil && err != model.ErrNotFound {
		return err
	}
	if err == model.ErrNotFound {
		return errors.New("visitor not found")
	}
	visitorUuid := visitor.Uuid
	return c.publishMessage(visitorUuid, message)
}

// SendPresetMessage 发送预设的公告消息
// 不经过 redis 订阅，直接写入ws
func (c *client) sendPresetMessage(msgKey string, kefuId, visitorId int64) error {
	if msgKey = strings.TrimSpace(msgKey); msgKey == "" {
		return nil
	}

	msgList, err := c.svcCtx.MessageConfigModel.FindsByMsgkeyProjectId(c.ctx, msgKey, c.user.GetProjectId())
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	// 如果存在，则发送
	var messages []*model.Message
	for _, msg := range msgList {
		if content := strings.TrimSpace(msg.Content); content != "" {
			messages = append(messages, c.NewPublicMessage(content, kefuId, visitorId))
		}
	}
	return c.writeToWebsocket(messages...)
}

// message写入ws
func (c *client) writeToWebsocket(messages ...*model.Message) error {
	list := dataconv.MdMessages2ApiMessages(messages)
	for _, msg := range list {
		if content := strings.TrimSpace(msg.Content); content != "" {
			data, err := json.Marshal(*msg)
			if err != nil {
				err = fmt.Errorf("json.Marshal error: %s", err)
				c.Logger.Error(err)
				return err
			}
			err = c.conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 分配客服
func (c *client) assignKefu() (kefuId int64) {
	id := OnlineKefu.AssignKefu()
	return id
}

func (c *client) publishMessage(channel string, message *model.Message) error {
	payload, err := json.Marshal(*message)
	if err != nil {
		return err
	}
	// 发布消息
	if err = c.svcCtx.RedisCli.Publish(c.ctx, channel, string(payload)).Err(); err != nil {
		return err
	}
	return nil
}

// ws读取
func (c *client) readPump(payloadHandler payloadHandler) {
	defer func() {
		c.Stop()
	}()

	for {
		select {
		case <-c.stop:
			return
		default:
			messageType, payload, err := c.conn.ReadMessage()
			if err != nil {
				return
			}
			if messageType == websocket.CloseMessage {
				// 连接已被客户端关闭
				return
			} else if messageType == websocket.PingMessage {
				continue
			} else if messageType == websocket.PongMessage {
				continue
			} else if messageType == websocket.BinaryMessage {
				//二进制消息，丢弃，暂不支持
				continue
			}

			message, err := payloadHandler(payload)
			if err != nil {
				//有错误，关闭
				return
			}
			if message == nil {
				//丢弃
				continue
			}

			//改用redis发布方案：
			if err = c.assignMessage(message); err != nil {
				c.Logger.Error(err)
				return
			}
		}
	}
}

// ws写入
func (c *client) writePump(recevierMessageHandler recevierMessageHandler) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Stop()
	}()

	for {
		select {
		case <-c.stop:
			return
		case msg, ok := <-c.pubsub.Channel():
			if !ok {
				//ch已经关闭
				return
			}
			var message model.Message
			if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
				//不合规的消息体。丢弃
				continue
			}
			if err := recevierMessageHandler(&message); err != nil {
				return
			}
			if err := c.writeMessage(&message); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 往websocket中写入消息
func (c *client) writeMessage(message *model.Message) error {
	if message == nil {
		return nil
	}
	msg := dataconv.MdMessageToApiMessage(message)
	body, err := json.Marshal(*msg)
	if err != nil {
		err = fmt.Errorf("json.Marshal error: %s", err)
		c.Logger.Error(err)
		return err
	}
	if err := c.conn.WriteMessage(websocket.TextMessage, body); err != nil {
		//如果发送失败了，如果是客服或者游客的消息，避免该消息丢失，应该保存到redis中，等以后再发
		if message.MsgType == KefuMsgType || message.MsgType == VisitorMsgType {
			if e := c.assignMessage(message); e != nil {
				e = fmt.Errorf("c.sendMessage error: %s", e)
				c.Logger.Error(e)
				return e
			}
		}
		return err
	}
	return nil
}
