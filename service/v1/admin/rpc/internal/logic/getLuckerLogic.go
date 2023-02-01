package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyLogic {
	return &GetLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckyLogic) GetLucky(in *admin_pb.ItemReq) (*admin_pb.GetLuckyResp, error) {
	res, err := l.svcCtx.LuckyModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &admin_pb.GetLuckyResp{
		Lucky: mdLuckyToPbLucky(res),
	}, nil
}
