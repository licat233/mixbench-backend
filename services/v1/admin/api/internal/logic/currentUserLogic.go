/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 11:21:05
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"

	"mixbench/common/ctxdata"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CurrentUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCurrentUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CurrentUserLogic {
	return &CurrentUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CurrentUserLogic) CurrentUser() (any, error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	resp, err := l.svcCtx.AdminRpc.GetAdminer(l.ctx, &admin_pb.GetAdminerReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	resp.Adminer.Password = ""
	return respx.NewData(resp.Adminer), nil
}
