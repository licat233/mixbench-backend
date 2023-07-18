package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/logic/dataconv"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutTokenBlacklistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutTokenBlacklistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutTokenBlacklistLogic {
	return &PutTokenBlacklistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新token黑名单
func (l *PutTokenBlacklistLogic) PutTokenBlacklist(in *authentication_pb.PutTokenBlacklistReq) (*authentication_pb.PutTokenBlacklistResp, error) {
	data := dataconv.PbTokenBlacklistToMdTokenBlacklist(in.TokenBlacklist)
	if err := l.svcCtx.TokenBlacklistModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.PutTokenBlacklistResp{}, nil
}
