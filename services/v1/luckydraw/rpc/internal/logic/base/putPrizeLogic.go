package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPrizeLogic {
	return &PutPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新奖品名单
func (l *PutPrizeLogic) PutPrize(in *luckydraw_pb.PutPrizeReq) (*luckydraw_pb.PutPrizeResp, error) {
	data := &model.Prize{
		Id:     in.Id,
		Uuid:   in.Uuid,
		Name:   in.Name,
		Image:  in.Image,
		Price:  in.Price,
		Status: in.Status,
	}

	if err := l.svcCtx.PrizeModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutPrizeResp{}, nil
}
