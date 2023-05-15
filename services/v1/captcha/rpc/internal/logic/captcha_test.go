package logic

import (
	"fmt"
	"testing"
	"time"

	"github.com/mojocn/base64Captcha"
)

func Test_digitCaptcha(t *testing.T) {
	store := base64Captcha.NewMemoryStore(10240, time.Minute*5)
	//生成captcha_id
	driver := base64Captcha.NewDriverDigit(40, 150, 6, 0, 0)
	c := base64Captcha.NewCaptcha(driver, store)

	// captchaId, question, answer := c.Driver.GenerateIdQuestionAnswer()
	// fmt.Printf("captchaId:%v, question:%v, answer:%v\n", captchaId, question, answer)

	captchaId, _, err := c.Generate()
	if err != nil {
		fmt.Printf("c.Generate() err:%+v\n", err)
		return
	}

	//获取captcha_id的answer
	answer := store.Get(captchaId, false)
	fmt.Printf("captchaId:%v, answer:%v\n", captchaId, answer)

	ok := store.Verify(captchaId, answer, false)
	fmt.Println("verify result: ", ok)
}
