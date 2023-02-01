/*
 * @Author: licat
 * @Date: 2023-01-11 11:47:48
 * @LastEditors: licat
 * @LastEditTime: 2023-01-11 14:38:16
 * @Description: licat233@gmail.com
 */
package response

import (
	"fmt"
	"net"
	"testing"
)

func Test_Response(t *testing.T) {
	res, _ := LocalIp()
	fmt.Println(res)
}

func LocalIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	res := "0.0.0.0"
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				res = ipnet.IP.String()
			}
		}
	}
	return res, nil
}
