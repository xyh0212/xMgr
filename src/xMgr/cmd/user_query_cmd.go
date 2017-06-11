package cmd

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"testing"
	//"fmt"
	"xMgr/db"
	//"xMgr/websocket"
	//"github.com/bitly/go-simplejson"
	"strings"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":6,"param":{"name":"xxxxx"}}
func DoQueryUserCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoQueryUserCmd")
	name, _ := cmd.jsParam.CheckGet("name")
	if name == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_NAME_PARAM)
		return msg, x.XErrStr("name not exist")
	}
	gname, err := name.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	return checkUser(gname)
}
func checkUser(name string) ([]byte, *x.Error) {
	for _, value := range db.TbUserMap {
		if value.Name == name && value.UserGroupId == 1 {
			msg, _ := CreateManageCmdMsg()
			return msg, x.XErrStr("you are a manager")
		}
	}
	msg, _ := CreateErrMsg(sys.CODE_MANAGE_NOT_AUTHORIZATION)
	return msg, x.XErrStr("you are not a manager")
}

// {'cmd':6}发送数据

func CreateManageCmdMsg() ([]byte, *x.Error) {
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_QUERY_USER
	msgMap["json"] = getSUserNameJStr()
	return CreateAnyMsg(msgMap)
}
func getSUserNameJStr() string {
	var str string
	var arr []*db.SUser
	arr = make([]*db.SUser, len(db.TbUserMap))
	str = `{"userjson":[`
	i := 0
	x.PrintInfo(db.TbUserMap)
	for _, value := range db.TbUserMap {
		arr[i] = value
		i++
	}
	arr = sortArrySUser(arr)
	for _, a := range arr {
		str += `"` + a.Name + `",`
	}
	str = strings.TrimRight(str, ",")
	str += `]}`
	x.PrintInfo(str)
	return str
}

//排序SGS数组
func sortArrySUser(arr []*db.SUser) []*db.SUser {
	var temp *db.SUser
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i].Id < arr[j].Id {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
