package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPrizeGradeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizeGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeGradeLogic {
	return &AddPrizeGradeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加奖品级别
func (l *AddPrizeGradeLogic) AddPrizeGrade(in *luckydraw_pb.AddPrizeGradeReq) (*luckydraw_pb.AddPrizeGradeResp, error) {
	data := &model.PrizeGrade{
		Uuid:  uniqueid.NewUUID(),
		Name:  in.Name,
		IsWin: in.IsWin,
	}
	_, err := l.svcCtx.PrizeGradeModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddPrizeGradeResp{}, nil
}
