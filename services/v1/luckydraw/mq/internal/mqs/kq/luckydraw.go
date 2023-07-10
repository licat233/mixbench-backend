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
	"fmt"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/mq/internal/svc"
	"mixbench/services/v1/luckydraw/mq/types"

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
	var msg types.LuckydrawResultMsg
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
func (l *GoodluckStatusMq) execService(message types.LuckydrawResultMsg) (err error) {
	if message.IsWin {
		return
	}

	//奖励积分
	var awardScore int64 = 10

	user, err := l.svcCtx.UserModel.FindById(l.ctx, message.UserId)
	if err != nil && err != model.ErrNotFound {
		err = fmt.Errorf("error finding user, err: %+v", err)
		return
	}
	if err == model.ErrNotFound {
		err = fmt.Errorf("the user %v is not found", message.UserId)
		return
	}

	user.Score += awardScore
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		err = fmt.Errorf("error updating user, err: %+v", err)
		return
	}
	return
}
