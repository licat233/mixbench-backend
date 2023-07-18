package utils

import (
	"fmt"
	"testing"
)

func Test_SpliteURL(t *testing.T) {
	s := "http://localhost:80/pages/chat/index"
	res := GetDomain(s)
	fmt.Println(res)
}

func Test_IsPublic(t *testing.T) {
	s := "/v1/api/public/authentication/auth"
	res := IsPublicPath(s)
	fmt.Println(res)
}
