/*
 * @Author: licat
 * @Date: 2023-02-20 13:16:52
 * @LastEditors: licat
 * @LastEditTime: 2023-02-20 13:58:51
 * @Description: licat233@gmail.com
 */
package gotest

import (
	"fmt"
	"regexp"
	"testing"
)

func Test_reg(t *testing.T) {
	content := `services Admin {

		// adminer rpc

		//添加管理员

	  }
	  services Admin {

		// adminer aaa

		//添加s管理员

	  }`
	pattern := `(?s)service\s*\w+\s*\{(.*?)}`
	// matched, err := regexp.MatchString(pattern, content)
	// fmt.Println(matched, err)

	re := regexp.MustCompile(pattern)
	match := re.FindAllStringSubmatch(content, -1)

	fmt.Printf("%v", match)
}

func Test_pwd(t *testing.T) {
	password := "Pwd123@qqqwq"
	re := regexp.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)
	fmt.Println(re.MatchString(password))
}
