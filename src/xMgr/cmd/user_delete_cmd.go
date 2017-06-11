package cmd

import (
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"testing"
	//"fmt"
	"xMgr/db"
	//"xMgr/websocket"
	//"github.com/bitly/go-simplejson"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":7,"param":{"name":"xxxxx"}}
func DoDeleteUserCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
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
	return deleteUser(gname)
}
func deleteUser(name string) ([]byte, *x.Error) {
	for key, value := range db.TbUserMap {
		if value.Name == name {
			_, err := db.G_OrmDB.Delete(value)
			if err != nil {
				x.LogErr(err)
				msg, _ := CreateErrMsg(sys.CODE_DELETE_USER_FAIL)
				return msg, x.XErrStr("fail")
			}
			delete(db.TbUserMap, key)
			msg, _ := CreateManageDeleteCmdMsg()
			return msg, x.XErrStr("succeed")
		}
	}
	msg, _ := CreateErrMsg(sys.CODE_NAME_NOT_EXIST)
	return msg, x.XErrStr("name not exit")

}

//cmd = 7
func CreateManageDeleteCmdMsg() ([]byte, *x.Error) {
	ss := make(map[string]interface{})
	ss["cmd"] = sys.CMD_DELETE_USER
	return CreateAnyMsg(ss)
}
