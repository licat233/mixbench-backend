package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

//GetReqFormValue 获取请求form字段值
func GetReqFormValue(r *http.Request, Key string) string {
	//r.Form 类型为 map[string][]string
	if r.Form.Has(Key) {
		return strings.TrimSpace(r.Form.Get(Key))
	}
	return ""
}

// RemoteIp 获取远程客户端的 IP
func RemoteIp(req *http.Request) string {
	addr := httpx.GetRemoteAddr(req)
	//可能会被仿冒，我们取最后一个IP，也就是原始IP
	strs := strings.Split(addr, ",")
	return strings.TrimSpace(strs[len(strs)-1])
}

// RemotePlatform 获取远程客户端系统：ipad,iphone,mac,windows,android,linux,ubuntu
func RemotePlatform(r *http.Request) (string, error) {
	u := r.Header.Get("User-Agent")
	if len(u) == 0 {
		u = r.Header.Get("user-agent")
		if len(u) == 0 {
			return u, errors.New("未知设备的请求！")
		}
	}
	//ipad
	padR := regexp.MustCompile(`(?i)ipad`)
	if padR.MatchString(u) {
		return "ipad", nil
	}
	//iphone
	iphoneR := regexp.MustCompile(`(?i)iphone`)
	if iphoneR.MatchString(u) {
		return "iphone", nil
	}
	//mac
	macR := regexp.MustCompile(`(?i)mac`)
	if macR.MatchString(u) {
		return "mac", nil
	}
	//windows
	windowsR := regexp.MustCompile(`(?i)windows`)
	if windowsR.MatchString(u) {
		return "windows", nil
	}
	//android
	androidR := regexp.MustCompile(`(?i)android`)
	if androidR.MatchString(u) {
		return "android", nil
	}
	//linux
	linuxR := regexp.MustCompile(`(?i)linux`)
	if linuxR.MatchString(u) {
		return "linux", nil
	}
	//ubuntu
	ubuntuR := regexp.MustCompile(`(?i)ubuntu`)
	if ubuntuR.MatchString(u) {
		return "ubuntu", nil
	}
	return u, errors.New("未知设备的请求")
}
