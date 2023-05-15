package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/logic/dataconv"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityListLogic {
	return &GetActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActivityListLogic) GetActivityList(req *types.GetActivityListReq) (any, error) {
	in := &luckydraw_pb.GetActivityListReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Activity: &luckydraw_pb.Activity{
			Id:        req.Id,
			Uuid:      req.Uuid,
			Name:      req.Name,
			ProjectId: req.ProjectId,
			Status:    req.Status,
			Pattern:   req.Pattern,
			MaxCount:  req.MaxCount,
			MaxPrizes: req.MaxPrizes,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		},
	}
	rpcResp, err := l.svcCtx.LuckydrawRpc.GetActivityList(l.ctx, in)
	if err != nil {
		//rpc的错误已经包装过了，无需再处理，直接返回即可
		return nil, err
	}
	pbList := rpcResp.Activities
	data := dataconv.PbActivities2ApiActivities(pbList)

	return respx.DefaultListResp(data, rpcResp.Total, req.PageSize, req.Page, nil)
}
