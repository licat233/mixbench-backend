/*
 * @Author: licat
 * @Date: 2023-02-01 15:34:22
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:44:04
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUALuckyDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUALuckyDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUALuckyDetailsLogic {
	return &GetUALuckyDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取某用户在某活动中的中奖详细列表
func (l *GetUALuckyDetailsLogic) GetUALuckyDetails(in *luckydraw_pb.GetUALuckyDetailsReq) (*luckydraw_pb.GetUALuckyDetailsResp, error) {
	aId := in.ActivityId
	if in.ActivityUuid != "" {
		//查找是否存在该活动
		activity, err := l.svcCtx.ActivityModel.FindOneByUuid(l.ctx, in.ActivityUuid)
		if err != nil && err != model.ErrNotFound {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == model.ErrNotFound {
			return nil, errorx.ExtRpcErr(errors.New("the activity is not found"))
		}
		aId = activity.Id
	}

	//查找中奖详细
	list, tatol, err := l.svcCtx.CommonModel.FindsLuckyDetails(l.ctx, &model.LuckyDetailOptions{
		Lucky:      &model.Lucky{UserId: in.UserId},
		ActivityId: aId,
	})
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &luckydraw_pb.GetUALuckyDetailsResp{
		LuckyDetails: tools.ConvertLuckyDetails(list),
		Total:        tatol,
	}, nil
}
