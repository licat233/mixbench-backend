/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:34:49
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetPrizePoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetPrizePoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetPrizePoolListLogic {
	return &BaseGetPrizePoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetPrizePoolListLogic) BaseGetPrizePoolList(in *luckydraw_pb.GetPrizePoolListReq) (*luckydraw_pb.GetPrizePoolListResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.PrizePoolModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, tools.PbPrizePoolToMdPrizePool(in.PrizePool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetPrizePoolListResp{
		PrizePools: tools.MdPrizePoolsToPbPrizePools(list),
		Total:      total,
	}, nil
}
