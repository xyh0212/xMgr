package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":6,"param":{"name":"xxxxx","pass":"xxxxxx"}}
func DoAddUserCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoAddUserCmd")
	name, _ := cmd.jsParam.CheckGet("name")
	if name == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_NAME_PARAM)
		return msg, x.XErrStr("name not exist")
	}
	password, _ := cmd.jsParam.CheckGet("pass")
	if password == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_PASS_PARAM)
		return msg, x.XErrStr("pass not exist")
	}
	gname, err := name.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	gpassword, err := password.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	return addUser(gname, gpassword)
}
func addUser(name, gpassword string) ([]byte, *x.Error) {
	user := new(db.SUser)
	user.Name = name
	user.Pwd = gpassword
	_, err := db.G_OrmDB.Insert(user)
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ADD_USER_FAIL)
		return msg, x.XErrStr("add  fail")
	}
	db.TbUserMap[user.Id] = user
	msg, _ := CreateManageAddCmdMsg()
	return msg, x.XErrStr("add  succeed")
}

//cmd=9
func CreateManageAddCmdMsg() ([]byte, *x.Error) {
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_ADD_USER
	return CreateAnyMsg(msgMap)

}
