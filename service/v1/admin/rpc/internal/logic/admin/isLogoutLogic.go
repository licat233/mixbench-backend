/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:16:58
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLogoutLogic {
	return &IsLogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否已经注销
func (l *IsLogoutLogic) IsLogout(in *admin_pb.IsLogoutReq) (*admin_pb.IsLogoutResp, error) {
	//优先从redis中查找
	exist, err := l.svcCtx.RedisCli.ExistsCtx(l.ctx, in.JwtToken)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	if exist {
		return &admin_pb.IsLogoutResp{
			Status: exist,
		}, nil
	}

	//再从mysql中查找
	_, err = l.svcCtx.JwtBlacklistModel.FindByToken(l.ctx, in.JwtToken)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	return &admin_pb.IsLogoutResp{
		Status: err == nil,
	}, nil
}
