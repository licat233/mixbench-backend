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
	return c, nil
}

func (c *KefuClient) Run() {
	if c.Init() != nil {
		return
	}
	c.wg.Add(2)
	go func() {
		defer c.wg.Done()
		c.writePump()
	}()
	go func() {
		defer c.wg.Done()
		c.readPump(c.payloadHandler)
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

func (c *KefuClient) Init() error {
	// 此时，可能存在需要接待的游客

	//注册online
	OnlineKefu.Set(c.user.GetId(), 0)
	//发送公告信息
	if err := c.sendPresetMessage(PublicMessageKey, c.Kefu.Id, -1); err != nil {
		c.Logger.Error(err)
		return err
	}
	return nil
}

// payloadHandler 用于处理ws消息
// 当返回nil,nil时，表示丢弃
func (c *KefuClient) payloadHandler(payload []byte) (*model.Message, error) {
	//这里的payload为完整的json序列
	if !validatePayload(payload) {
		//丢弃
		return nil, nil
	}
	var msg model.Message
	if err := json.Unmarshal(payload, &msg); err != nil {
		// 不符合规范的 messageBody，丢弃
		return nil, nil
	}
	//不知道发给哪个游客的msg，丢弃
	if msg.VisitorId < 1 {
		return nil, nil
	}
	if !validatePayload([]byte(msg.Content)) {
		//丢弃
		return nil, nil
	}
	message := c.newMessage(msg.Content, msg.VisitorId)
	return message, nil
}

func (c *KefuClient) newMessage(content string, visitorId int64) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     content,
		ContentType: TextContentType,
		VisitorId:   visitorId,
		KefuId:      c.Kefu.Id,
		ProjectId:   c.Kefu.ProjectId,
		MsgType:     KefuMsgType,
		Status:      0,
		IsRead:      0,
		IsSend:      0,
		SendTime:    time.Now().Local(),
		ExpireTime:  time.Now().Local().Add(time.Hour * 24 * 30), //30天
	}
}

// func (c *KefuClient) update() error {
// 	kefu := c.Kefu.Adminer
// 	_, err := c.svcCtx.AdminRpc.PutAdminer(c.ctx, &admin_pb.PutAdminerReq{
// 		Id:           kefu.Id,
// 		Uuid:         kefu.Uuid,
// 		Name:         kefu.Name,
// 		Avatar:       kefu.Avatar,
// 		Access:       kefu.Access,
// 		Passport:     kefu.Passport,
// 		Password:     kefu.Password,
// 		Email:        kefu.Email,
// 		Status:       kefu.Status,
// 		IsSuperAdmin: kefu.IsSuperAdmin,
// 		ProjectId:    kefu.ProjectId,
// 		LoginCount:   kefu.LoginCount,
// 		LastLogin:    kefu.LastLogin,
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
