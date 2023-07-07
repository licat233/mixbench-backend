package wshub

import (
	"context"
	"mixbench/services/v1/chat/api/internal/types"
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
	send chan *types.Message
	Stop chan bool //关闭信号
	w    http.ResponseWriter
	r    *http.Request
}

func (c *Client) Close() {
	c.Stop <- true
	close(c.send)
	if !isConnectionClosed(c.conn) {
		c.conn.Close()
	}
}

// Send 主要用来防止阻塞，用于其它逻辑，请注意区分消息类型
func (c *Client) Send(message *types.Message) error {
	//发送前，需要校验，该游客信息是否已经被修改了客服号
	if message.MsgType == VisitorMsgType {
		msg, err := c.hub.SvcCtx.MessageModel.FindOne(c.ctx, message.Id)
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
	msgList, _, err := c.hub.SvcCtx.MessageConfigModel.FindList(c.ctx, -1, -1, "", &model.MessageConfig{
		Msgkey: msgKey,
	})
	if err != nil {
		c.Logger.Error(err)
		c.SystemErrorHandler(kefuId, visitorId)
		return err
	}
	// 如果存在，则发送
	for _, msg := range msgList {
		if content := strings.TrimSpace(msg.Content); content != "" {
			welcomeMsg := NewPublicMessage(content, kefuId, visitorId)
			c.Send(welcomeMsg)
		}
	}
	return nil
}

// 丧失了灵活性，暂时不用
// // ws写入消息，
// func (c *Client) WritePump(logout func()) {
// 	ticker := time.NewTicker(pingPeriod)
// 	defer func() {
// 		ticker.Stop()
// 		logout()
// 	}()
// 	hub := c.hub
// 	conn := c.conn
// 	logger := c.Logger
// 	send := c.send
// 	for {
// 		select {
// 		case message, ok := <-send:
// 			conn.SetWriteDeadline(time.Now().Add(writeWait))
// 			if !ok {
// 				// The hub closed the channel.
// 				if err := conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
// 					logger.Error(err)
// 				}
// 				return
// 			}
// 			body, _ := json.Marshal(*message)
// 			if err := conn.WriteMessage(websocket.TextMessage, body); err != nil {
// 				logger.Error(err)
// 				return
// 			}
// 			_ = hub.SendAfterHandler(message)
// 		case <-ticker.C:
// 			conn.SetWriteDeadline(time.Now().Add(writeWait))
// 			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
// 				logger.Error(err)
// 				return
// 			}
// 		}
// 	}
// }
