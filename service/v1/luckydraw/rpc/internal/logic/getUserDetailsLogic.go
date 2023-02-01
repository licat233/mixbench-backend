/*
 * @Author: licat
 * @Date: 2023-02-01 14:59:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:45:58
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailsLogic {
	return &GetUserDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户详细列表
func (l *GetUserDetailsLogic) GetUserDetails(in *luckydraw_pb.GetUserDetailsReq) (*luckydraw_pb.GetUserDetailsResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	users, total, err := l.svcCtx.UserModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbUserToMdUser(in.User))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*luckydraw_pb.UserDetail{}
	userActivityLogic := NewGetUserActivitiesLogic(l.ctx, l.svcCtx)
	for _, user := range users {
		resp, err := userActivityLogic.GetUserActivities(&luckydraw_pb.GetUserActivitiesReq{
			UserId: user.Id,
		})
		if err != nil {
			return nil, err
		}
		data = append(data, &luckydraw_pb.UserDetail{
			User:           tools.MdUserToPbUser(user),
			UserActivities: resp.UserActivities,
		})
	}

	return &luckydraw_pb.GetUserDetailsResp{
		UserDetails: data,
		Total:       total,
	}, nil
}
