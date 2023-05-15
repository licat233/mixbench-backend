package baselogic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLuckyLogic {
	return &AddLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加中奖名单
func (l *AddLuckyLogic) AddLucky(in *luckydraw_pb.AddLuckyReq) (*luckydraw_pb.AddLuckyResp, error) {
	data := &model.Lucky{
		Uuid:        uniqueid.NewUUID(),
		UserId:      in.UserId,
		PrizePoolId: in.PrizePoolId,
		Status:      in.Status,
		LuckDate:    time.Unix(in.LuckDate, 0).Local(),
		Ip:          in.Ip,
	}
	_, err := l.svcCtx.LuckyModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddLuckyResp{}, nil
}
