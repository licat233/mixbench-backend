package gotest

import (
	"fmt"
	"testing"
)

func Test_defer(t *testing.T) {
	fmt.Println("before defer")
	ok := true
	if ok {
		return
	}
	defer func() {
		fmt.Println("deferred function")
	}()

	fmt.Println("before return")
}
