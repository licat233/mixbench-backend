package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"mixbench/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminerListLogic {
	return &GetAdminerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminerListLogic) GetAdminerList(in *admin_pb.AdminerListReq) (*admin_pb.GetAdminerListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.AdminerModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbAdminerToMdAdminer(in.Adminer))
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	// for _, adminer := range list {
	// 	adminer.Password = ""
	// }

	return &admin_pb.GetAdminerListResp{
		Data:  mdAdminerListToPbAdminerList(list),
		Total: total,
	}, nil
}
