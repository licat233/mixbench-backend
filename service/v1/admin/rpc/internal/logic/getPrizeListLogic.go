package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeListLogic {
	return &GetPrizeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeListLogic) GetPrizeList(in *admin_pb.PrizeListReq) (*admin_pb.GetPrizeListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizeModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbPrizeToMdPrize(in.Prize))
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizeListResp{
		Data:  mdPrizeListToPbPrizeList(list),
		Total: total,
	}, nil
}
