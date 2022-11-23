package utils

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
