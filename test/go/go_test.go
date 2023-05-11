/*
 * @Author: licat
 * @Date: 2023-02-18 15:13:52
 * @LastEditors: licat
 * @LastEditTime: 2023-02-18 15:20:09
 * @Description: licat233@gmail.com
 */
package gotest

import "testing"

type Animal interface {
	Cry() string
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
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

func Test_go(t *testing.T) {

}
