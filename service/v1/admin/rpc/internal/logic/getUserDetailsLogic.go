package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

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

func (l *GetUserDetailsLogic) GetUserDetails(in *admin_pb.UserListReq) (*admin_pb.GetUserDetailsListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	users, total, err := l.svcCtx.UserModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbUserToMdUser(in.User))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	data := []*admin_pb.UserDetails{}
	userActivityLogic := NewUserActivityLogic(l.ctx, l.svcCtx)
	for _, user := range users {
		resp, err := userActivityLogic.UserActivity(&admin_pb.UserActivityReq{
			UserId: user.Id,
		})
		if err != nil {
			return nil, err
		}
		data = append(data, &admin_pb.UserDetails{
			User:           mdUserToPbUser(user),
			UserActivities: resp.Data,
		})
	}

	return &admin_pb.GetUserDetailsListResp{
		Data:  data,
		Total: total,
	}, nil
}
