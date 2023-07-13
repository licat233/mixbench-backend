package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mixbench/services/v1/chat/mq/internal/svc"
	"mixbench/services/v1/chat/mq/types"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

/**
 * consumer 从redis中取出task进行处理
 */
type AsynqConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Server *asynq.Server
	logx.Logger
}

func NewAsynqConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqConsumer {
	return &AsynqConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Server: asynq.NewServer(svcCtx.RedisClientOpt, asynq.Config{Concurrency: 10}),
		Logger: logx.WithContext(ctx),
	}
}

var _ service.Service = (*AsynqConsumer)(nil)

func (p *AsynqConsumer) Start() {
	if err := p.Server.Run(asynq.HandlerFunc(p.handler)); err != nil {
		log.Fatal(err)
	}
}

func (p *AsynqConsumer) Stop() {
	p.Server.Stop()
}

func (p *AsynqConsumer) handler(ctx context.Context, t *asynq.Task) error {
	switch t.Type() {
	case types.TaskTypeSendMessage:
		var msg types.SendMessageReq
		if err := json.Unmarshal(t.Payload(), &msg); err != nil {
			return err
		}
		if err := p.SendMessageMq(&msg); err != nil {
			p.Logger.Error(err)
			return err
		}
	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
	return nil
}

func (p *AsynqConsumer) SendMessageMq(msg *types.SendMessageReq) error {
	if msg == nil || msg.Message == nil {
		return nil
	}
	return nil
}
