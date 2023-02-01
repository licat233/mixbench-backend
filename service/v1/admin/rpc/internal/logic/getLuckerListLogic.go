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

type GetLuckyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyListLogic {
	return &GetLuckyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLuckyListLogic) GetLuckyList(in *admin_pb.LuckyListReq) (*admin_pb.GetLuckyListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	if in.Lucky != nil && in.Lucky.Ip != "" && !govalidator.IsIP(in.Lucky.Ip) {
		return nil, errorx.ExtRpcErr(errors.New("ip field invalid"))
	}

	list, total, err := l.svcCtx.LuckyModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbLuckyToMdLucky(in.Lucky))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetLuckyListResp{
		Data:  mdLuckyListToPbLuckyList(list),
		Total: total,
	}, nil
}
