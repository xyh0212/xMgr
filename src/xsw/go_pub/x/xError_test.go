package x

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Error(t *testing.T) {
	err := errors.New("错误")
	e := XErr(err)
	fmt.Println(e)

}
