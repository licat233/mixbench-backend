package activity

import (
	"context"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/api/internal/types"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"time"

	"mixbench/common/respx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityListLogic {
	return &ActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ActivityListLogic) ActivityList(req *types.ActivityListReq) (any, error) {
	var startTime int64
	var endTime int64
	n := len(req.RangeTime)
	if n == 2 {
		startTime = timeHandler(req.RangeTime[0])
		endTime = timeHandler(req.RangeTime[1])
	}
	// fmt.Printf("\n 开始时间:%v \n 结束时间:%v \n", startTime, endTime)
	resp, err := l.svcCtx.AdminRpc.GetActivityList(l.ctx, &admin_pb.ActivityListReq{
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

func timeHandler(timeStr string) int64 {
	timeLayout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(timeLayout, timeStr, time.Local)
	if err != nil {
		return -1
	}
	return t.Local().Unix()
}
