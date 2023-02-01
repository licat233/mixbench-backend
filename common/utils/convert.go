/*
 * @Author: licat
 * @Date: 2022-11-17 14:55:03
 * @LastEditors: licat
 * @LastEditTime: 2023-01-31 15:14:09
 * @Description: licat233@gmail.com
 */
package utils

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// expToNumber 科学计算数字转int64
// 例： ExpNum = 1.641020629e+09
func ExpToNumber(ExpNum string) (int64, error) {
	if res, err := strconv.ParseInt(ExpNum, 10, 64); err == nil {
		return res, nil
	}
	var newNum float64
	_, err := fmt.Sscanf(ExpNum, "%e", &newNum)
	if err != nil {
		return 0, fmt.Errorf("fmt.Sscanf error, ExpiresAt:%s, err:%v", ExpNum, err)
	}
	num := fmt.Sprintf("%.f", newNum)
	res, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("strconv.ParseInt error, num:%s, err:%v", num, err)
	}
	return res, nil
}

func Int64ToTime(timestamp int64) time.Time {
	if timestamp <= 0 {
		return time.Time{}
	}
	timestamp = TimestampHandle(timestamp)
	return time.Unix(timestamp, 0).Local()
}

func StrToTime(timestr string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
}

func TimeToInt64(t time.Time) int64 {
	return t.Unix()
}

func Int64ToBool(n int64) bool {
	return n != 0
}

func BoolToInt64(ok bool) int64 {
	if ok {
		return 1
	}
	return 0
}

func JwtTokenToMd5(token string) (string, error) {
	if token == "" {
		return "", errors.New("token is empty")
	}
	return Md5V(token), nil
}

// StructToSlice 结构体转切片
func StructToSlice(f interface{}) []string {
	v := reflect.ValueOf(f)
	ss := make([]string, v.NumField())
	for i := range ss {
		ss[i] = fmt.Sprintf("%v", v.Field(i))
	}
	return ss
}

// 驼峰法to蛇形法
func ToSnakeCase(str string) string {
	var matchNonAlphaNumeric = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
	str = matchNonAlphaNumeric.ReplaceAllString(str, "_")     //非常规字符转化为 _
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}") //拆分出连续大写
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")  //拆分单词
	return strings.ToLower(snake)                             //全部转小写
}

// 蛇形法to大驼峰法
func ToCamelCase(str string) string {
	wds := strings.Split(str, "_")
	res := ""
	for _, wd := range wds {
		res += StrFirstToUpper(wd)
	}
	return res
}

// 如果是小写字母, 则变换为大写字母
func StrFirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str) // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 { // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				// fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}
