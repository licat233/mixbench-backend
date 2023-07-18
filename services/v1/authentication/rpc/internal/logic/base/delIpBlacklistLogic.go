package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIpBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelIpBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIpBlacklistLogic {
	return &DelIpBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除IP黑名单
func (l *DelIpBlacklistLogic) DelIpBlacklist(in *authentication_pb.DelIpBlacklistReq) (*authentication_pb.DelIpBlacklistResp, error) {
	if err := l.svcCtx.IpBlacklistModel.SoftDelete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.DelIpBlacklistResp{}, nil
}
