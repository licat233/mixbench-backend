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

type GetIpBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIpBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpBlacklistLogic {
	return &GetIpBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取IP黑名单
func (l *GetIpBlacklistLogic) GetIpBlacklist(in *authentication_pb.GetIpBlacklistReq) (*authentication_pb.GetIpBlacklistResp, error) {
	res, err := l.svcCtx.IpBlacklistModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GetIpBlacklistResp{
		IpBlacklist: dataconv.MdIpBlacklistToPbIpBlacklist(res),
	}, nil
}
