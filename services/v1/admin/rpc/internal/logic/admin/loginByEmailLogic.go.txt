/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:17:33
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"errors"
	"time"

	"mixbench/common/errorx"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByEmailLogic {
	return &LoginByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 邮箱登录，注意避免同个邮箱，多点同时登陆带来的耦合性问题
func (l *LoginByEmailLogic) LoginByEmail(in *admin_pb.LoginByEmailReq) (*admin_pb.LoginResp, error) {
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
	//检查adminer的状态
	if err = verifyAdminerStatus(adminer); err != nil {
		return nil, errorx.ExtRpcErr(err)
	}
	//生成jwt-token
	jwtData := &admin_pb.JwtData{
		UserId:   adminer.Id,
		Ip:       in.JwtData.Ip,
		Platform: in.JwtData.Platform,
	}

	generateToken := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	res, err := generateToken.GenerateToken(&admin_pb.GenerateTokenReq{
		JwtData:   jwtData,
		AutoLogin: in.AutoLogin,
	})
	if err != nil {
		return nil, err
	}

	//更新登录次数和登录时间
	adminer.LoginCount++
	adminer.LastLogin = time.Now()
	if err = l.svcCtx.AdminerModel.Update(l.ctx, adminer); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//登录成功
	return &admin_pb.LoginResp{
		JwtToken: &admin_pb.JwtToken{
			AccessToken:  res.JwtToken.AccessToken,
			AccessExpire: res.JwtToken.AccessExpire,
			RefreshAfter: res.JwtToken.RefreshAfter,
		},
	}, nil
}
