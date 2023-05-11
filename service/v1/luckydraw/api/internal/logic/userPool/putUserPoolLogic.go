package userPool

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/luckydraw/api/internal/svc"
	"mixbench/service/v1/luckydraw/api/internal/types"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutUserPoolLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPutUserPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutUserPoolLogic {
	return &PutUserPoolLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutUserPoolLogic) PutUserPool(req *types.PutUserPoolReq) (any, error) {
	in := &luckydraw_pb.PutUserPoolReq{
		Id:         req.Id,
		Uuid:       req.Uuid,
		ActivityId: req.ActivityId,
		UserId:     req.UserId,
		Status:     req.Status,
		Score:      req.Score,
	}
	_, err := l.svcCtx.LuckydrawRpc.PutUserPool(l.ctx, in)
	if err != nil {
		//若rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}

	return respx.DefaultStateResp(nil)
}
