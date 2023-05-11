package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLuckyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLuckyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLuckyListLogic {
	return &GetLuckyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取中奖名单列表
func (l *GetLuckyListLogic) GetLuckyList(in *luckydraw_pb.GetLuckyListReq) (*luckydraw_pb.GetLuckyListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.LuckyModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbLuckyToMdLucky(in.Lucky))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetLuckyListResp{
		Luckies: dataconv.MdLuckies2PbLuckies(list),
		Total:   total,
	}, nil
}
