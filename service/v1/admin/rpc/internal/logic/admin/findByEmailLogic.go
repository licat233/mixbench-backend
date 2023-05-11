/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:15:34
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"errors"
	"mixbench/common/errorx"
	"mixbench/service/v1/admin/model"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/logic/dataconv"
	"mixbench/service/v1/admin/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByEmailLogic {
	return &FindByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 通过邮箱查找管理员
func (l *FindByEmailLogic) FindByEmail(in *admin_pb.FindByEmailReq) (*admin_pb.FindByEmailResp, error) {
	//校验邮箱
	if !govalidator.IsEmail(in.Email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}
	//查找是否有绑定该邮箱的用户
	adminer, err := l.svcCtx.AdminerModel.FindByEmail(l.ctx, in.Email)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//账户不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("account does not exist"))
	}

	return &admin_pb.FindByEmailResp{
		Adminer: dataconv.MdAdminerToPbAdminer(adminer),
	}, nil
}
