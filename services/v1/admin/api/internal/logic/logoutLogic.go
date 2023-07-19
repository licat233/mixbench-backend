/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 12:23:43
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"net/http"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/services/v1/admin/api/internal/svc"
	"mixbench/services/v1/admin/rpc/admin_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	Request *http.Request
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *LogoutLogic {
	return &LogoutLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		Request: r,
	}
}

func (l *LogoutLogic) Logout() (any, error) {
	token := l.Request.Header.Get(jwtx.ReqHeaderTokenKey)
	if token == "" {
		return nil, errorx.ExternalError("token not exists")
	}

	_, err := l.svcCtx.AdminRpc.Logout(l.ctx, &admin_pb.LogoutReq{
		Token: token,
	})
	return respx.DefaultStatusResp(err)
}
