/*
 * @Author: licat
 * @Date: 2023-01-31 17:18:24
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 15:52:15
 * @Description: licat233@gmail.com
 */
package luckydrawlogic

import (
	"context"
	"fmt"
	"time"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/services/v1/luckydraw/rpc/internal/svc"
	"mixbench/services/v1/luckydraw/rpc/luckydraw_pb"

	jwt "github.com/dgrijalva/jwt-go"
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

// 生产jwtToken
func (l *GenerateTokenLogic) GenerateToken(in *luckydraw_pb.GenerateTokenReq) (*luckydraw_pb.GenerateTokenResp, error) {
	if in.JwtData == nil {
		return nil, errorx.ErrEmptyParam
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.JwtData)
	if err != nil {
		return nil, errorx.IntRpcErr(fmt.Errorf("GenerateToken getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId) error: %v", err))
	}

	return &luckydraw_pb.GenerateTokenResp{
		JwtToken: &luckydraw_pb.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
		},
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secret string, now, expire int64, data *luckydraw_pb.JwtData) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = now + expire
	claims["iat"] = now
	claims[ctxdata.CtxKeyJwtUserId] = data.UserId
	claims[ctxdata.CtxKeyJwtIp] = data.Ip
	claims[ctxdata.CtxKeyJwtPlatform] = data.Platform
	return jwtx.CreateToken(claims, secret)
}
