package xcrypt

import (
	"fmt"
	"testing"
)

func Test_HashPassword(t *testing.T) {
	res, _ := HashPassword("Abcd12345")
	fmt.Println(res)
}

func Test_CheckPassword(t *testing.T) {
	a := "Abcd12345"
	b := "$2a$10$in/bw.cASq41miAYh2soTOeUg7e1Z/0r0LlKazmiSTlWvA5dSjzsi"
	res := CheckPassword(a, b)
	fmt.Println(res)
}
