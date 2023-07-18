package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/logic/dataconv"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenBlacklistListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTokenBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenBlacklistListLogic {
	return &GetTokenBlacklistListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取token黑名单列表
func (l *GetTokenBlacklistListLogic) GetTokenBlacklistList(in *authentication_pb.GetTokenBlacklistListReq) (*authentication_pb.GetTokenBlacklistListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.TokenBlacklistModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbTokenBlacklistToMdTokenBlacklist(in.TokenBlacklist))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GetTokenBlacklistListResp{
		TokenBlacklists: dataconv.MdTokenBlacklists2PbTokenBlacklists(list),
		Total:           total,
	}, nil
}
