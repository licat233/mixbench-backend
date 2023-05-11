package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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

// 获取奖品级别列表
func (l *GetPrizeGradeListLogic) GetPrizeGradeList(in *luckydraw_pb.GetPrizeGradeListReq) (*luckydraw_pb.GetPrizeGradeListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.PrizeGradeModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbPrizeGradeToMdPrizeGrade(in.PrizeGrade))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizeGradeListResp{
		PrizeGrades: dataconv.MdPrizeGrades2PbPrizeGrades(list),
		Total:       total,
	}, nil
}
