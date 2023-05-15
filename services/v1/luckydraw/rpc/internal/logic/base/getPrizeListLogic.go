package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeListLogic {
	return &GetPrizeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取奖品名单列表
func (l *GetPrizeListLogic) GetPrizeList(in *luckydraw_pb.GetPrizeListReq) (*luckydraw_pb.GetPrizeListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.PrizeModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbPrizeToMdPrize(in.Prize))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizeListResp{
		Prizes: dataconv.MdPrizes2PbPrizes(list),
		Total:  total,
	}, nil
}
