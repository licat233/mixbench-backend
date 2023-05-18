package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/services/v1/luckydraw/api/internal/svc"
	"mixbench/services/v1/luckydraw/api/internal/types"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityDetailsLogic {
	return &ActivityDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityDetailsLogic) ActivityDetails(req *types.GetActivityListReq) (any, error) {
	// fmt.Printf("\n 开始时间:%v \n 结束时间:%v \n", startTime, endTime)
	resp, err := l.svcCtx.LuckydrawRpc.GetActivityDetails(l.ctx, &luckydraw_pb.GetActivityDetailsReq{
		ListReq: &luckydraw_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Activity: &luckydraw_pb.Activity{
			Name:      req.Name,
			ProjectId: req.ProjectId,
			Status:    req.Status,
			Pattern:   req.Pattern,
			MaxCount:  req.MaxCount,
			MaxPrizes: req.MaxPrizes,
			StartTime: req.StartTime,
			EndTime:   req.EndTime,
		},
	})
	if err != nil {
		return nil, err
	}
	data := []*types.ActivityDetail{}
	for _, v := range resp.ActivityDetails {
		data = append(data, PbActivityDetailToApiActivityDetail(v))
	}
	return respx.DefaultListResp(data, resp.Total, req.PageSize, req.Page, err)
}
