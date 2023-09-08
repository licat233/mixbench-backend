package ws

import (
	"errors"
	"mixbench/services/v1/chat/api/internal/ws/store"
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

	// 上线信息 key，务必保持与 数据库 中的key一致
	OnlineMessageKey = "online_message"

	//消息类型

	VisitorMsgType = 0 //游客消息
	KefuMsgType    = 1 //客服消息
	PublicMsgType  = 2 //公告消息
	SystemMsgType  = 3 //系统消息
	OnlineMsgType  = 4 //上线消息
	OfflineMsgType = 5 //离线消息

	//消息内容类型:0文本，1图片，2表情，3视频, 4错误信息
	TextContentType  = 0
	ImgContentType   = 1
	EmojiContentType = 2
	VideoContentType = 3
	ErrorContentType = 4

	//无客服接待的游客的redis channel值
	WaiteKefuChannelValue = "waite-kefu-redis-channel-value"

	//是否开启ChatGPT
	UseChatGPT = false
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

type recevierMessageHandler func(message *model.Message) (*model.Message, error)

// 对于kefu来说：map[kefuId]count 客服ID-游客数
var OnlineKefu = store.NewOnlineKefu()

// 对于visitor来说：map[visitorId]kefuId 游客ID-客服ID
// 约定，如果kefuId == 0，表示无客服
var OnlineVisitor = store.NewOnlineVisitor()
