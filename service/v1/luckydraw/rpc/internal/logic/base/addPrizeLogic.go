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

type AddPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPrizeLogic {
	return &AddPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加奖品名单
func (l *AddPrizeLogic) AddPrize(in *luckydraw_pb.AddPrizeReq) (*luckydraw_pb.AddPrizeResp, error) {
	data := &model.Prize{
		Uuid:   uniqueid.NewUUID(),
		Name:   in.Name,
		Image:  in.Image,
		Price:  in.Price,
		Status: in.Status,
	}
	_, err := l.svcCtx.PrizeModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddPrizeResp{}, nil
}
