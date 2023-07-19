package middleware

import (
	"context"
	"errors"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/redisx"
	"mixbench/common/utils"
	"mixbench/services/v1/admin/rpc/admin_pb"
	adminClient "mixbench/services/v1/admin/rpc/client"
	"mixbench/services/v1/authentication/api/internal/config"
	"net/http"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	ctx         context.Context
	Config      config.Config
	AdminRpc    *adminClient.AdminRpc
	RedisClient *redis.Redis
}

func NewAuthMiddleware(c config.Config, RedisClient *redis.Redis, adminRpc *adminClient.AdminRpc) *AuthMiddleware {
	return &AuthMiddleware{
		ctx:         context.Background(),
		Config:      c,
		AdminRpc:    adminRpc,
		RedisClient: RedisClient,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := m.AuthHandler(w, r)
		if err != nil {
			httpx.ErrorCtx(m.ctx, w, err)
			return
		}
		next(w, r)
	}
}

func (m *AuthMiddleware) AuthHandler(w http.ResponseWriter, r *http.Request) error {
	jwtToken := r.Header.Get(jwtx.ReqHeaderTokenKey)
	if jwtToken == "" {
		return errorx.ErrAuth
	}

	//检查token是否已被拉黑
	exists, err := m.RedisClient.Exists(redisx.FormatTokenblacklistKey(jwtToken))
	if err != nil {
		return errorx.InternalError(err)
	}
	if exists {
		return errorx.ErrAuth
	}

	//解析 token
	claims, err := jwtx.ParseToken(jwtToken, m.Config.Auth.AccessSecret)
	if err != nil {
		return errorx.InternalError(err)
	}

	//校验是否过期
	if jwtx.IsExpired(claims) {
		return errorx.ErrAuth
	}

	//获取请求设备
	reqPlatform, err := utils.RemotePlatform(r)
	if err != nil {
		return errorx.ErrPlatform
	}

	jwtPlatform := jwtx.GetClaimsValue(jwtx.JwtPlatformKey, claims)
	//设备 校验
	if reqPlatform != jwtPlatform {
		return errorx.ErrPlatform
	}

	jwtIp := jwtx.GetClaimsValue(jwtx.JwtIpKey, claims)
	//获取请求ip
	reqIp := utils.RemoteIp(r)
	//ip 校验
	if reqIp != jwtIp {
		return errorx.ErrIp
	}
	//校验ip是否已经加入黑名单
	exists, err = m.RedisClient.Exists(redisx.FormatIpblacklistKey(reqIp))
	if err != nil {
		return errorx.InternalError(err)
	}
	if exists {
		return errorx.ErrAuth
	}

	//检查用户状态
	uid, exist := jwtx.GetInt64(claims, jwtx.JwtUidKey)
	if !exist || uid < 1 {
		return errorx.ErrAuth
	}

	resp, err := m.AdminRpc.GetAdminer(m.ctx, &admin_pb.GetAdminerReq{
		Id: uid,
	})
	if err != nil {
		return err
	}
	if resp.Adminer == nil {
		return errorx.AuthError("user does not exist")
	}
	adminer := resp.Adminer
	if adminer.Status == 0 {
		return errorx.AuthError(errors.New("当前用户不可用"))
	}
	return nil
}
