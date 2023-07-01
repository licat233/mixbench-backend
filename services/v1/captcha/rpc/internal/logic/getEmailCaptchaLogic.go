package logic

import (
	"context"
	"fmt"
	"mixbench/services/v1/captcha/rpc/captcha_pb"
	"mixbench/services/v1/captcha/rpc/internal/svc"
	"strings"
	"time"

	"mixbench/common/emailx"
	"mixbench/common/errorx"

	"github.com/asaskevich/govalidator"
	"github.com/zeromicro/go-zero/core/logx"
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
			l.Logger.Error(fmt.Errorf("has panic: %+v", e))
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
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}
	//获取captcha_id的answer
	answer := l.svcCtx.CaptchaStore.Get(captchaId, false)

	key := formatEmailCaptchaKey(captchaId)

	//先存入redis中，key为邮箱，value为验证码，过期时间为 expireTime
	if err := l.svcCtx.RedisClient.SetexCtx(l.ctx, key, email, int(l.svcCtx.CaptchaExpire.Seconds())); err != nil {
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	message := &emailx.CaptchaMessage{
		Captcha:  answer,
		Account:  in.Account,
		Datetime: in.Datetime,
		Address:  in.Address,
		IP:       in.Ip,
	}

	//发送邮件到邮箱
	err = emailx.SendCaptcha(l.ctx, "mixbench验证码", message, []string{email})
	if err != nil {
		// 如果发送失败了，就要删除刚刚存入redis的数据
		if _, e := l.svcCtx.RedisClient.DelCtx(l.ctx, key); e != nil {
			l.Logger.Error(err)
			return nil, errorx.IntRpcErr(err)
		}
		if err == emailx.ErrEmailTimed {
			return nil, errorx.ExtRpcErr(err)
		}
		if err == emailx.ErrEmailSendFailed {
			return nil, errorx.ExtRpcErr(err)
		}
		l.Logger.Error(err)
		return nil, errorx.IntRpcErr(err)
	}

	//发送成功
	expiresAt := time.Now().Add(l.svcCtx.CaptchaExpire).Unix()
	return &captcha_pb.GetEmailCaptchaResp{
		CaptchaId: captchaId,
		ExpiresAt: expiresAt,
	}, nil
}

// func (l *GetEmailCaptchaLogic) SendCaptcha(ctx context.Context, message *emailMessage, receiver string) (err error) {
// 	defer func() {
// 		e := recover()
// 		if e != nil {
// 			err = fmt.Errorf("SendCaptcha has panic: %+v", e)
// 		}
// 	}()
// 	if message == nil {
// 		return errors.New("message is nil")
// 	}
// 	if len(receiver) == 0 {
// 		return errors.New("receiver is nil")
// 	}

// 	key := formatEmailCaptchaReceiverKey(receiver)
// 	exists, err := l.svcCtx.RedisClient.ExistsCtx(ctx, key)
// 	if err != nil {
// 		return err
// 	}

// 	if exists {
// 		return emailx.ErrEmailTimed
// 	}

// 	msg := gomail.NewMessage()
// 	msg.SetHeader("From", l.svcCtx.Config.Email.Username)
// 	msg.SetHeader("To", l.svcCtx.Config.Email.Receiver...)
// 	msg.SetHeader("Subject", "mixbench验证码")
// 	// 新建拨号器
// 	dialer := gomail.NewDialer(l.svcCtx.Config.Email.Host, l.svcCtx.Config.Email.Port, l.svcCtx.Config.Email.Username, l.svcCtx.Config.Email.Secret)
// 	// 关闭SSL协议认证
// 	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

// 	body, err := emailCaptchaBody(ctx, message)
// 	if err != nil {
// 		e := fmt.Errorf("get emailCaptchaBody error: %s", err.Error())
// 		l.Logger.Error(e)
// 		return e
// 	}
// 	msg.SetBody("text/html", body)
// 	if err := dialer.DialAndSend(msg); err != nil {
// 		l.Logger.Error(fmt.Sprintf("邮件发送失败:%#+v", err))
// 		return emailx.ErrEmailSendFailed
// 	}
// 	if err := l.svcCtx.RedisClient.SetexCtx(ctx, key, message.Captcha, 60); err != nil {
// 		return err
// 	}
// 	return nil
// }

// type emailMessage struct {
// 	Captcha  string
// 	Account  string
// 	Datetime string
// 	Address  string
// 	IP       string
// }

// func emailCaptchaBody(ctx context.Context, message *emailMessage) (string, error) {
// 	root := utils.CurrentProjectRootPath()
// 	filename := path.Join(root, "deploy/email/template/captcha.html")
// 	data, err := os.ReadFile(filename)
// 	if err != nil {
// 		logx.WithContext(ctx).Error("File reading error", err)
// 		return "", err
// 	}
// 	body := string(data)
// 	body = strings.Replace(body, "__CAPTCHA__", message.Captcha, -1)
// 	body = strings.Replace(body, "__ACCOUNT__", message.Account, -1)
// 	body = strings.Replace(body, "__DATETIME__", message.Datetime, -1)
// 	body = strings.Replace(body, "__ADDRESS__", message.Address, -1)
// 	body = strings.Replace(body, "__IP__", message.IP, -1)
// 	return body, nil
// }
