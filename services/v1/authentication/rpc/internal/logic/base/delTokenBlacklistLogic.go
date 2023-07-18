package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTokenBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTokenBlacklistLogic {
	return &DelTokenBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除token黑名单
func (l *DelTokenBlacklistLogic) DelTokenBlacklist(in *authentication_pb.DelTokenBlacklistReq) (*authentication_pb.DelTokenBlacklistResp, error) {
	if err := l.svcCtx.TokenBlacklistModel.Delete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.DelTokenBlacklistResp{}, nil
}
