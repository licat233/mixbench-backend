package client

// // client需要实现的接口，暂时不用
// type Pump interface {
// 	Hub() *Hub
// 	Conn() *websocket.Conn
// 	Logger() logx.Logger
// 	Logout()
// 	SendChan() chan *types.Message
// }

// func Write(p Pump) {
// 	ticker := time.NewTicker(pingPeriod)
// 	defer func() {
// 		ticker.Stop()
// 		p.Logout()
// 	}()
// 	hub := p.Hub()
// 	conn := p.Conn()
// 	logger := p.Logger()
// 	send := p.SendChan()
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
