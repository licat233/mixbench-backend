package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPrizesPoolLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPrizesPoolLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizesPoolLogic {
	return &SearchPrizesPoolLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPrizesPoolLogic) SearchPrizesPool(in *admin_pb.SearchPrizesPoolReq) (*admin_pb.SearchPrizesPoolResp, error) {
	list, err := l.svcCtx.PrizesPoolModel.FindsByActivityId(l.ctx, in.ActivityId)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.SearchPrizesPoolResp{
		Data: mdPrizesPoolListToPbPrizesPoolList(list),
	}, nil
}
