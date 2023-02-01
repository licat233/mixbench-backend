/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:28:34
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/luckydraw/model"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type BaseGetActivityEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetActivityEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetActivityEnumsLogic {
	return &BaseGetActivityEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetActivityEnumsLogic) BaseGetActivityEnums(in *luckydraw_pb.GetActivityEnumsReq) (*luckydraw_pb.GetEnumsResp, error) {
	var list = []*model.Activity{}
	var err error
	var projectId = in.ParentId
	if projectId == -1 {
		list, _, err = l.svcCtx.ActivityModel.FindList(l.ctx, -1, -1, "", nil)
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	} else {
		list, err = l.svcCtx.ActivityModel.FindsByProjectId(l.ctx, projectId)
		if err != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
	}

	enums := []*luckydraw_pb.Enum{}
	for _, item := range list {
		enums = append(enums, &luckydraw_pb.Enum{
			Label: item.Name,
			Value: item.Id,
		})
	}
	return &luckydraw_pb.GetEnumsResp{
		Enums: enums,
	}, nil
}
