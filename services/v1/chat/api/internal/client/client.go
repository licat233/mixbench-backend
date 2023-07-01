package client

import (
	"context"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"
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
}

func (c *Client) Close() {
	c.Stop <- true
	close(c.send)
	if !isConnectionClosed(c.conn) {
		c.conn.Close()
	}
}

// 主要用来防止阻塞，用于其它逻辑，请注意区分消息类型
func (c *Client) Send(message *types.Message) {
	select {
	case c.send <- message:
		// 消息成功发送
	default:
		// 连接忙碌，无法发送消息
	}
}

func (c *Client) SystemErrorHandler() {
	c.Send(NewSystemMessage(ErrSystem))
}

func (c *Client) SendWelcomeMessage() error {
	mc, err := c.hub.MessageConfigModel.FindByKey(c.ctx, WelcomeKey)
	if err != nil && err != model.ErrNotFound {
		c.Logger.Error(err)
		c.SystemErrorHandler()
		return err
	}
	// 如果有欢迎信息，则发送欢迎信息
	if err == nil {
		if content := strings.TrimSpace(mc.Content); content != "" {
			welcomeMsg := NewPublicMessage(content)
			c.Send(welcomeMsg)
		}
	}
	return nil
}

func (c *Client) SendPublicMessage() error {
	mc, err := c.hub.MessageConfigModel.FindByKey(c.ctx, PublicKey)
	if err != nil && err != model.ErrNotFound {
		c.Logger.Error(err)
		c.SystemErrorHandler()
		return err
	}
	// 如果有公告，则发送公告信息
	if err == nil {
		if content := strings.TrimSpace(mc.Content); content != "" {
			publicMsg := NewPublicMessage(content)
			c.Send(publicMsg)
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
