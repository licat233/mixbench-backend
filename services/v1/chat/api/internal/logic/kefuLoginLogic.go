package logic

import (
	"context"
	"errors"
	"time"

	"mixbench/common/errorx"
	"mixbench/common/jwtx"
	"mixbench/common/respx"
	"mixbench/common/xcrypt"
	"mixbench/services/v1/authentication/rpc/authentication_pb"
	"mixbench/services/v1/chat/api/internal/svc"
	"mixbench/services/v1/chat/api/internal/types"
	"mixbench/services/v1/chat/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

type KefuLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKefuLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KefuLoginLogic {
	return &KefuLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KefuLoginLogic) KefuLogin(req *types.KefuLoginReq) (any, error) {
	//调用captcha.rpc服务校验验证码
	res, err := l.svcCtx.AuthenticationRpc.VerifyCaptcha(l.ctx, &authentication_pb.VerifyCaptchaReq{
		CaptchaId: req.CaptchaId,
		Solution:  req.Solution,
	})
	if err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errorx.ExternalError(nil, errors.New("captcha auth failed"))
	}

	if len(req.Passport) == 0 {
		return nil, errorx.ExternalError(errors.New("passport is empty"))
	}
	err = xcrypt.PasswordValidator(req.Password)
	if err != nil {
		return nil, errorx.ExternalError(err)
	}
	//根据账户查找用户
	//进行账号登录
	kefu, err := l.svcCtx.KefuModel.FindByPassport(l.ctx, req.Passport)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}

	//账户不存在
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("account does not exist"))
	}

	//查找项目是否还在
	_, err = l.svcCtx.ProjectModel.FindById(l.ctx, kefu.ProjectId)
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}
	if err == model.ErrNotFound {
		return nil, errorx.ExternalError(errors.New("peoject does not exist"))
	}

	if !xcrypt.CheckPassword(req.Password, kefu.Password) {
		return nil, errorx.ExternalError(errors.New("account and password do not match"))
	}

	//检查账户的状态
	if kefu.Status == 2 {
		err := errors.New("this account has been disabled and cannot be logged in")
		return nil, errorx.ExternalError(err, err)
	}

	now := time.Now().Unix()

	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	if req.AutoLogin {
		accessExpire = 86400 * 7 //7天免登录
	}

	refreshAfter := now + accessExpire/2

	jwtToken, err := GenKefuToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, kefu)
	if err != nil {
		l.Logger.Error(err)
		return nil, errorx.InternalError(err)
	}

	return respx.NewData(&types.JwtToken{
		AccessToken:  jwtToken,
		AccessExpire: accessExpire,
		RefreshAfter: refreshAfter,
	}), nil
}

func GenKefuToken(secret string, now, expire int64, kefu *model.Kefu) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = now + expire
	claims["iat"] = now
	claims["id"] = kefu.Id
	return jwtx.CreateToken(claims, secret)
}
