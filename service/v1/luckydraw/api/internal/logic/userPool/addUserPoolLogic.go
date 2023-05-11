package userPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserPoolLogic {
	return &AddUserPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserPoolLogic) AddUserPool(req *types.AddUserPoolReq) (any, error) {
	in := &luckydraw_pb.AddUserPoolReq{
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Status:     req.Status,
		Score:      req.Score,
	}
	_, err := l.svcCtx.LuckydrawRpc.AddUserPool(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
