package wshub

import (
	"context"
	"fmt"
	"mixbench/common/jwtx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/model"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	Logger logx.Logger
	ctx    context.Context
	hub    *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send      chan *model.Message
	Stop      chan bool //关闭信号
	Status    bool      //状态，是否正常
	ProjectId int64
	w         http.ResponseWriter
	r         *http.Request
}

func (c *Client) Close() {
	c.Stop <- true
	close(c.send)
	CloseWsConn(c.conn)
}

// Send 主要用来防止阻塞，用于其它逻辑，请注意区分消息类型
func (c *Client) Send(message *model.Message) error {
	//发送前，需要校验，该游客信息是否已经被修改了客服号
	if message.MsgType == VisitorMsgType {
		msg, err := c.hub.SvcCtx.MessageModel.FindById(c.ctx, message.Id)
		if err != nil && err != model.ErrNotFound {
			c.Logger.Error(err)
			return err
		}
		if err == nil {
			if msg.KefuId != message.KefuId {
				//已经被修改，需要即使修改
				message.KefuId = msg.KefuId
			}
		}
	}
	//发送前，务必检查 send 通道关闭了没有
	select {
	case c.send <- message:
		// 消息成功发送
	default:
		// 连接忙碌，无法发送消息
	}
	return nil
}

// SystemErrorHandler 给send chan发送系统错误消息，如果涉及写入和chan的错误，请勿使用
func (c *Client) SystemErrorHandler(kefuId, visitorId int64) {
	c.Send(NewSystemMessage(ErrSystem, kefuId, visitorId))
}

// SendPresetMessage 发送预设的消息
func (c *Client) SendPresetMessage(msgKey string, kefuId, visitorId int64) error {
	if msgKey = strings.TrimSpace(msgKey); msgKey == "" {
		return nil
	}

	msgList, err := c.hub.SvcCtx.MessageConfigModel.FindsByMsgkeyProjectId(c.ctx, msgKey, c.ProjectId)
	if err != nil {
		c.Logger.Error(err)
		c.SystemErrorHandler(kefuId, visitorId)
		return err
	}
	// 如果存在，则发送
	for _, msg := range msgList {
		if content := strings.TrimSpace(msg.Content); content != "" {
			welcomeMsg := c.NewPublicMessage(content, kefuId, visitorId)
			c.Send(welcomeMsg)
		}
	}
	return nil
}

// 获取第一条信息 token 做鉴权，以及从中获取 id 值
// 如果err!=nil，表示内部错误
func (c *Client) Identify() (ok bool, id int64, err error) {
	//读取第一条信息，即 token
	_, token, err := c.conn.ReadMessage()
	if err != nil {
		c.Logger.Error("error reading the first message: ", err)
		err = fmt.Errorf("error reading the first message: %v", err)
		return
	}

	//鉴权
	claims, ok := jwtx.ValidateToken(string(token), c.hub.SvcCtx.Config.VisitorAuthSecret)
	if !ok {
		//如果鉴权失败
		return
	}
	//提取token信息
	id, ok = jwtx.GetId(claims)
	return
}

func (c *Client) IdentifyNotice(ok bool) error {
	body := "no"
	if ok {
		body = "yes"
	}
	if err := c.conn.WriteMessage(websocket.TextMessage, []byte(body)); err != nil {
		c.Logger.Error(err)
		return fmt.Errorf("conn.WriteMessage error: %s", err)
	}
	return nil
}

func (c *Client) NewPublicMessage(content string, kefuId, visitorId int64) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     content,
		ContentType: TextContentType,
		VisitorId:   visitorId,
		KefuId:      kefuId,
		ProjectId:   c.ProjectId,
		MsgType:     PublicMsgType,
		Status:      0,
		IsRead:      0,
		IsSend:      0,
	}
}
