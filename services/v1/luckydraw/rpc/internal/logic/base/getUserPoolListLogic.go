package baselogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/luckydraw/rpc/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserPoolListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserPoolListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserPoolListLogic {
	return &GetUserPoolListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取活动用户池，可以参与该活动抽奖的用户，及其在该活动中的积分管理列表
func (l *GetUserPoolListLogic) GetUserPoolList(in *luckydraw_pb.GetUserPoolListReq) (*luckydraw_pb.GetUserPoolListResp, error) {
	pageSize, page, keyword := dataconv.ListReqParams(in.ListReq)
	list, total, err := l.svcCtx.UserPoolModel.FindList(l.ctx, pageSize, page, keyword, dataconv.PbUserPoolToMdUserPool(in.UserPool))
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetUserPoolListResp{
		UserPools: dataconv.MdUserPools2PbUserPools(list),
		Total:     total,
	}, nil
}
