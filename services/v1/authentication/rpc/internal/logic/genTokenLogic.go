package logic

import (
	"context"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/authentication/rpc/internal/svc"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type GenTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenTokenLogic {
	return &GenTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 生成jwt-token
func (l *GenTokenLogic) GenToken(in *authentication_pb.GenTokenReq) (*authentication_pb.GenTokenResp, error) {
	if in.JwtData == nil {
		return nil, errorx.SecRpcErr(errorx.ErrEmptyParam)
	}
	now := time.Now().Local().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	if in.AutoLogin {
		accessExpire = 86400 * 7 //7天免登录
	}
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.JwtData)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &authentication_pb.GenTokenResp{
		JwtToken: &authentication_pb.JwtToken{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			RefreshAfter: now + accessExpire/2,
		},
	}, nil
}

func (l *GenTokenLogic) getJwtToken(secret string, now, expire int64, data *authentication_pb.JwtData) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = now + expire
	claims["iat"] = now
	claims[jwtx.JwtUidKey] = data.UserId
	claims[jwtx.JwtIpKey] = data.Ip
	claims[jwtx.JwtPlatformKey] = data.Platform
	return jwtx.CreateToken(claims, secret)
}
