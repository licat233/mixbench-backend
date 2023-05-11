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

	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/common/utils"
	"mixbench/service/v1/admin/api/internal/svc"
	"mixbench/service/v1/admin/rpc/admin_pb"

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
	token := l.Request.Header.Get(jwtx.ReqHeaderKey)
	reqIp := utils.RemoteIp(l.Request)
	platform, _ := utils.RemotePlatform(l.Request)

	_, err := l.svcCtx.AdminRpc.Logout(l.ctx, &admin_pb.LogoutReq{
		JwtToken: token,
		Platform: platform,
		Ip:       reqIp,
	})
	return respx.DefaultStateResp(err)
}
