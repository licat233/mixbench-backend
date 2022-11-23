package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

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

func (l *GetUserDetailLogic) GetUserDetail(in *admin_pb.GetUserDetailReq) (*admin_pb.GetUserDetailResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("not found user"))
	}

	userActivityLogic := NewUserActivityLogic(l.ctx, l.svcCtx)

	resp, err := userActivityLogic.UserActivity(&admin_pb.UserActivityReq{
		UserId: user.Id,
	})
	if err != nil {
		return nil, err
	}

	return &admin_pb.GetUserDetailResp{
		Data: &admin_pb.UserDetails{
			User:           mdUserToPbUser(user),
			UserActivities: resp.Data,
		},
	}, nil
}
