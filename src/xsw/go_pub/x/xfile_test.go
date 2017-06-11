package x

import (
	"testing"
)

func Test_xFile(t *testing.T) {
	var str1 string
	strPath := "c:/d1/d2/f1.ext"
	str1 = GetDirName(strPath)
	if str1 != "c:/d1/d2/" {
		t.Errorf("GetDirName fail:%s", str1)
	}
	str1 = GetFileName(strPath)
	if str1 != "f1.ext" {
		t.Errorf("GetFileName fail:%s", str1)
	}
	str1 = GetParentFileName(strPath)
	if str1 != "d2/f1.ext" {
		t.Errorf("GetParentFileName fail:%s", str1)
	}
	// names := GetDirNames(strPath)

	str1 = GetParentName(strPath)
	if str1 != "d2" {
		t.Errorf("GetParentName2 fail:%s", str1)
	}
}
