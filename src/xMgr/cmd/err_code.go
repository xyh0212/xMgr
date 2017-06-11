package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	// "github.com/server-nado/orm"
	"github.com/bitly/go-simplejson"
	"strconv"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

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

// {'cmd':10, 'code':0, 'info':'成功'}

func CreateErrMsg(code int) ([]byte, *x.Error) {
	var js *simplejson.Json
	var err error
	js, err = simplejson.NewJson([]byte(`{}`))
	if err != nil {
		return nil, x.XErr(err)
	}
	js.Set("cmd", sys.CMD_ERROR_MSG)
	js.Set("code", code)
	js.Set("info", GetCodeString(code))
	var strMsg []byte
	strMsg, err = js.Encode()
	return strMsg, x.XErr(err)
}

func CreateAnyMsg(any map[string]interface{}) ([]byte, *x.Error) {
	var js *simplejson.Json
	var err error
	js, err = simplejson.NewJson([]byte(`{}`))
	if err != nil {
		return nil, x.XErr(err)
	}
	for key, value := range any {
		js.Set(key, value)
		//x.PrintDbg(key, value)
	}
	var strMsg []byte
	strMsg, err = js.Encode()
	return strMsg, nil
}

//cmd = 7
func CreateSucceedMsg(cmd int) ([]byte, *x.Error) {
	ss := make(map[string]interface{})
	ss["cmd"] = cmd
	Msg, err := CreateAnyMsg(ss)
	return Msg, err
}
