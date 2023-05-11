package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizePoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizePoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizePoolListLogic {
	return &GetPrizePoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动奖品池，某活动中的奖品信息列表
func (l *GetPrizePoolListLogic) GetPrizePoolList(in *luckydraw_pb.GetPrizePoolListReq) (*luckydraw_pb.GetPrizePoolListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.PrizePoolModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbPrizePoolToMdPrizePool(in.PrizePool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizePoolListResp{
		PrizePools: dataconv.MdPrizePools2PbPrizePools(list),
		Total:      total,
	}, nil
}
