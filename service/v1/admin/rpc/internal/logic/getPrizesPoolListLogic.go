package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizesPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizesPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizesPoolListLogic {
	return &GetPrizesPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizesPoolListLogic) GetPrizesPoolList(in *admin_pb.PrizesPoolListReq) (*admin_pb.GetPrizesPoolListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizesPoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, pbPrizesPoolToMdPrizesPool(in.PrizesPool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizesPoolListResp{
		Data:  mdPrizesPoolListToPbPrizesPoolList(list),
		Total: total,
	}, nil
}
