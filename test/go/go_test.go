/*
 * @Author: licat
 * @Date: 2023-02-18 15:13:52
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 15:20:09
 * @Description: licat233@gmail.com
 */
package gotest

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

type Animal interface {
	Cry() string
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

type Room struct {
	*Cat
}

func (c Cat) Cry() string {
	return c.Name
}

func (c Dog) Cry() string {
	return c.Name
}

func Demo() Animal {
	var p = Dog{}
	return p
}

func (r *Room) Do() error {
	panic("发生了恐慌")
}

func Test_go(t *testing.T) {
	rawURL := "https://example.com/path/to/resource"
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("URL 解析失败:", err)
		return
	}

	hostParts := strings.Split(parsedURL.Hostname(), ".")
	numParts := len(hostParts)

	// 提取主域名
	mainDomain := hostParts[numParts-2] + "." + hostParts[numParts-1]
	fmt.Println("主域名:", mainDomain)
}
