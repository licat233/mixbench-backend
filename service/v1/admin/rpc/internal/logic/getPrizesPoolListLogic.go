package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizePoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolListLogic {
	return &GetPrizePoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizePoolListLogic) GetPrizePoolList(in *admin_pb.PrizePoolListReq) (*admin_pb.GetPrizePoolListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizePoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, pbPrizePoolToMdPrizePool(in.PrizePool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizePoolListResp{
		Data:  mdPrizePoolListToPbPrizePoolList(list),
		Total: total,
	}, nil
}
