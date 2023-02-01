/*
 * @Author: licat
 * @Date: 2023-02-01 14:59:03
 * @LastEditors: licat
 * @LastEditTime: 2023-02-01 15:01:18
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/service/v1/luckydraw/rpc/internal/logic/tools"
	"mixbench/service/v1/luckydraw/rpc/internal/svc"
	"mixbench/service/v1/luckydraw/rpc/luckydraw_pb"

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
func (l *GetActivityPatternEnumsLogic) GetActivityPatternEnums(in *luckydraw_pb.NilReq) (*luckydraw_pb.GetEnumsResp, error) {
	mdList := l.svcCtx.ActivityModel.PatternEnumList()
	return &luckydraw_pb.GetEnumsResp{
		Enums: tools.MdEnumsToPbEnums(mdList),
	}, nil
}
