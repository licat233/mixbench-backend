package authenticationlogic

import (
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

var authKeyPrefix = "cache:auth:rpc"

// 防止某个ip请求过多的key
// 用于限制某ip请求次数过多
func formatIpKey(ip string) string {
	return fmt.Sprintf("%s:ip:%s", authKeyPrefix, ip)
}

// 防止ip使用email请求过多的key
// 用于限制某个ip给某email发送邮件次数过多
func formatIpEmailKey(email string, ip string) string {
	return fmt.Sprintf("%s:ip:%s:email%s", authKeyPrefix, ip, email)
}

// 方式使用email和captcha请求过多的key，校验阶段，防止硬破解，不限制ip是因为考虑局域网
// 用于校验邮箱和验证码
func formatEmailCaptchaKey(email string, captchaId string) string {
	return fmt.Sprintf("%s:email:%s:capcthaid:%s", authKeyPrefix, email, captchaId)
}

// 查找captcha对应email的key
func formatCaptchaKey(captchaId string) string {
	return fmt.Sprintf("%s:captchaId:%s", authKeyPrefix, captchaId)
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
