/*
 * @Author: licat
 * @Date: 2023-02-01 14:59:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 17:45:23
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"

	"mixbench/services/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActivityPatternEnumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActivityPatternEnumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActivityPatternEnumsLogic {
	return &GetActivityPatternEnumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 活动规则枚举
func (l *GetActivityPatternEnumsLogic) GetActivityPatternEnums(in *luckydraw_pb.NilReq) (*luckydraw_pb.Enums, error) {
	mdList := l.svcCtx.ActivityModel.PatternEnumList()
	return &luckydraw_pb.Enums{
		Enums: tools.MdEnumsToPbEnums(mdList),
	}, nil
}
