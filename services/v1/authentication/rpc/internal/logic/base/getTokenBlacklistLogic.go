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

type GetTokenBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenBlacklistLogic {
	return &GetTokenBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取token黑名单
func (l *GetTokenBlacklistLogic) GetTokenBlacklist(in *authentication_pb.GetTokenBlacklistReq) (*authentication_pb.GetTokenBlacklistResp, error) {
	res, err := l.svcCtx.TokenBlacklistModel.FindById(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GetTokenBlacklistResp{
		TokenBlacklist: dataconv.MdTokenBlacklistToPbTokenBlacklist(res),
	}, nil
}
