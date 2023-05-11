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
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/mq/internal/svc"
	"mixbench/service/v1/luckydraw/mq/types"

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
	case types.TypeLuckydrawResult:
		var msg types.LuckydrawResultMsg
		if err := json.Unmarshal(t.Payload(), &msg); err != nil {
			return err
		}
		return p.LuckydrawResultHandler(msg)
	default:
		return fmt.Errorf("unexpected task type: %s", t.Type())
	}
}

// 如果没有中奖，那就给予10个积分奖励
// 如果中奖了，不给额外的奖励
func (p *AsynqConsumer) LuckydrawResultHandler(msg types.LuckydrawResultMsg) (err error) {
	if msg.IsWin {
		return
	}

	//奖励积分
	var awardScore int64 = 10

	user, err := p.svcCtx.UserModel.FindOne(p.ctx, msg.UserId)
	if err != nil && err != model.ErrNotFound {
		err = fmt.Errorf("find user error, err: %+v", err)
		return
	}
	if err == model.ErrNotFound {
		err = fmt.Errorf("user %v not found", msg.UserId)
		return
	}
	user.Score += awardScore
	err = p.svcCtx.UserModel.Update(p.ctx, user)
	if err != nil {
		err = fmt.Errorf("update user error, err: %+v", err)
		return
	}
	return
}
