package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerLogic {
	return &GetLuckerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerLogic) GetLucker(in *admin_pb.ItemReq) (*admin_pb.GetLuckerResp, error) {
	res, err := l.svcCtx.LuckerModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &admin_pb.GetLuckerResp{
		Lucker: mdLuckerToPbLucker(res),
	}, nil
}
