package userPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolListLogic {
	return &GetUserPoolListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPoolListLogic) GetUserPoolList(req *types.GetUserPoolListReq) (any, error) {
	in := &luckydraw_pb.GetUserPoolListReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		UserPool: &luckydraw_pb.UserPool{
			Id:         req.Id,
			Uuid:       req.Uuid,
			ActivityId: req.ActivityId,
			UserId:     req.UserId,
			Status:     req.Status,
			Score:      req.Score,
		},
	}
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetUserPoolList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.UserPools
	data := dataconv.PbUserPools2ApiUserPools(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
