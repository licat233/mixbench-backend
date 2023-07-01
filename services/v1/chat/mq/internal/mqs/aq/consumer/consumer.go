/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-02 10:23:23
 * @Description: licat233@gmail.com
 */
package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mixbench/services/v1/chat/model"
	"mixbench/services/v1/chat/mq/internal/svc"
	"mixbench/services/v1/chat/mq/types"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/service"
)

/**
 * consumer 从redis中取出task进行处理
 */
type AsynqConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	Server *asynq.Server
}

func NewAsynqConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *AsynqConsumer {
	return &AsynqConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
		Server: asynq.NewServer(svcCtx.RedisClientOpt, asynq.Config{Concurrency: 10}),
	}
}

var _ service.Service = &AsynqConsumer{}

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
	case types.MessageTypename:
		var msg model.Message
		if err := json.Unmarshal(t.Payload(), &msg); err != nil {
			return err
		}
		return p.LuckydrawResultHandler(msg)
	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
}

func (p *AsynqConsumer) LuckydrawResultHandler(msg model.Message) (err error) {
	//todo: logic
	return
}
