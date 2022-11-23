package logic

import (
	"context"
	"mixbench/service/v1/luckdraw/api/internal/svc"
	"mixbench/service/v1/luckdraw/api/internal/types"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp any, err error) {
	in := &luckdraw_pb.SearchLuckerReq{
		Passport: req.Keyword,
	}
	res, err := l.svcCtx.LuckdrawRpc.SearchLucker(l.ctx, in)
	if err != nil {
		return nil, err
	}
	return respx.DefaultSingleResp(res, err)
}
