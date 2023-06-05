package crypt_test

import (
	"fmt"
	"mixbench/common/xcrypt"
	"testing"
)

func TestCrypt(t *testing.T) {
	pwd := "Abcd12345"
	// hashv, _ := xcrypt.HashPassword(pwd)
	// fmt.Println("hash:", hashv)
	hashv := "$2a$10$HoIY098HXE.4qb3x/EcIvu3YCKNp6pMIaBD7r1T0LQHTMNBk1e4fy"
	ok := xcrypt.CheckPassword(pwd, hashv)
	fmt.Println("结果:", ok)
}
