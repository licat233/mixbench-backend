package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/logic/dataconv"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutIpBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutIpBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutIpBlacklistLogic {
	return &PutIpBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新IP黑名单
func (l *PutIpBlacklistLogic) PutIpBlacklist(in *authentication_pb.PutIpBlacklistReq) (*authentication_pb.PutIpBlacklistResp, error) {
	data := dataconv.PbIpBlacklistToMdIpBlacklist(in.IpBlacklist)
	if err := l.svcCtx.IpBlacklistModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.PutIpBlacklistResp{}, nil
}
