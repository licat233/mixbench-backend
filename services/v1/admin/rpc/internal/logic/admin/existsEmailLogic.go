/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:15:02
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type ExistsEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExistsEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExistsEmailLogic {
	return &ExistsEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱是否存在
func (l *ExistsEmailLogic) ExistsEmail(in *admin_pb.ExistsEmailReq) (*admin_pb.ExistsEmailResp, error) {
	//校验邮箱
	if !govalidator.IsEmail(in.Email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}
	//查找是否有绑定该邮箱的用户
	_, err := l.svcCtx.AdminerModel.ExistsByEmail(l.ctx, in.Email)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	return &admin_pb.ExistsEmailResp{
		Exists: err == nil,
	}, nil
}
