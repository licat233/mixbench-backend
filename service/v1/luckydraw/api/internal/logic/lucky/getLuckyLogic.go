package lucky

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyLogic {
	return &GetLuckyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLuckyLogic) GetLucky(req *types.GetLuckyReq) (any, error) {
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetLucky(l.ctx, &luckydraw_pb.GetLuckyReq{
		Id: req.Id,
	})
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbLucky := rpcResp.Lucky
	data := dataconv.PbLuckyToApiLucky(pbLucky)

	return respx.DefaultSingleResp(data, nil)
}
