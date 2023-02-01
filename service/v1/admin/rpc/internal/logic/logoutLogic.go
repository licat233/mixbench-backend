/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:32
 * @LastEditors: licat
 * @LastEditTime: 2023-01-14 16:10:20
 * @Description: licat233@gmail.com
 */
package logic

import (
	"context"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"

	"mixbench/common/ctxdata"
	"mixbench/common/errorx"

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
	uid := ctxdata.GetUidFromCtx(l.ctx)
	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, uid)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return &admin_pb.LogoutResp{}, nil
	}
	//注销方案
	data := &model.JwtBlacklist{
		Uuid:     adminer.Uuid,
		Token:    in.JwtToken,
		Platform: in.Platform,
		Ip:       in.Ip,
	}

	_, err = l.svcCtx.JwtBlacklistModel.Insert(l.ctx, data)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.LogoutResp{}, nil
}
