/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:19:32
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"errors"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutPasswordLogic {
	return &PutPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改密码
func (l *PutPasswordLogic) PutPassword(in *admin_pb.PutPasswordReq) (*admin_pb.StatusResp, error) {
	adminer, err := l.svcCtx.AdminerModel.FindOne(l.ctx, in.AdminerId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("this adminer cannot be found"))
	}
	if adminer.Password != in.OldPassword {
		return nil, errorx.ExtRpcErr(errors.New("原密码错误"))
	}
	adminer.Password = in.NewPassword
	err = l.svcCtx.AdminerModel.Update(l.ctx, adminer)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.StatusResp{}, nil
}
