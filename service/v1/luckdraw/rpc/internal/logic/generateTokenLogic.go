package logic

import (
	"context"
	"fmt"
	"mixbench/service/v1/luckdraw/rpc/internal/svc"
	"mixbench/service/v1/luckdraw/rpc/luckdraw_pb"
	"time"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/common/jwtx"

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

// 生产jwtToken
func (l *GenerateTokenLogic) GenerateToken(in *luckdraw_pb.GenerateTokenReq) (*luckdraw_pb.GenerateTokenResp, error) {
	if in.JwtData == nil {
		return nil, errorx.ErrEmptyParam
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.JwtData)
	if err != nil {
		return nil, errorx.IntRpcErr(fmt.Errorf("GenerateToken getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId) error: %v", err))
	}

	return &luckdraw_pb.GenerateTokenResp{
		JwtToken: &luckdraw_pb.JwtToken{
			AccessToken:  jwtToken,
			AccessExpire: now + accessExpire,
		},
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secret string, now, expire int64, data *luckdraw_pb.JwtData) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = now + expire
	claims["iat"] = now
	claims[ctxdata.CtxKeyJwtUserId] = data.UserId
	claims[ctxdata.CtxKeyJwtIp] = data.Ip
	claims[ctxdata.CtxKeyJwtPlatform] = data.Platform
	return jwtx.CreateToken(claims, secret)
}
