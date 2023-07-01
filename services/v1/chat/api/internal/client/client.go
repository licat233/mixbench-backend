package client

import (
	"context"
	"encoding/json"
	"mixbench/services/v1/chat/api/internal/types"
	"time"

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

func (c *Client) Send(message *types.Message) {
	select {
	case c.send <- message:
		// 消息成功发送
	default:
		// 连接忙碌，无法发送消息
	}
}

// ws写入消息，
func (c *Client) WritePump(logout func()) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		logout()
	}()
	hub := c.hub
	conn := c.conn
	logger := c.Logger
	send := c.send
	for {
		select {
		case message, ok := <-send:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				if err := conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					logger.Error(err)
				}
				return
			}
			body, _ := json.Marshal(*message)
			if err := conn.WriteMessage(websocket.TextMessage, body); err != nil {
				logger.Error(err)
				return
			}
			_ = hub.SendAfterHandler(message)
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				logger.Error(err)
				return
			}
		}
	}
}
