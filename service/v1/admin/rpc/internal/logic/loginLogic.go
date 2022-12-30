package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/admin/rpc/admin_pb"
	"mixbench/service/v1/admin/rpc/internal/svc"
	"mixbench/service/v1/model"
	"time"

	"mixbench/common/errorx"

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
		return nil, errorx.ExtRpcErr(errorx.New("passport is empty"))
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
	adminer, err := l.svcCtx.AdminerModel.FindOneByPassport(l.ctx, in.Passport)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//账户不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExtRpcErr(errors.New("account does not exist"))
	}

	if adminer.Passport != in.Passport {
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
