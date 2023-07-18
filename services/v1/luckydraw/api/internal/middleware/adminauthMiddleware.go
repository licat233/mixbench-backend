/*
 * @Author: licat
 * @Date: 2023-02-20 18:27:39
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 18:29:33
 * @Description: licat233@gmail.com
 */
package middleware

import (
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/utils"
	"mixbench/services/v1/luckydraw/api/internal/config"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AdminAuthMiddleware struct {
	Config config.Config
}

func NewAdminAuthMiddleware(c config.Config) *AdminAuthMiddleware {
	return &AdminAuthMiddleware{
		Config: c,
	}
}

func (m *AdminAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(jwtx.ReqHeaderTokenKey)
		if jwtToken == "" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		//解析 token
		claims, err := jwtx.ParseToken(jwtToken, m.Config.Auth.AccessSecret)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, errorx.InternalError(err))
			return
		}

		//获取请求设备
		reqPlatform, err := utils.RemotePlatform(r)
		if err != nil {
			httpx.Error(w, errorx.ErrPlatform)
			return
		}

		jwtPlatform := jwtx.GetClaimsValue(jwtx.JwtPlatformKey, claims)
		//设备 校验
		if reqPlatform != jwtPlatform {
			httpx.Error(w, errorx.ErrPlatform)
			return
		}

		jwtIp := jwtx.GetClaimsValue(jwtx.JwtIpKey, claims)
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
