/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:23:51
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/common/uniqueid"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseAddActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseAddActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseAddActivityLogic {
	return &BaseAddActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseAddActivityLogic) BaseAddActivity(in *luckydraw_pb.AddActivityReq) (*luckydraw_pb.AddActivityResp, error) {
	data := &model.Activity{
		Uuid:      uniqueid.NewUUID(),
		Name:      in.Name,
		ProjectId: in.ProjectId,
		Status:    in.Status,
		Pattern:   in.Pattern,
		MaxCount:  in.MaxCount,
		MaxPrizes: in.MaxPrizes,
		StartTime: utils.Int64ToTime(utils.TimestampHandle(in.StartTime)),
		EndTime:   utils.Int64ToTime(utils.TimestampHandle(in.EndTime)),
	}
	result, err := l.svcCtx.ActivityModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	activity, err := l.svcCtx.ActivityModel.FindOne(l.ctx, id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.AddActivityResp{
		Activity: tools.MdActivityToPbActivity(activity),
	}, nil
}
