/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:28:47
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

type BaseGetActivityListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetActivityListLogic {
	return &BaseGetActivityListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetActivityListLogic) BaseGetActivityList(in *luckydraw_pb.GetActivityListReq) (*luckydraw_pb.GetActivityListResp, error) {
	in.ListReq = tools.NewListReq(in.ListReq)
	list, total, err := l.svcCtx.ActivityModel.FindList(l.ctx, in.ListReq.PageSize, in.ListReq.Page, in.ListReq.Keyword, tools.PbActivityToMdActivity(in.Activity))

	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetActivityListResp{
		Activities: tools.MdActivitiesToPbActivities(list),
		Total:      total,
	}, nil
}
