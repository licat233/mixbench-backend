package logic

import (
	"errors"
	"fmt"

	"github.com/mojocn/base64Captcha"
)

type CaptchaConfig struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var CaptchaKeyPrefix = "LicatCaptcha"

var (
	ErrEmailTimed      = errors.New("requests are too frequent. Please try again in a minute")
	ErrEmailSendFailed = errors.New("mail sending failed, please check email address")
)

func formatEmailCaptchaKey(captchaId interface{}) string {
	return fmt.Sprintf("%s:%v", CaptchaKeyPrefix, captchaId)
}

func formatEmailCaptchaReceiverKey(receiver any) string {
	return fmt.Sprintf("captchaReceiver:%v", receiver)
}

func generateDigitCaptcha(store base64Captcha.Store) (captchaId string, err error) {
	defer func() {
		e := recover()
		if e != nil {
			err = fmt.Errorf("panic err:%+v", e)
		}
	}()

	driver := base64Captcha.NewDriverDigit(40, 150, 6, 0, 0)
	c := base64Captcha.NewCaptcha(driver, store)
	captchaId, _, err = c.Generate()
	if err != nil {
		err = fmt.Errorf("generateDigitCaptcha c.Generate() err:%+v", err)
	}
	return
}
