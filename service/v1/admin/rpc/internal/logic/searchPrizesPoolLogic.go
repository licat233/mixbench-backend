package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPrizePoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPrizePoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizePoolLogic {
	return &SearchPrizePoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPrizePoolLogic) SearchPrizePool(in *admin_pb.SearchPrizePoolReq) (*admin_pb.SearchPrizePoolResp, error) {
	list, err := l.svcCtx.PrizePoolModel.FindsByActivityId(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.SearchPrizePoolResp{
		Data: mdPrizePoolListToPbPrizePoolList(list),
	}, nil
}
