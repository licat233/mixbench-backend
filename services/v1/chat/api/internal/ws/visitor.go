package ws

import (
	"context"
	"errors"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/model"
	"net/http"
	"time"
)

type VisitorInfo struct {
	*model.Visitor
}

var _ user = (*VisitorInfo)(nil)

func (i *VisitorInfo) GetId() int64        { return i.Visitor.Id }
func (i *VisitorInfo) GetUuid() string     { return i.Visitor.Uuid }
func (i *VisitorInfo) GetProjectId() int64 { return i.Visitor.ProjectId }
func (i *VisitorInfo) IsKefu() bool        { return false }
func (i *VisitorInfo) IsVisitor() bool     { return true }

type VisitorClient struct {
	*client
	Visitor *VisitorInfo
}

func NewVisitorClient(ctx context.Context, svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request, visitor *VisitorInfo) (*VisitorClient, error) {
	if visitor == nil || visitor.Uuid == "" || visitor.ProjectId <= 0 {
		return nil, errors.New("visitor field value error")
	}
	client, err := newClient(ctx, svcCtx, w, r, visitor)
	if err != nil {
		return nil, err
	}
	c := &VisitorClient{
		client:  client,
		Visitor: visitor,
	}
	return c, nil
}

func (c *VisitorClient) Run() {
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

func (c *VisitorClient) done() error {
	c.client.Close() //关闭client
	return nil
}

func (c *VisitorClient) Init() error {
	// 查看该游客是否有专属客服接待
	if c.Visitor.KefuId > 0 {
		//判断客服是否在线
		isOnline := OnlineKefu.Exist(c.Visitor.KefuId)
		//如果不在线，
		if !isOnline {
			//如果不在线，则应该给该游客分配其他客服
			newKefuId := c.assignKefu()
			//确保分配到了客服，且跟原来的不一样，否则不改变原kefuId
			if newKefuId > 0 && newKefuId != c.Visitor.KefuId {
				c.Visitor.KefuId = newKefuId
				//更新数据库
				if err := c.update(); err != nil {
					c.Logger.Error(err)
					c.client.Close()
					return err
				}
			}
		}
		//如果在线，则不用管
	}
	// 注册online
	OnlineKefu.Set(c.user.GetId(), 0)
	// 如果该游客有客服接待
	if c.Visitor.KefuId > 0 {
		// 发送欢迎信息
		if err := c.sendPresetMessage(WelcomeMessageKey, -1, c.Visitor.Id); err != nil {
			c.Logger.Error(err)
			return err
		}
	} else {
		// 没有分配接待的客服，该怎样处理？说明没有客服在线，无法分配，此时应该发送留言信息
		// 发送留言信息
		if err := c.sendPresetMessage(LeaveMessageKey, -1, c.Visitor.Id); err != nil {
			c.Logger.Error(err)
			return err
		}
	}

	return nil
}

func (c *VisitorClient) payloadHandler(payload []byte) (*model.Message, error) {
	//这里的payload仅为message.content
	if !validatePayload(payload) {
		//丢弃
		return nil, nil
	}
	// 如果该游客没有客服接待
	if c.Visitor.KefuId < 1 {
		//尝试分配客服
		kefuId := c.assignKefu()
		//如果有客服
		if kefuId > 0 {
			//如果有客服，将其分配给该游客
			c.Visitor.KefuId = kefuId
			//同时更新数据库
			if err := c.update(); err != nil {
				c.Logger.Error(err)
				return nil, err
			}
		}
		//不管有无客服，都不应该丢弃该条信息
	}
	message := c.newMessage(string(payload))
	return message, nil
}

func (c *VisitorClient) update() error {
	return c.svcCtx.VisitorModel.Update(c.ctx, c.Visitor.Visitor)
}

func (c *VisitorClient) newMessage(content string) *model.Message {
	return &model.Message{
		Id:          0,
		Uuid:        uniqueid.NewUUID(),
		Content:     content,
		ContentType: TextContentType,
		VisitorId:   c.Visitor.Id,
		KefuId:      c.Visitor.KefuId,
		ProjectId:   c.Visitor.ProjectId,
		MsgType:     VisitorMsgType,
		Status:      0,
		IsRead:      0,
		IsSend:      0,
		SendTime:    time.Now().Local(),
		ExpireTime:  time.Now().Local().Add(time.Hour * 24 * 30), //30天
	}
}
