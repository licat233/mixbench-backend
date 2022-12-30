package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeGradeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeGradeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeGradeListLogic {
	return &GetPrizeGradeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeGradeListLogic) GetPrizeGradeList(in *admin_pb.PrizeGradeListReq) (*admin_pb.GetPrizeGradeListResp, error) {
	in.ListReq = newListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizeGradeModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, pbPrizeGradeToMdPrizeGrade(in.PrizeGrade))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GetPrizeGradeListResp{
		Data:  mdPrizeGradeListToPbPrizeGradeList(list),
		Total: total,
	}, nil
}
