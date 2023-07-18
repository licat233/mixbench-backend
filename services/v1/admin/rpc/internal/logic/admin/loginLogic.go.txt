/*
 * @Author: licat
 * @Date: 2023-02-18 13:59:10
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 14:18:42
 * @Description: licat233@gmail.com
 */
package adminlogic

import (
	"context"
	"errors"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/xcrypt"
	"mixbench/services/v1/admin/model"
	"mixbench/services/v1/admin/rpc/admin_pb"
	"mixbench/services/v1/admin/rpc/internal/svc"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 账号登录
func (l *LoginLogic) Login(in *admin_pb.LoginReq) (*admin_pb.LoginResp, error) {
	if len(in.Passport) == 0 {
		return nil, errorx.ExtRpcErr(errors.New("passport is empty"))
	}
	err := xcrypt.PasswordValidator(in.Password)
	if err != nil {
		return nil, errorx.ExtRpcErr(err)
	}
	//根据账户查找用户
	//passport也可能是邮箱
	if govalidator.IsEmail(in.Passport) {
		//进行邮箱登录
		loginByEmail := NewLoginByEmailLogic(l.ctx, l.svcCtx)
		return loginByEmail.LoginByEmail(&admin_pb.LoginByEmailReq{
			Email:     in.Passport,
			AutoLogin: in.AutoLogin,
		})
	}

	//进行账号登录
	adminer, err := l.svcCtx.AdminerModel.FindByPassport(l.ctx, in.Passport)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//账户不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("account does not exist"))
	}

	if !xcrypt.CheckPassword(in.Password, adminer.Password) {
		return nil, errorx.ExtRpcErr(errors.New("account and password do not match"))
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

	return &admin_pb.LoginResp{
		JwtToken: &admin_pb.JwtToken{
			AccessToken:  res.JwtToken.AccessToken,
			AccessExpire: res.JwtToken.AccessExpire,
			RefreshAfter: res.JwtToken.RefreshAfter,
		},
	}, nil
}
