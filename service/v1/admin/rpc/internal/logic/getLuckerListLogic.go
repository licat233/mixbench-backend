package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckerListLogic {
	return &GetLuckerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckerListLogic) GetLuckerList(in *admin_pb.LuckerListReq) (*admin_pb.GetLuckerListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	if in.Lucker != nil && in.Lucker.Ip != "" && !govalidator.IsIP(in.Lucker.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}

	list, total, err := l.svcCtx.LuckerModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbLuckerToMdLucker(in.Lucker))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetLuckerListResp{
		Data:  mdLuckerListToPbLuckerList(list),
		Total: total,
	}, nil
}
