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
	"testing"
	"time"
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

func echo(v any) {
	fmt.Printf("結果:%s\n", v)
}

func Test_go(t *testing.T) {
	var err error
	res := fmt.Sprintf("%s", err)
	echo(res == "")
}

func Test_Time(t *testing.T) {
	a := time.Unix(-1, 0).Local()
	// a := time.Time{}
	fmt.Println(a)
	fmt.Println(isZeroTime(a))
}

func isZeroTime(t time.Time) bool {
	return t.Before(time.Unix(0, 0))
}
