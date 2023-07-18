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
	defer c.Close()
	if c.Init() != nil {
		return
	}
	c.wg.Add(2)
	go func() {
		defer c.wg.Done()
		c.writePump(c.recevierMessageHandler)
	}()
	go func() {
		defer c.wg.Done()
		c.readPump(c.wsPayloadHandler)
	}()
	c.wg.Wait()
	c.done()
}

func (c *KefuClient) done() error {
	//客服离线后的操作
	//此时，如果其对接的游客还在线，应该将他的游客转接给其它客服，而不是放任不管
	//逻辑：转移当前客服的在线游客，分配给其它客服
	//如果没有其它在线客服可分配，则无法分配，无需分配
	return nil
}

func (c *KefuClient) Init() error {
	// 此时，可能存在需要接待的游客

	//查找该客服的在线游客
	ids := OnlineVisitor.FindsByVal(c.Kefu.Id)
	count := len(ids)

	//注册online
	OnlineKefu.Set(c.Kefu.Id, int64(count))

	//发送公告信息
	if err := c.sendPresetMessage(PublicMessageKey, c.Kefu.Id, -1); err != nil {
		return err
	}
	//给该客服分配等待接待的游客
	if err := c.assignVisitor(); err != nil {
		return err
	}

	return nil
}

func (c *KefuClient) assignVisitor() error {
	//需要分配客服的在线的游客
	ids := OnlineVisitor.FindsByVal(0)
	visitors, err := c.svcCtx.VisitorModel.FindsByIds(c.ctx, ids)
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	for _, visitor := range visitors {
		//先更新数据库，确保无误后再更新其它
		if err := c.svcCtx.VisitorModel.Update(c.ctx, visitor); err != nil {
			return err
		}
		//给他们指定客服为当前客服
		OnlineVisitor.Set(visitor.Id, c.Kefu.Id)
		visitor.KefuId = c.Kefu.Id
		//还需要更新visitorClient中的visitor信息，在visitorClient中实现
	}
	//更新其游客数量
	count := OnlineKefu.Get(c.Kefu.Id)
	OnlineKefu.Set(c.Kefu.Id, count+int64(len(ids)))
	return nil
}

// wsPayloadHandler 用于处理ws消息
// 当返回nil,nil时，表示丢弃
func (c *KefuClient) wsPayloadHandler(payload []byte) (*model.Message, error) {
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

func (c *KefuClient) recevierMessageHandler(message *model.Message) error {
	//客服接受到了游客发送过来的消息
	if message.KefuId < 1 && message.MsgType == VisitorMsgType && message.VisitorId > 0 {
		//会有多个客服接收到该信息，谁先拿到算谁的，所以要确认是否已被其他客服拿到了
		kefuId := OnlineVisitor.Get(message.VisitorId)
		if kefuId < 1 {
			OnlineVisitor.Set(message.VisitorId, c.Kefu.Id)
			//检查更新数据库中的数据
			visitor, err := c.svcCtx.VisitorModel.FindById(c.ctx, message.Id)
			if err == model.ErrNotFound {
				return err
			} else if err != nil {
				c.Logger.Error(err)
				return err
			}
			if visitor.KefuId != c.Kefu.Id {
				//更新为当前客服的id
				message.KefuId = c.Kefu.Id
				err = c.svcCtx.VisitorModel.Update(c.ctx, visitor)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
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

// sendHistoryMessages 发送历史聊天记录给当前用户
// 不经过 redis 订阅，直接写入ws
func (c *KefuClient) SendHistoryMessages(visitorId int64) error {
	messages, err := c.svcCtx.MessageModel.FindsByVisitorIdAndKefuId(c.ctx, visitorId, c.Kefu.Id)
	if err != nil {
		c.Logger.Error(err)
		return err
	}
	return c.writeToWebsocket(messages...)
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
