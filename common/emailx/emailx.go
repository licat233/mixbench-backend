package emailx

import (
	"context"
	"crypto/tls"
	"fmt"
	"mixbench/common/utils"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

type Config struct {
	Subject      string // Subject字段通常出现在邮件的标题栏或邮件列表中，以便收件人可以快速浏览和筛选邮件
	HeaderParams map[string][]string
}

var (
	host         string = os.Getenv("EMAIL_HOST")
	port         int
	username     string = os.Getenv("EMAIL_USERNAME")
	password     string = os.Getenv("EMAIL_PASSWORD")
	AdminerEmail string = os.Getenv("ADMINER_EMAIL")
)

func init() {
	var err error
	port, err = strconv.Atoi(os.Getenv("EMAIL_PORT"))
	if err != nil {
		port = 465
	}
}

// 注意，发送邮件会出现比较大的延迟，所以最好设置ctx
func sendEmail(ctx context.Context, subject string, body string, receiver []string, headerParams map[string][]string) error {
	newCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	msg := gomail.NewMessage()
	msg.SetHeader("From", username)
	msg.SetHeader("To", receiver...)
	msg.SetHeader("Subject", subject)
	for k, v := range headerParams {
		msg.SetHeader(k, v...)
	}

	//发送邮件到邮箱，为了避免超时，交给协程去执行
	go func() {
		// 新建拨号器
		dialer := gomail.NewDialer(host, port, username, password)
		// 关闭SSL协议认证
		dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		msg.SetBody("text/html", body)
		if err := dialer.DialAndSend(msg); err != nil {
			logx.WithContext(newCtx).Error(fmt.Sprintf("email sending failed:%#+v", err))
			// return ErrEmailSendFailed
		}
	}()

	return nil
}

func genBody(ctx context.Context, htmlname string, data map[string]string) (string, error) {
	filename := path.Join(utils.ProjectRoot, "deploy/email/template/", htmlname)
	body, err := os.ReadFile(filename)
	if err != nil {
		logx.WithContext(ctx).Error("File reading error", err)
		return "", fmt.Errorf("file reading error: %s", err)
	}
	html := string(body)
	for k, v := range data {
		html = strings.ReplaceAll(html, k, v)
	}
	return html, nil
}

func SendNotification(ctx context.Context, subject, message string, receiver []string) error {
	body, err := genBody(ctx, "notification.html", map[string]string{
		"__TITLE__":   subject,
		"__CONTENT__": message,
	})
	if err != nil {
		e := fmt.Errorf("SendNotification error: %s", err.Error())
		logx.WithContext(ctx).Error(e)
		return e
	}
	return sendEmail(ctx, subject, body, receiver, nil)
}

type CaptchaMessage struct {
	Captcha  string
	Account  string
	Datetime string
	Address  string
	IP       string
}

func SendCaptcha(ctx context.Context, subject string, message *CaptchaMessage, receiver []string) error {
	body, err := genBody(ctx, "captcha.html", map[string]string{
		"__TITLE__":    subject,
		"__CAPTCHA__":  message.Captcha,
		"__ACCOUNT__":  message.Account,
		"__DATETIME__": message.Datetime,
		"__ADDRESS__":  message.Address,
		"__IP__":       message.IP,
	})
	if err != nil {
		e := fmt.Errorf("SendCaptcha error: %s", err.Error())
		logx.WithContext(ctx).Error(e)
		return e
	}
	return sendEmail(ctx, subject, body, receiver, nil)
}
