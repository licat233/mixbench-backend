package ws

import (
	"context"
	"encoding/json"
	"errors"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/model"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type KefuInfo struct {
	*admin_pb.Adminer
}

var _ user = (*KefuInfo)(nil)

func (i *KefuInfo) GetId() int64        { return i.Id }
func (i *KefuInfo) GetUuid() string     { return i.Uuid }
func (i *KefuInfo) GetProjectId() int64 { return i.ProjectId }
func (i *KefuInfo) IsKefu() bool        { return true }
func (i *KefuInfo) IsVisitor() bool     { return false }

type KefuClient struct {
	*client
	Kefu *KefuInfo
}

func NewKefuClient(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request, kefu *KefuInfo) (*KefuClient, error) {
	if kefu == nil || kefu.Uuid == "" || kefu.ProjectId <= 0 {
		return nil, errors.New("kefu field value error")
	}
	client, err := newClient(ctx, svcCtx, w, r, kefu)
	if err != nil {
		return nil, err
	}
	c := &KefuClient{
		client: client,
		Kefu:   kefu,
	}
	c.client.status = c.Init() == nil
	return c, nil
}

func (c *KefuClient) Init() error {
	// 此时，可能存在需要接待的游客

	//注册online
	OnlineKefu.Set(c.user.GetId(), 0)
	return nil
}

func (c *KefuClient) Run() {
	if !c.client.status {
		return
	}
	c.wg.Add(2)
	go func() {
		defer c.wg.Done()
		c.writePump()
	}()
	go func() {
		defer c.wg.Done()
		c.readPump()
	}()
	c.wg.Wait()
	c.done()
}

func (c *KefuClient) done() error {
	c.client.Close() //关闭client
	//此时，如果其对接的游客还在线，应该将他的游客转接给其它客服，而不是放任不管
	//逻辑：转移当前客服的在线游客，分配给其它客服
	//如果没有其它在线客服可分配，则无法分配，无需分配
	return nil
}

// 读取客服通过websocket.Conn发送过来的信息，发布至redis中
func (c *KefuClient) readPump() {
	defer func() {
		close(c.stop)
	}()

	for {
		select {
		case <-c.client.stop:
			return
		default:
			// messageBody 为一个完整的 message 的 json 数据
			_, messageBody, err := c.conn.ReadMessage()
			if err != nil {
				//ws已经关闭
				return
			}

			var msg model.Message
			if err = json.Unmarshal(messageBody, &msg); err != nil {
				// 不符合规范的 messageBody，丢弃
				continue
			}
			//不知道发给哪个游客的msg，丢弃
			if msg.VisitorId < 1 {
				// c.Logger.Error("unvalidtor visitor id: %d", message.VisitorId)
				continue
			}

			message := c.newKefuMessage(msg.Content, msg.VisitorId)

			//改用redis发布方案：
			if err = c.client.sendMessage(message); err != nil {
				c.Logger.Error(err)
				return
			}
		}
	}
}

// 给客服的ws写入消息
func (c *KefuClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		close(c.stop)
	}()

	//发送公告信息
	if err := c.sendPresetMessage(PublicMessageKey, c.Kefu.Id, -1); err != nil {
		c.Logger.Error(err)
		return
	}

	for {
		select {
		case msg, ok := <-c.pubsub.Channel():
			//如果通道已经关闭
			if !ok {
				// 关闭ws.conn
				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
						//正常关闭
						return
					} else {
						//异常关闭
						c.Logger.Error("websocket conn IsUnexpectedCloseError: %v", err)
						return
					}
				}
				// 已关闭，正常退出
				return
			}
			var message model.Message
			if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
				//不合规的消息体。丢弃
				continue
			}
			if err := c.writeMessage(&message); err != nil {
				c.Logger.Error(err)
				return
			}
		case <-c.client.stop:
			//关闭了
			return
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				c.Logger.Error(err)
				return
			}
		}
	}
}

func (c *KefuClient) Update() error {
	kefu := c.Kefu.Adminer
	_, err := c.svcCtx.AdminRpc.PutAdminer(c.ctx, &admin_pb.PutAdminerReq{
		Id:           kefu.Id,
		Uuid:         kefu.Uuid,
		Name:         kefu.Name,
		Avatar:       kefu.Avatar,
		Access:       kefu.Access,
		Passport:     kefu.Passport,
		Password:     kefu.Password,
		Email:        kefu.Email,
		Status:       kefu.Status,
		IsSuperAdmin: kefu.IsSuperAdmin,
		ProjectId:    kefu.ProjectId,
		LoginCount:   kefu.LoginCount,
		LastLogin:    kefu.LastLogin,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *KefuClient) newKefuMessage(content string, visitorId int64) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     content,
		ContentType: TextContentType,
		VisitorId:   visitorId,
		KefuId:      c.Kefu.Id,
		ProjectId:   c.user.GetProjectId(),
		MsgType:     KefuMsgType,
		Status:      0,
		IsRead:      0,
		IsSend:      0,
		SendTime:    time.Now().Local(),
		ExpireTime:  time.Now().Local().Add(time.Hour * 24 * 30), //30天
	}
}
