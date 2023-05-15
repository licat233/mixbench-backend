package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutLuckyLogic {
	return &PutLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新中奖名单
func (l *PutLuckyLogic) PutLucky(in *luckydraw_pb.PutLuckyReq) (*luckydraw_pb.PutLuckyResp, error) {
	data := &model.Lucky{
		Id:          in.Id,
		Uuid:        in.Uuid,
		UserId:      in.UserId,
		PrizePoolId: in.PrizePoolId,
		Status:      in.Status,
		LuckDate:    time.Unix(in.LuckDate, 0).Local(),
		Ip:          in.Ip,
	}

	if err := l.svcCtx.LuckyModel.Update(l.ctx, data); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.PutLuckyResp{}, nil
}
