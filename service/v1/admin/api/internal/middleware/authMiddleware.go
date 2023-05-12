/*
 * @Author: licat
 * @Date: 2023-02-17 15:41:34
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 16:45:21
 * @Description: licat233@gmail.com
 */
package middleware

import (
	"context"
	"errors"
	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/utils"
	"mixbench/service/v1/admin/api/internal/config"
	"mixbench/service/v1/admin/model"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthMiddleware struct {
	ctx                context.Context
	Config             config.Config
	AdminerModel       model.AdminerModel
	JwtBlacklistsModel model.JwtBlacklistModel
}

func NewAuthMiddleware(c config.Config, jwtModel model.JwtBlacklistModel, adminerModel model.AdminerModel) *AuthMiddleware {
	return &AuthMiddleware{
		ctx:                context.Background(),
		Config:             c,
		AdminerModel:       adminerModel,
		JwtBlacklistsModel: jwtModel,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get(jwtx.ReqHeaderKey)
		if jwtToken == "" {
			httpx.ErrorCtx(m.ctx, w, errorx.ErrAuth)
			return
		}

		//解析 token
		claims, err := jwtx.ParseToken(jwtToken, m.Config.Auth.AccessSecret)
		if err != nil {
			httpx.ErrorCtx(m.ctx, w, errorx.InternalError(err))
			return
		}

		//获取请求设备
		reqPlatform, err := utils.RemotePlatform(r)
		if err != nil {
			httpx.ErrorCtx(m.ctx, w, errorx.ErrPlatform)
			return
		}

		jwtPlatform := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtPlatform, *claims)
		//设备 校验
		if reqPlatform != jwtPlatform {
			httpx.ErrorCtx(m.ctx, w, errorx.ErrPlatform)
			return
		}

		jwtIp := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtIp, *claims)
		//获取请求ip
		reqIp := utils.RemoteIp(r)
		//ip 校验
		if reqIp != jwtIp {
			httpx.ErrorCtx(m.ctx, w, errorx.ErrIp)
			return
		}

		//检查用户状态
		uidValue := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtUserId, *claims)
		if uidValue == "" {
			httpx.ErrorCtx(m.ctx, w, errorx.ErrAuth)
			return
		}
		uid, err := strconv.Atoi(uidValue)
		if err != nil {
			logx.WithContext(m.ctx).Error(err)
			httpx.ErrorCtx(m.ctx, w, errorx.ErrAuth)
			return
		}

		adminer, err := m.AdminerModel.FindOne(m.ctx, int64(uid))
		if err != nil && err != model.ErrNotFound {
			logx.WithContext(m.ctx).Error(err)
			httpx.ErrorCtx(m.ctx, w, errorx.InternalError(err))
			return
		}
		if err == model.ErrNotFound {
			httpx.ErrorCtx(m.ctx, w, errorx.ErrAuth)
			return
		}
		if adminer.Status == 0 {
			httpx.ErrorCtx(m.ctx, w, errorx.AuthError(errors.New("当前用户不可用")))
			return
		}

		// 校验token是否已被加入黑名单
		_, err = m.JwtBlacklistsModel.FindOneByToken(m.ctx, jwtToken)
		if err != nil {
			logx.WithContext(m.ctx).Error(err)
			httpx.ErrorCtx(m.ctx, w, errorx.InternalError(err))
			return
		}
		if err == nil {
			//如果存在了，则失效
			httpx.ErrorCtx(m.ctx, w, errorx.ErrAuth)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}
