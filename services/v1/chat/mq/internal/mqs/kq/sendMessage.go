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
	"mixbench/services/v1/chat/mq/internal/svc"
	"mixbench/services/v1/chat/mq/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageMq struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageMq(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageMq {
	return &SendMessageMq{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageMq) Consume(key, val string) error {
	var req types.SendMessageReq
	if err := json.Unmarshal([]byte(val), &req); err != nil {
		l.Logger.Errorf("Consume unmarshal error, err: %+v", err)
		return err
	}
	return l.execService(&req)
}

func (l *SendMessageMq) execService(req *types.SendMessageReq) (err error) {
	//todo: logic
	return
}
