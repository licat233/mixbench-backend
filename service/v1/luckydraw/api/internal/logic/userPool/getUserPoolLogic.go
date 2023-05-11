package userPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolLogic {
	return &GetUserPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserPoolLogic) GetUserPool(req *types.GetUserPoolReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetUserPool(l.ctx, &luckydraw_pb.GetUserPoolReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbUserPool := rpcResp.UserPool
	data := dataconv.PbUserPoolToApiUserPool(pbUserPool)

	return respx.DefaultSingleResp(data, nil)
}
