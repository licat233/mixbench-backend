package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIpblackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIpblackLogic {
	return &DelIpblackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除IP黑名单
func (l *DelIpblackLogic) DelIpblack(in *authentication_pb.DelIpblackReq) (*authentication_pb.DelIpblackResp, error) {
	if err := l.svcCtx.IpblackModel.SoftDelete(l.ctx, in.Id); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.DelIpblackResp{}, nil
}
