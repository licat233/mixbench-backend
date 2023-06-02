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
	"time"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
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
	exp := ctxdata.GetExpFromCtx(l.ctx)
	if exp == 0 {
		return nil, errorx.IntRpcErr(errors.New("failed to obtain exp value in ctx"))
	}
	//注销方案
	data := &model.JwtBlacklist{
		Id:       0,
		Token:    in.JwtToken,
		Platform: in.Platform,
		Ip:       in.Ip,
		ExpireAt: time.Unix(exp, 0),
	}

	_, err := l.svcCtx.JwtBlacklistModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.LogoutResp{}, nil
}
