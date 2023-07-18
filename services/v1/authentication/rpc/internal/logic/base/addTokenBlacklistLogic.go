package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/model"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTokenBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTokenBlacklistLogic {
	return &AddTokenBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加token黑名单
func (l *AddTokenBlacklistLogic) AddTokenBlacklist(in *authentication_pb.AddTokenBlacklistReq) (*authentication_pb.AddTokenBlacklistResp, error) {
	data := &model.TokenBlacklist{
		Token:    in.Token,
		Platform: in.Platform,
		Ip:       in.Ip,
	}
	if _, err := l.svcCtx.TokenBlacklistModel.Insert(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.AddTokenBlacklistResp{}, nil
}
