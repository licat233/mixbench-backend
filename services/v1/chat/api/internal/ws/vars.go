package ws

import (
	"errors"
	"mixbench/services/v1/chat/model"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512

	// send buffer size
	// bufSize = 256

	// 欢迎信息 key，务必保持与 数据库 中的key一致
	WelcomeMessageKey = "welcome_message"

	// 留言信息 key，务必保持与 数据库 中的key一致
	LeaveMessageKey = "leave_message"

	// 公告信息 key，务必保持与 数据库 中的key一致
	PublicMessageKey = "public_message"

	//消息来源
	VisitorMsgType = 0
	KefuMsgType    = 1
	PublicMsgType  = 2
	SystemMsgType  = 3

	//消息内容类型:0文本，1图片，2表情，3视频, 4错误信息
	TextContentType  = 0
	ImgContentType   = 1
	EmojiContentType = 2
	VideoContentType = 3
	ErrorContentType = 4
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域请求
		return true
	},
}

var ErrSystem = errors.New("server internal error! 伺服器故障! ")

// 将payload转化为message
type payloadHandler func(payload []byte) (*model.Message, error)
