package websocket

import (
	//"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
	//	"strings"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

func CreateMsg(m map[string]interface{}) ([]byte, *x.Error) {
	var js *simplejson.Json
	var err error
	js, err = simplejson.NewJson([]byte(`{}`))
	if err != nil {
		return nil, x.XErr(err)
	}
	for key, value := range m {
		js.Set(key, value)
	}
	var strMsg []byte
	strMsg, err = js.Encode()
	return strMsg, x.XErr(err)
}
func GetCodeString(code int) string {
	if db.TbCodeStringMap == nil {
		return strconv.Itoa(code)
	}

	rec := db.TbCodeStringMap.GetObj(int64(code))
	if rec == nil {
		return strconv.Itoa(code)
	}
	str, _ := rec.GetString(1)
	if str != nil {
		return *str
	}
	return ""
}

///for example
//return  {"cmd":"10","info":"2122"}
//the "cmd" is  key of map ,the "10" is value of map

//return  {"cmd":"10","info":"err"}

func CreateErrMsg(code int) []byte {
	m := make(map[string]interface{})
	m["cmd"] = sys.HTTP_ERR
	m["code"] = code
	m["info"] = GetCodeString(code)
	b, _ := CreateMsg(m)
	return b
}

func CreateErrInfoMsg(code int, info string) []byte {
	m := make(map[string]interface{})
	m["cmd"] = sys.HTTP_ERR
	m["code"] = code
	m["info"] = info
	b, _ := CreateMsg(m)
	return b
}

//return  {"cmd":"1","info":"suc"} CODE_SUCCEED
func CreateSucceedMsg(info string) []byte {
	m := make(map[string]interface{})
	m["cmd"] = sys.HTTP_SUCCEED
	m["info"] = info
	b, _ := CreateMsg(m)
	return b
}

//return  {"cmd":"1","info":"suc","info2":"suc"} CODE_SUCCEED
func CreateSucceedMultiMsg(info1, info2 string) []byte {
	m := make(map[string]interface{})
	m["cmd"] = sys.HTTP_SUCCEED
	m["info"] = info1
	m["info2"] = info2
	b, _ := CreateMsg(m)
	return b
}
