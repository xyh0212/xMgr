package x

import (
	"fmt"
	"testing"
)

func Test_String(t *testing.T) {
	str := "0123456789"
	str1 := SubStr(str, 1, 3)
	if str1 != "123" {
		t.Error("SubStr fail")
	}

	var id int64
	id = 999999999991
	fmt.Printf("hotgame_hl%d\n", id)
}
