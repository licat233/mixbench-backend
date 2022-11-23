package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLogoutLogic {
	return &IsLogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否已经注销
func (l *IsLogoutLogic) IsLogout(in *admin_pb.IsLogoutReq) (*admin_pb.IsLogoutResp, error) {
	//优先从redis中查找
	exist, err := l.svcCtx.RedisCli.ExistsCtx(l.ctx, in.JwtToken)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	if exist {
		return &admin_pb.IsLogoutResp{
			Status: exist,
		}, nil
	}

	//再从mysql中查找
	exist, err = l.svcCtx.JwtBlacklistsModel.ExistByToken(l.ctx, in.JwtToken)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.IsLogoutResp{
		Status: exist,
	}, nil
}
