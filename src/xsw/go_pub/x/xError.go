package x

import (
	"errors"
	"fmt"
	"runtime"
)

type Error struct {
	str string
}

func (self *Error) GetStr() string {
	return self.str
}

func XErr(eRaw error) (e *Error) {
	if eRaw == nil {
		return XErrStr("err nil")
	}

	strFile, nLine, _ := GetSourceFileLine(2)
	e = new(Error)
	e.str = fmt.Sprintf("%s:%d %s", strFile, nLine, eRaw.Error())
	return e
}

func XErrStr(str string) *Error {
	strFile, nLine, _ := GetSourceFileLine(2)
	e := new(Error)
	e.str = fmt.Sprintf("%s:%d %s", strFile, nLine, str)

	return e
}

func GetSourceFileLine(nDepth int) (string, int, error) {
	var err error
	var ok bool
	var log_lineno int
	var file_name, short_name string
	_, file_name, log_lineno, ok = runtime.Caller(nDepth)
	if !ok {
		err = errors.New("call runtime.Caller() fail")
	}
	file_len := len(file_name)
	j := 0
	for i := file_len - 1; i > 0; i-- {
		c := file_name[i]
		if c == '\\' || c == '/' {
			j += 1
			if j > 2 {
				short_name = file_name[i+1:]
				break
			}
		}
	}

	return short_name, log_lineno, err
}
