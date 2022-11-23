package logic

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"mixbench/service/v1/captcha/rpc/captcha_pb"
	"mixbench/service/v1/captcha/rpc/internal/svc"
	"strings"
	"time"

	"mixbench/common/errorx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

type GetEmailCaptchaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmailCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmailCaptchaLogic {
	return &GetEmailCaptchaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/** GetEmailCaptcha 获取邮箱验证码
 */
func (l *GetEmailCaptchaLogic) GetEmailCaptcha(in *captcha_pb.GetEmailCaptchaReq) (resp *captcha_pb.GetEmailCaptchaResp, err error) {
	defer func() {
		e := recover()
		if e != nil {
			resp = nil
			logx.Error(fmt.Errorf("has panic: %+v", e))
			err = errorx.IntRpcErr(fmt.Errorf("has panic: %+v", e))
		}
	}()
	email := strings.TrimSpace(in.Email)
	//校验邮箱格式是否正确
	if !govalidator.IsEmail(email) {
		return nil, errorx.ExtRpcErr(errorx.ErrInvalidEmail)
	}

	//生成captcha_id
	captchaId, err := generateDigitCaptcha(l.svcCtx.CaptchaStore)
	if err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//获取captcha_id的answer
	answer := l.svcCtx.CaptchaStore.Get(captchaId, false)

	key := formatEmailCaptchaKey(captchaId)

	//先存入redis中，key为邮箱，value为验证码，过期时间为 expireTime
	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, key, email, int(l.svcCtx.CaptchaExpire.Seconds())); err != nil {
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//发送邮件到邮箱
	err = l.SendCaptcha(l.ctx, answer, email)
	if err != nil {
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, e := l.svcCtx.RedisClient.DelCtx(l.ctx, key); e != nil {
			logx.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == ErrEmailTimed {
			return nil, errorx.ExtRpcErr(err)
		}
		if err == ErrEmailSendFailed {
			return nil, errorx.ExtRpcErr(err)
		}
		logx.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//发送成功
	expiresAt := time.Now().Add(l.svcCtx.CaptchaExpire).Unix()
	return &captcha_pb.GetEmailCaptchaResp{
		CaptchaId: captchaId,
		ExpiresAt: expiresAt,
	}, nil
}

func (l *GetEmailCaptchaLogic) SendCaptcha(ctx context.Context, chaptcha string, receiver string) (err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("SendCaptcha has panic: %+v", e)
		}
	}()
	if len(chaptcha) == 0 {
		return errors.New("captcha is nil")
	}
	if len(receiver) == 0 {
		return errors.New("receiver is nil")
	}

	key := formatEmailCaptchaReceiverKey(receiver)
	exists, err := l.svcCtx.RedisClient.ExistsCtx(ctx, key)
	if err != nil {
		return err
	}

	if exists {
		return ErrEmailTimed
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", l.svcCtx.Config.Email.Username)
	msg.SetHeader("To", l.svcCtx.Config.Email.Receiver...)
	msg.SetHeader("Subject", "mixbench验证码")
	// 新建拨号器
	dialer := gomail.NewDialer(l.svcCtx.Config.Email.Host, l.svcCtx.Config.Email.Port, l.svcCtx.Config.Email.Username, l.svcCtx.Config.Email.Secret)
	// 关闭SSL协议认证
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	msg.SetBody("text/html", emailCaptchaBody(chaptcha))
	if err := dialer.DialAndSend(msg); err != nil {
		logx.Error(fmt.Sprintf("邮件发送失败:%#+v", err))
		return ErrEmailSendFailed
	}
	if err := l.svcCtx.RedisClient.SetexCtx(ctx, key, chaptcha, 60); err != nil {
		return err
	}
	return nil
}

func emailCaptchaBody(captcha string) string {
	return fmt.Sprintf(`<div style="max-width:600px;min-width:300px;margin:0 auto;background:#fff;border-collapse:collapse;font-size:18px;overflow:hidden">
	<div style=width:100vw>
		<div style=background:#000;padding:30px;color:#fff>
          系统通知-验证码
		</div>
	</div>
	<div style=width:100vw>
		<div style="padding:20px 15px">
			<h1 style="color:#000;border-bottom:1px solid;text-align:center;margin-bottom:30px">验证码</h1>
			<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 36px">
            <p>你本次的验证码为：%s <br/>为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>
		</div>
	</div>
	<div style=position:fixed;left:0;bottom:0;width:100vw>
		<div style=background:#eee;padding:10px;font-size:10px;text-align:center>系统邮件</div>
	</div>
</div>`, captcha)
}
