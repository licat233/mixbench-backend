package emailx

import "errors"

var (
	ErrEmailTimed      = errors.New("requests are too frequent. Please try again in a minute")
	ErrEmailSendFailed = errors.New("mail sending failed, please check email address")
)
