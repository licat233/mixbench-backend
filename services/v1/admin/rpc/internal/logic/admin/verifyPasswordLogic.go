/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:19:50
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/common/xcrypt"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVerifyPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyPasswordLogic {
	return &VerifyPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 验证管理员密码
func (l *VerifyPasswordLogic) VerifyPassword(in *admin_pb.VerifyPasswordReq) (*admin_pb.VerifyPasswordResp, error) {
	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, in.AdminerId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("this administrator cannot be found"))
	}
	return &admin_pb.VerifyPasswordResp{
		Ok: xcrypt.CheckPassword(in.Password, adminer.Password),
	}, nil
}
