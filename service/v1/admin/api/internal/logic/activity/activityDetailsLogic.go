package activity

import (
	"context"

	"mixbench/common/respx"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"

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

func (l *ActivityDetailsLogic) ActivityDetails(req *types.ActivityListReq) (any, error) {
	var startTime int64
	var endTime int64
	n := len(req.RangeTime)
	if n == 2 {
		startTime = timeHandler(req.RangeTime[0])
		endTime = timeHandler(req.RangeTime[1])
	}
	// fmt.Printf("\n 开始时间:%v \n 结束时间:%v \n", startTime, endTime)
	resp, err := l.svcCtx.AdminRpc.GetActivityDetails(l.ctx, &admin_pb.ActivityListReq{
		ListReq: &admin_pb.ListReq{
			PageSize: req.PageSize,
			Page:     req.Page,
			Keyword:  req.Keyword,
		},
		Activity: &admin_pb.Activity{
			Name:      req.Name,
			ProjectId: req.ProjectId,
			Status:    req.Status,
			Pattern:   req.Pattern,
			MaxCount:  req.MaxCount,
			MaxPrizes: req.MaxPrizes,
			StartTime: startTime,
			EndTime:   endTime,
		},
	})
	if err != nil {
		return nil, err
	}
	return respx.DefaultListResp(resp.Data, resp.Total, req.PageSize, req.Page, err)
}
