package middleware

import (
	"context"
	"errors"
	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/utils"
	"mixbench/service/v1/admin/api/internal/config"
	"mixbench/service/v1/model"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JwtMiddleware struct {
	ctx                context.Context
	Config             config.Config
	AdminerModel       model.AdminerModel
	JwtBlacklistsModel model.JwtBlacklistsModel
}

func NewJwtMiddleware(c config.Config, jwtModel model.JwtBlacklistsModel, adminerModel model.AdminerModel) *JwtMiddleware {
	return &JwtMiddleware{
		ctx:                context.Background(),
		Config:             c,
		AdminerModel:       adminerModel,
		JwtBlacklistsModel: jwtModel,
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

		//检查用户状态
		uidValue := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtUserId, *claims)
		if uidValue == "" {
			httpx.Error(w, errorx.ErrAuth)
			return
		}
		uid, err := strconv.Atoi(uidValue)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		adminer, err := m.AdminerModel.FindOne(m.ctx, int64(uid))
		if err != nil && err != model.ErrNotFound {
			logx.Error(err)
			httpx.Error(w, errorx.InternalError(err))
			return
		}
		if err == model.ErrNotFound {
			httpx.Error(w, errorx.ErrAuth)
			return
		}
		if adminer.Status == 0 {
			httpx.Error(w, errorx.AuthError(errors.New("当前用户不可用")))
			return
		}

		// 校验token是否已被加入黑名单
		exist, err := m.JwtBlacklistsModel.ExistByToken(m.ctx, jwtToken)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, errorx.InternalError(err))
			return
		}
		if exist {
			//如果存在了，则失效
			httpx.Error(w, errorx.ErrAuth)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
