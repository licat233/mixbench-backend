package prize

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/errorx"
	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrizeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrizeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrizeAddLogic {
	return &PrizeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrizeAddLogic) PrizeAdd(req *types.PrizeAddReq) (any, error) {
	if req.Name == "" {
		return nil, errorx.ErrInvalidParam
	}
	resp, err := l.svcCtx.AdminRpc.AddPrize(l.ctx, &admin_pb.Prize{
		Name:   req.Name,
		Image:  req.Image,
		Price:  req.Price,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Prize, err)
}
