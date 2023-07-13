package utils

import (
	"errors"
	"net"
	"net/http"
	"regexp"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// GetReqFormValue 获取请求form字段值
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

// 获取请求头中的Origin
func RemoteOrigin(r *http.Request) string {
	return r.Header.Get("Origin")
}

// 获取请求头中的Referer
func RemoteReferer(r *http.Request) string {
	return r.Header.Get("Referer")
}

// 获取请求头中的User-Agent
func RemoteUserAgent(r *http.Request) string {
	return r.Header.Get("User-Agent")
}

func RemoteLanguage(r *http.Request) (languages []string) {
	acceptLanguage := r.Header.Get("Accept-Language")
	if acceptLanguage == "" {
		return
	}
	// 将Accept-Language字符串拆分成多个语言设置
	languageTags := strings.Split(acceptLanguage, ",")

	// 创建一个存储语言设置的切片
	languages = make([]string, 0, len(languageTags))

	// 迭代处理每个语言设置
	for _, tag := range languageTags {
		parts := strings.SplitN(tag, ";", 2)

		// 获取语言标签，忽略语言权重参数
		language := strings.TrimSpace(parts[0])

		// 添加到语言切片中
		languages = append(languages, language)
	}

	return
}

// RemotePlatform 获取远程客户端系统：ipad,iphone,mac,windows,android,linux,ubuntu
func RemotePlatform(r *http.Request) (string, error) {
	ua := RemoteUserAgent(r)
	if len(ua) == 0 {
		return ua, errors.New("未知设备")
	}
	//ipad
	padR := regexp.MustCompile(`(?i)ipad`)
	if padR.MatchString(ua) {
		return "ipad", nil
	}
	//iphone
	iphoneR := regexp.MustCompile(`(?i)iphone`)
	if iphoneR.MatchString(ua) {
		return "iphone", nil
	}
	//mac
	macR := regexp.MustCompile(`(?i)mac`)
	if macR.MatchString(ua) {
		return "mac", nil
	}
	//windows
	windowsR := regexp.MustCompile(`(?i)windows`)
	if windowsR.MatchString(ua) {
		return "windows", nil
	}
	//android
	androidR := regexp.MustCompile(`(?i)android`)
	if androidR.MatchString(ua) {
		return "android", nil
	}
	//linux
	linuxR := regexp.MustCompile(`(?i)linux`)
	if linuxR.MatchString(ua) {
		return "linux", nil
	}
	//ubuntu
	ubuntuR := regexp.MustCompile(`(?i)ubuntu`)
	if ubuntuR.MatchString(ua) {
		return "ubuntu", nil
	}
	return ua, errors.New("未知设备")
}

// 获取本地服务器的ip
func LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "0.0.0.0", err
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "0.0.0.0", nil
}

func IsWebsocketRequest(r *http.Request) bool {
	if r == nil {
		return false
	}
	if r.Header.Get("Upgrade") == "websocket" && r.Header.Get("Connection") == "Upgrade" {
		// 是 WebSocket 请求
		return true
	}
	return false
}

func GetDomain(host string) string {
	hostParts := strings.Split(host, ".")
	numParts := len(hostParts)

	// 提取主域名
	mainDomain := hostParts[numParts-2] + "." + hostParts[numParts-1]
	return mainDomain
}
