/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2023-02-02 10:39:36
 * @Description: licat233@gmail.com
 */
package kq

import (
	"context"
	"encoding/json"
	"mixbench/services/v1/chat/model"
	"mixbench/services/v1/chat/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageStatusMq struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodluckStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *MessageStatusMq {
	return &MessageStatusMq{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageStatusMq) Consume(_, val string) error {
	var msg model.Message
	if err := json.Unmarshal([]byte(val), &msg); err != nil {
		l.Logger.Errorf("Consume unmarshal error, err: %+v", err)
		return err
	}

	if err := l.execService(msg); err != nil {
		l.Logger.Errorf("GoodluckStatus consume execService error: %v", err)
		return err
	}

	return nil
}

func (l *MessageStatusMq) execService(message model.Message) (err error) {
	//todo: logic
	return
}
