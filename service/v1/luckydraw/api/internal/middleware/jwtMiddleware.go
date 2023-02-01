/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:33
 * @LastEditors: licat
 * @LastEditTime: 2022-12-30 11:58:01
 * @Description: licat233@gmail.com
 */
package middleware

import (
	"context"
	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/utils"
	"mixbench/service/v1/luckydraw/api/internal/config"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JwtMiddleware struct {
	ctx    context.Context
	Config config.Config
}

func NewJwtMiddleware(c config.Config) *JwtMiddleware {
	return &JwtMiddleware{
		ctx:    context.Background(),
		Config: c,
	}
}

func (m *JwtMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(jwtx.ReqHeaderKey)
		if jwtToken == "" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		//解析 token
		claims, err := jwtx.ParseToken(jwtToken, m.Config.JwtAuth.AccessSecret)
		if err != nil {
			logx.WithContext(m.ctx).Error(err)
			httpx.Error(w, errorx.InternalError(err))
			return
		}

		//获取请求设备
		reqPlatform, err := utils.RemotePlatform(r)
		if err != nil {
			httpx.Error(w, errorx.ErrPlatform)
			return
		}

		jwtPlatform := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtPlatform, *claims)
		//设备 校验
		if reqPlatform != jwtPlatform {
			httpx.Error(w, errorx.ErrPlatform)
			return
		}

		jwtIp := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtIp, *claims)
		//获取请求ip
		reqIp := utils.RemoteIp(r)
		//ip 校验
		if reqIp != jwtIp {
			httpx.Error(w, errorx.ErrIp)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
