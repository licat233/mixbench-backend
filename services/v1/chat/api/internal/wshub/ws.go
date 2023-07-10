package wshub

import (
	"bytes"
	"errors"
	"fmt"
	"mixbench/common/utils"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

func NewWsConn(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	if !utils.IsWebsocketRequest(r) {
		// 不是 WebSocket 请求
		return nil, errors.New("the current request is not a WebSocket request")
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return conn, err
	}
	if err := conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		return conn, fmt.Errorf("SetWriteDeadline error: %s", err)
	}
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	return conn, nil
}

// 关闭ws.conn
func CloseWsConn(conn *websocket.Conn) {
	isClose := isConnectionClosed(conn)
	if !isClose {
		conn.Close()
	}
}

// 判断 ws.conn是否已经关闭
func isConnectionClosed(wsConn *websocket.Conn) bool {
	if wsConn == nil {
		return true
	}
	_, _, err := wsConn.ReadMessage()
	return err != nil

	// if err != nil {
	// 	// 判断是否为预期的关闭错误
	// 	if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
	// 		return true
	// 	} else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	// 		// 是否为意外的关闭错误
	// 		// logx.Error("websocket conn.close error: %v", err)
	// 		return true
	// 	}
	// 	return true // 说明有问题，全部判定为已关闭
	// }

	// return false
}

// 预处理消息内容
func preprocessMessageContent(content []byte) []byte {
	return bytes.TrimSpace(bytes.Replace(content, newline, space, -1))
}
