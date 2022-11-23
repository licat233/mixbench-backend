package logic

import (
	"context"
	"errors"
	"mixbench/service/v1/captcha/rpc/captcha_pb"
	"mixbench/service/v1/captcha/rpc/internal/svc"
	"strings"

	"mixbench/common/errorx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
)

type AuthEmailCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthEmailCaptchaLogic {
	return &AuthEmailCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/** AuthEmailCaptcha 邮箱验证码校验
 */
func (l *AuthEmailCaptchaLogic) AuthEmailCaptcha(in *captcha_pb.AuthEmailCaptchaReq) (*captcha_pb.AuthEmailCaptchaResp, error) {
	//检查邮箱格式
	email := strings.TrimSpace(in.Email)
	//校验邮箱格式是否正确
	if !govalidator.IsEmail(email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}
	//处理captchaId
	CaptchaId := strings.TrimSpace(in.CaptchaId)
	if len(CaptchaId) == 0 {
		return nil, errorx.ExtRpcErr(errors.New("captcha_id is nil"))
	}
	key := formatEmailCaptchaKey(CaptchaId)
	exists, err := l.svcCtx.RedisClient.ExistsCtx(l.ctx, key)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//如果不存在，说明该captcha_id可能是假的
	if !exists {
		return &captcha_pb.AuthEmailCaptchaResp{
			Success: false,
		}, nil
	}

	//存在，则取出
	emailV, err := l.svcCtx.RedisClient.GetCtx(l.ctx, key)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//检查邮箱与captcha_id是否对应
	if emailV != email {
		//不相等，说明该captcha_id不属于该email
		return &captcha_pb.AuthEmailCaptchaResp{
			Success: false,
		}, nil
	}

	//校验captcha与solution
	solution := in.Solution
	ok := l.svcCtx.CaptchaStore.Verify(CaptchaId, solution, false)
	if !ok {
		return &captcha_pb.AuthEmailCaptchaResp{
			Success: false,
		}, nil
	}

	//清除验证码
	_ = l.svcCtx.CaptchaStore.Get(CaptchaId, true)

	//返回
	return &captcha_pb.AuthEmailCaptchaResp{
		Success: ok,
	}, nil
}
