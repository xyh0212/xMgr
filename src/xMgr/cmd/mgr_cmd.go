package cmd

import (
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"reflect"
	//	"xMgr/db"
	"xMgr/sys"
	"xMgr/user"
	"xMgr/websocket"
	"xsw/go_pub/x"
)

type CmdInfo struct {
	nCmd    int
	jsParam *simplejson.Json
}

var (
	g_funcsCmd map[int]interface{}
)

func init() {

	g_funcsCmd = make(map[int]interface{})
	g_funcsCmd[sys.CMD_LOGIN] = DoLoginCmd
	g_funcsCmd[sys.CMD_QUERY_GS] = DoQueryGSCmd
	g_funcsCmd[sys.CMD_QUERY_SQL] = DoQuerySqlCmd
	g_funcsCmd[sys.CMD_CHAT] = DoChatCmd
	g_funcsCmd[sys.CMD_CHECK_SESSION] = DoCheckSessionCmd
	g_funcsCmd[sys.CMD_QUERY_USER] = DoQueryUserCmd
	g_funcsCmd[sys.CMD_DELETE_USER] = DoDeleteUserCmd
	g_funcsCmd[sys.CMD_EDIT_USER] = DoEditUserCmd
	g_funcsCmd[sys.CMD_ADD_USER] = DoAddUserCmd
	g_funcsCmd[sys.CMD_ADD_GS] = DoAddGSCmd
	g_funcsCmd[sys.CMD_DELETE_GS] = DoDeleteGSCmd
	g_funcsCmd[sys.CMD_EDIT_GS] = DoEditGSCmd
	g_funcsCmd[sys.CMD_QUERY_MANAGE_GS] = DoQueryGSManageCmd

}

func DoCmd(strJson string, cid int) *x.Error {
	c, err := pharseCmdStr(strJson)
	if err != nil {
		return err
		x.PrintDbg(err)
	}
	x.PrintDbg(cid, c.nCmd)
	ok := user.GSessionMgr.CheckPriviligeByCid(cid, c.nCmd) //AddSession(1, "ldkfnl")
	if !ok {
		msg, _ := CreateErrMsg(sys.CODE_NO_PRIVILIGE)
		websocket.SendMsg(cid, msg)
		return x.XErrStr("no Privilige")
	}
	oFunc := g_funcsCmd[c.nCmd]
	if oFunc == nil {
		return x.XErr(errors.New(fmt.Sprintf("cmd[%d] not exist", c.nCmd)))
	}
	results, _ := callReflect(g_funcsCmd, c.nCmd, c, cid)
	if len(results) != 2 {
		var str string
		str = fmt.Sprint("DoCmd:", c.nCmd, " 对应的函数返回值个数为", len(results))
		xerr := x.XErrStr(str)
		return xerr
	}
	v1 := results[0]
	websocket.SendMsg(cid, v1.Bytes())
	return nil
}

func callReflect(m map[int]interface{}, name int, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return nil, err
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return result, err
}

func pharseCmdStr(strJson string) (CmdInfo, *x.Error) {
	var c CmdInfo

	js, js_err := simplejson.NewJson([]byte(strJson))
	if js_err != nil {
		return c, x.XErr(js_err)
	}
	jsCmd := js.Get("cmd")
	if jsCmd == nil {
		return c, x.XErr(errors.New("not find cmd"))
	}

	var err error
	c.nCmd, err = jsCmd.Int()
	if err != nil {
		return c, x.XErr(err)
	}

	c.jsParam = js.Get("param")
	return c, nil
}
