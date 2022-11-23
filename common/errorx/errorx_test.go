package errorx

import (
	"errors"
	"fmt"
	"testing"

	"google.golang.org/grpc/status"
)

func Test_InternalError(t *testing.T) {
	err := InternalError(errors.New("调试错误信息"), errors.New("提示错误信息"))
	fmt.Printf("%#v\n", err)
}

func Test_StatusErr(t *testing.T) {
	err := errors.New("错误消息")
	// s := status.Convert(err)
	// fmt.Printf("%v %v", (*s).Code(), (*s).Message())

	errs := status.Error(403, err.Error())
	e := FromRpcErr(errs)
	fmt.Printf("%#v", e)
}
