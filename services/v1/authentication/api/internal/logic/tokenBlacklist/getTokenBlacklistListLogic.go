package tokenBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/logic/dataconv"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTokenBlacklistListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTokenBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTokenBlacklistListLogic {
	return &GetTokenBlacklistListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTokenBlacklistListLogic) GetTokenBlacklistList(req *types.GetTokenBlacklistListReq) (any, error) {
	in := &authentication_pb.GetTokenBlacklistListReq{
		ListReq: &authentication_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		TokenBlacklist: &authentication_pb.TokenBlacklist{
			Id:       req.Id,
			Token:    req.Token,
			Platform: req.Platform,
			Ip:       req.Ip,
		},
	}
	rpcResp, err := l.svcCtx.AuthenticationRpc.GetTokenBlacklistList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.TokenBlacklists
	data := dataconv.PbTokenBlacklists2ApiTokenBlacklists(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
