/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:15:59
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"time"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生产token
func (l *GenerateTokenLogic) GenerateToken(in *admin_pb.GenerateTokenReq) (*admin_pb.GenerateTokenResp, error) {
	if in.JwtData == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	if in.AutoLogin {
		accessExpire = 86400 * 7 //7天免登录
	}
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.JwtData)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.GenerateTokenResp{
		JwtToken: &admin_pb.JwtToken{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secret string, now, expire int64, data *admin_pb.JwtData) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = now + expire
	claims["iat"] = now
	claims[ctxdata.CtxKeyJwtUserId] = data.UserId
	claims[ctxdata.CtxKeyJwtIp] = data.Ip
	claims[ctxdata.CtxKeyJwtPlatform] = data.Platform
	return jwtx.CreateToken(claims, secret)
}
