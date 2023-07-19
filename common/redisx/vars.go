package redisx

import "fmt"

func FormatIpblacklistKey(ip string) string {
	return fmt.Sprintf("ipblacklist:%s", ip)
}

func FormatTokenblacklistKey(token string) string {
	return fmt.Sprintf("tokenblacklist:%s", token)
}
