package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIpBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIpBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIpBlacklistLogic {
	return &AddIpBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加IP黑名单
func (l *AddIpBlacklistLogic) AddIpBlacklist(in *authentication_pb.AddIpBlacklistReq) (*authentication_pb.AddIpBlacklistResp, error) {
	data := &model.IpBlacklist{
		Ip: in.Ip,
	}
	if _, err := l.svcCtx.IpBlacklistModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.AddIpBlacklistResp{}, nil
}
