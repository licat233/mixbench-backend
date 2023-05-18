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
	"mixbench/services/v1/admin/api/internal/config"
	"mixbench/services/v1/admin/model"
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
		err := AuthLogic(m.ctx, r, m.Config, m.AdminerModel, m.JwtBlacklistsModel)
		if err != nil {
			httpx.ErrorCtx(m.ctx, w, err)
			return
		}

		// Passthrough to next handler if need
		next(w, r)
	}
}

func AuthLogic(ctx context.Context, r *http.Request, config config.Config, adminerModel model.AdminerModel, jwtBlacklistsModel model.JwtBlacklistModel) error {
	jwtToken := r.Header.Get(jwtx.ReqHeaderKey)
	if jwtToken == "" {
		return errorx.ErrAuth
	}

	//解析 token
	claims, err := jwtx.ParseToken(jwtToken, config.Auth.AccessSecret)
	if err != nil {
		return errorx.InternalError(err)
	}

	//获取请求设备
	reqPlatform, err := utils.RemotePlatform(r)
	if err != nil {
		return errorx.ErrPlatform
	}

	jwtPlatform := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtPlatform, *claims)
	//设备 校验
	if reqPlatform != jwtPlatform {
		return errorx.ErrPlatform
	}

	jwtIp := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtIp, *claims)
	//获取请求ip
	reqIp := utils.RemoteIp(r)
	//ip 校验
	if reqIp != jwtIp {
		return errorx.ErrIp
	}

	//检查用户状态
	uidValue := jwtx.GetClaimsValue(ctxdata.CtxKeyJwtUserId, *claims)
	if uidValue == "" {
		return errorx.ErrAuth
	}
	uid, err := strconv.Atoi(uidValue)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return errorx.ErrAuth
	}

	adminer, err := adminerModel.FindOne(ctx, int64(uid))
	if err != nil && err != model.ErrNotFound {
		logx.WithContext(ctx).Error(err)
		return errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		return errorx.ErrAuth
	}
	if adminer.Status == 0 {
		return errorx.AuthError(errors.New("当前用户不可用"))
	}

	// 校验token是否已被加入黑名单
	exists, err := jwtBlacklistsModel.ExistsByToken(ctx, jwtToken)
	if err != nil {
		logx.WithContext(ctx).Error(err)
		return errorx.InternalError(err)
	}
	if exists {
		return errorx.ErrAuth
	}
	return nil
}
