package types

import "mixbench/common/emailx"

const (
	TaskTypeSendEmailCaptcha = "handle_send_email_captcha"
)

type SendEmailCaptchaReq struct {
	EmailAddrees []string
	RedisKey     string
	Message      *emailx.CaptchaMessage
}
