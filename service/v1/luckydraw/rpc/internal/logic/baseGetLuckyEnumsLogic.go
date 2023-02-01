/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 17:29:17
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

type BaseGetLuckyEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBaseGetLuckyEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BaseGetLuckyEnumsLogic {
	return &BaseGetLuckyEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BaseGetLuckyEnumsLogic) BaseGetLuckyEnums(in *luckydraw_pb.GetLuckyEnumsReq) (*luckydraw_pb.GetEnumsResp, error) {
	list, _, err := l.svcCtx.LuckyModel.FindList(l.ctx, -1, -1, "", nil)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	enums := []*luckydraw_pb.Enum{}
	for _, item := range list {
		enums = append(enums, &luckydraw_pb.Enum{
			Label: item.Uuid,
			Value: item.Id,
		})
	}

	return &luckydraw_pb.GetEnumsResp{
		Enums: enums,
	}, nil
}
