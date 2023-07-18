package ipBlacklist

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/logic/dataconv"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpBlacklistListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpBlacklistListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpBlacklistListLogic {
	return &GetIpBlacklistListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpBlacklistListLogic) GetIpBlacklistList(req *types.GetIpBlacklistListReq) (any, error) {
	in := &authentication_pb.GetIpBlacklistListReq{
		ListReq: &authentication_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		IpBlacklist: &authentication_pb.IpBlacklist{
			Id: req.Id,
			Ip: req.Ip,
		},
	}
	rpcResp, err := l.svcCtx.AuthenticationRpc.GetIpBlacklistList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.IpBlacklists
	data := dataconv.PbIpBlacklists2ApiIpBlacklists(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
