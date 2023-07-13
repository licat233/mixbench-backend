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

type user interface {
	GetId() int64
	GetUuid() string
	GetProjectId() int64
	IsKefu() bool
	IsVisitor() bool
}

type client struct {
	Logger logx.Logger
	ctx    context.Context

	svcCtx *svc.ServiceContext

	conn *websocket.Conn

	stop   chan bool //关闭信号
	status bool      //状态

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

	return &client{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		conn:   conn,
		stop:   make(chan bool),
		w:      w,
		r:      r,
		pubsub: svcCtx.RedisCli.Subscribe(ctx, u.GetUuid()),
		wg:     sync.WaitGroup{},
		user:   u,
	}, nil
}

func (c *client) Close() {
	c.status = false
	c.pubsub.Close()
	c.conn.Close()
	close(c.stop)
	if c.user.IsKefu() {
		OnlineKefu.Del(c.user.GetId())
	}
	if c.user.IsVisitor() {
		OnlineVisitor.Del(c.user.GetId())
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

func (c *client) sendMessage(message *model.Message) error {
	message.Content = preprocessMessageContent(message.Content)
	if message.Content == "" {
		//空消息
		return nil
	}
	switch message.MsgType {
	case VisitorMsgType:
		// 游客消息，转发给客服
		if err := c.sendMessageToKefu(message); err != nil {
			return err
		}
	case KefuMsgType:
		// 客服消息，公告消息，转发给游客
		if err := c.sendMessageToVisitor(message); err != nil {
			return err
		}
	case PublicMsgType, SystemMsgType:
		//公告信息，可能发给客服，或者游客
		//依据kefuId和visitorId来做判断
		if message.KefuId > 0 {
			//发给客服
			// 游客消息，转发给客服
			if err := c.sendMessageToKefu(message); err != nil {
				return err
			}
		}
		if message.VisitorId > 0 {
			//发给游客
			if err := c.sendMessageToVisitor(message); err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func (c *client) sendMessageToKefu(message *model.Message) error {
	// TODO: 查找出客服账号，取出uuid，发布至redis
	kefuId := message.KefuId
	resp, err := c.svcCtx.AdminRpc.GetAdminer(c.ctx, &admin_pb.GetAdminerReq{Id: kefuId})
	if err != nil {
		return err
	}
	if resp.Adminer == nil {
		return errors.New("adminer(kefu) not found")
	}
	kefuUuid := resp.Adminer.Uuid
	return c.publishMessage(kefuUuid, message)
}

func (c *client) sendMessageToVisitor(message *model.Message) error {
	// TODO: 查找出游客账号，取出uuid，发布至redis
	visitorId := message.VisitorId
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

// SendPresetMessage 发送预设的消息
func (c *client) sendPresetMessage(msgKey string, kefuId, visitorId int64) error {
	if msgKey = strings.TrimSpace(msgKey); msgKey == "" {
		return nil
	}

	msgList, err := c.svcCtx.MessageConfigModel.FindsByMsgkeyProjectId(c.ctx, msgKey, c.user.GetProjectId())
	if err != nil {
		return err
	}
	// 如果存在，则发送
	for _, msg := range msgList {
		if content := strings.TrimSpace(msg.Content); content != "" {
			welcomeMsg := c.NewPublicMessage(content, kefuId, visitorId)
			if err := c.sendMessage(welcomeMsg); err != nil {
				return err
			}
		}
	}
	return nil
}

// 往websocket中写入消息
func (c *client) writeMessage(message *model.Message) error {
	msg := dataconv.MdMessageToApiMessage(message)
	body, err := json.Marshal(*msg)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %s", err)
	}
	if err := c.conn.WriteMessage(websocket.TextMessage, body); err != nil {
		//如果发送失败了，避免该消息丢失，应该保存到redis中，等以后再发
		if e := c.sendMessage(message); e != nil {
			return fmt.Errorf("client.SendMessage error: %s", err)
		}
		return fmt.Errorf("conn.WriteMessage error: %s", err)
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
