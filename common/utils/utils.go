package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func RandomDigits(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().Unix()))
	var container []string
	for i := 0; i < length; i++ {
		n := rnd.Int()
		container = append(container, strconv.Itoa(n))
	}
	return strings.Join(container, "")
}

func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func TimestampHandle(timestamp int64) int64 {
	str := strconv.FormatInt(timestamp, 10)
	n := len(str)
	if n > 10 {
		str = str[:10]
		res, _ := strconv.ParseInt(str, 10, 64)
		return res
	} else if n < 10 {
		//返回zero时间戳
		zeroTime := time.Time{}
		return zeroTime.Unix()
	} else {
		return timestamp
	}
}

func IdValid(id int64) bool {
	return id > 0
}

func NameValid(name string) bool {
	return strings.TrimSpace(name) != ""
}

func PasswordPattern(password string) bool {
	matched, _ := regexp.MatchString("^.{6,34}$", password)
	return matched
}
