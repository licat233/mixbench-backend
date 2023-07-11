package types

import "mixbench/common/emailx"

const (
	SendEmailCaptchaKey = "handle_send_emailcaptcha"
)

type SendEmailCaptchaReq struct {
	EmailAddrees []string
	RedisKey     string
	Message      emailx.CaptchaMessage
}
