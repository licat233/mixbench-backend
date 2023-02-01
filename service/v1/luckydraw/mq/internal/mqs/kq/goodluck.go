/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 11:59:24
 * @Description: licat233@gmail.com
 */
package kq

import (
	"context"
	"encoding/json"
	"fmt"
	"mixbench/service/v1/luckydraw/mq/internal/svc"
	"mixbench/service/v1/luckydraw/mq/types"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodluckStatusMq struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodluckStatusMq(ctx context.Context, svcCtx *svc.ServiceContext) *GoodluckStatusMq {
	return &GoodluckStatusMq{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodluckStatusMq) Consume(_, val string) error {
	var msg types.LuckdrawResultMsg
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

// 如果没有中奖，那就给予10个积分奖励
// 如果中奖了，不给额外的奖励
func (l *GoodluckStatusMq) execService(message types.LuckdrawResultMsg) (err error) {
	if message.IsWin {
		return
	}

	//奖励积分
	var awardScore int64 = 10

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, message.UserId)
	if err != nil && err != model.ErrNotFound {
		err = fmt.Errorf("find user error, err: %+v", err)
		return
	}
	if err == model.ErrNotFound {
		err = fmt.Errorf("user %v not found", message.UserId)
		return
	}
	user.Score += awardScore
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		err = fmt.Errorf("update user error, err: %+v", err)
		return
	}
	return
}
