package client

import (
	"bytes"

	"github.com/gorilla/websocket"
)

// 判断 ws.conn是否已经关闭
func isConnectionClosed(wsConn *websocket.Conn) bool {
	if wsConn == nil {
		return true
	}
	_, _, err := wsConn.ReadMessage()
	//如果 err == nil，说明conn还未关闭
	return err != nil
	// if err != nil {
	// 	// // 判断是否为预期的关闭错误
	// 	// if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
	// 	// 	return true
	// 	// } else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	// 	// 	// 是否为意外的关闭错误
	// 	// 	logx.Error("websocket conn.close error: %v", err)
	// 	// 	return true
	// 	// }
	// 	return true // 说明有问题，全部判定为已关闭
	// }
	// return false
}

func preprocessMessageContent(content []byte) []byte {
	return bytes.TrimSpace(bytes.Replace(content, newline, space, -1))
}
