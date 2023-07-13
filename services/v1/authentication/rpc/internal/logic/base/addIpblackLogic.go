package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIpblackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIpblackLogic {
	return &AddIpblackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加IP黑名单
func (l *AddIpblackLogic) AddIpblack(in *authentication_pb.AddIpblackReq) (*authentication_pb.AddIpblackResp, error) {
	data := &model.Ipblack{
		Ip: in.Ip,
	}
	if _, err := l.svcCtx.IpblackModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.AddIpblackResp{}, nil
}
