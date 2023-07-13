package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutIpblackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutIpblackLogic {
	return &PutIpblackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新IP黑名单
func (l *PutIpblackLogic) PutIpblack(in *authentication_pb.PutIpblackReq) (*authentication_pb.PutIpblackResp, error) {
	data := &model.Ipblack{
		Id: in.Id,
		Ip: in.Ip,
	}

	if err := l.svcCtx.IpblackModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.PutIpblackResp{}, nil
}
