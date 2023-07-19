/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:18:54
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/redisx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注销登录
func (l *LogoutLogic) Logout(in *admin_pb.LogoutReq) (*admin_pb.LogoutResp, error) {
	if in.Token == "" {
		return nil, errorx.ExtRpcErr(errors.New("token equal nil"))
	}
	claims, ok := jwtx.ValidateToken(in.Token, os.Getenv("JWT_SECRET"))
	if !ok {
		return nil, errorx.ExtRpcErr(errors.New("validate token failed"))
	}
	id, _ := jwtx.GetId(claims)
	exp := jwtx.ExpiredTime(claims)
	currentTime := time.Now().Local()
	expirationTime := time.Unix(int64(exp), 0).Local()
	var err error
	if expirationTime.After(currentTime) {
		seconds := int(expirationTime.Sub(currentTime).Seconds())
		err = l.svcCtx.RedisCli.SetexCtx(l.ctx, redisx.FormatTokenblacklistKey(in.Token), fmt.Sprint(id), seconds)
	} else {
		err = l.svcCtx.RedisCli.SetCtx(l.ctx, redisx.FormatTokenblacklistKey(in.Token), fmt.Sprint(id))
	}
	if err != nil {
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.LogoutResp{}, nil
}
