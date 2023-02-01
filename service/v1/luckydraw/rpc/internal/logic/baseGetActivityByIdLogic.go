/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:28:18
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetActivityByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetActivityByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetActivityByIdLogic {
	return &BaseGetActivityByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetActivityByIdLogic) BaseGetActivityById(in *luckydraw_pb.GetActivityByIdReq) (*luckydraw_pb.GetActivityByIdResp, error) {
	data, err := l.svcCtx.ActivityModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	// if err == model.ErrNotFound {
	// 	return nil, errorx.ExtRpcErr(errors.New("the activity is not found"))
	// }
	return &luckydraw_pb.GetActivityByIdResp{
		Activity: tools.MdActivityToPbActivity(data),
	}, nil
}
