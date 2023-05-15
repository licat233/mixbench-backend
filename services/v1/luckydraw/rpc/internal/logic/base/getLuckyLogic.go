package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/model"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyLogic {
	return &GetLuckyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取中奖名单
func (l *GetLuckyLogic) GetLucky(in *luckydraw_pb.GetLuckyReq) (*luckydraw_pb.GetLuckyResp, error) {
	res, err := l.svcCtx.LuckyModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, nil
	}

	return &luckydraw_pb.GetLuckyResp{
		Lucky: dataconv.MdLuckyToPbLucky(res),
	}, nil
}
