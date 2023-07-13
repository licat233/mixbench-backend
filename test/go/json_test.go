package gotest

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Unmarshal1(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	jsonData := `{"name":"Alice","age":25}`

	var p *Person
	json.Unmarshal([]byte(jsonData), &p)
	fmt.Println("原始结构体指针:", p)

	p.Age = 30
	pJSON, _ := json.Marshal(p)
	fmt.Println("修改后的JSON:", string(pJSON))
}

func Test_Unmarshal2(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	jsonData := `{"name":"Alice","age":25}`

	var p Person
	json.Unmarshal([]byte(jsonData), &p)
	fmt.Println("原始结构体值:", p)

	p.Age = 30
	pJSON, _ := json.Marshal(p)
	fmt.Println("修改后的JSON:", string(pJSON))
}
