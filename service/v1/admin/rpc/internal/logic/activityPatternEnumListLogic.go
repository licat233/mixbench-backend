/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 09:37:17
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityPatternEnumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewActivityPatternEnumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityPatternEnumListLogic {
	return &ActivityPatternEnumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ActivityPatternEnumListLogic) ActivityPatternEnumList(in *admin_pb.NilReq) (*admin_pb.EnumListResp, error) {
	mdList := l.svcCtx.ActivityModel.PatternEnumList()
	return &admin_pb.EnumListResp{
		Data: mdEnumListToPbEnumList(mdList),
	}, nil
}
