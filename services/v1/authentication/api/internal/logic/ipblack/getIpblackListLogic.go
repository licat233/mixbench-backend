package ipblack

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/authentication/api/internal/logic/dataconv"
	"mixbench/services/v1/authentication/api/internal/svc"
	"mixbench/services/v1/authentication/api/internal/types"
	"mixbench/services/v1/authentication/rpc/authentication_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIpblackListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetIpblackListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIpblackListLogic {
	return &GetIpblackListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetIpblackListLogic) GetIpblackList(req *types.GetIpblackListReq) (any, error) {
	in := &authentication_pb.GetIpblackListReq{
		ListReq: &authentication_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Ipblack: &authentication_pb.Ipblack{
			Id: req.Id,
			Ip: req.Ip,
		},
	}
	rpcResp, err := l.svcCtx.AuthenticationRpc.GetIpblackList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Ipblacks
	data := dataconv.PbIpblacks2ApiIpblacks(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
