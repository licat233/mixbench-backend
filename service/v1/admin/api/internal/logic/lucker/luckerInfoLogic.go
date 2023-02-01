package lucky

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type LuckyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLuckyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LuckyInfoLogic {
	return &LuckyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LuckyInfoLogic) LuckyInfo(req *types.LuckyGetReq) (any, error) {
	resp, err := l.svcCtx.AdminRpc.GetLucky(l.ctx, &admin_pb.ItemReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(resp.Lucky, err)
}
