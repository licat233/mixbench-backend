/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:55:15
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailLogic {
	return &GetUserDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户详细
func (l *GetUserDetailLogic) GetUserDetail(in *luckydraw_pb.GetUserDetailReq) (*luckydraw_pb.GetUserDetailResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("not found user"))
	}

	userActivityLogic := NewGetUserActivitiesLogic(l.ctx, l.svcCtx)

	resp, err := userActivityLogic.GetUserActivities(&luckydraw_pb.GetUserActivitiesReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, err
	}

	return &luckydraw_pb.GetUserDetailResp{
		UserDetail: &luckydraw_pb.UserDetail{
			User:           tools.MdUserToPbUser(user),
			UserActivities: resp.UserActivities,
		},
	}, nil
}
