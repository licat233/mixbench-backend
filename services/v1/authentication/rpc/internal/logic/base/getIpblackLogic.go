package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/logic/dataconv"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIpblackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackLogic {
	return &GetIpblackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取IP黑名单
func (l *GetIpblackLogic) GetIpblack(in *authentication_pb.GetIpblackReq) (*authentication_pb.GetIpblackResp, error) {
	res, err := l.svcCtx.IpblackModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &authentication_pb.GetIpblackResp{
		Ipblack: dataconv.MdIpblackToPbIpblack(res),
	}, nil
}
