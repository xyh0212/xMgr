package cmd

import (
	//"github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":8,"param":{"name":"xxxxx","password":"xxxxxx"}}
func DoEditUserCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoEditUserCmd")
	name, _ := cmd.jsParam.CheckGet("name")
	if name == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_NAME_PARAM)
		return msg, x.XErrStr("name not exist")
	}
	password, _ := cmd.jsParam.CheckGet("pass")
	if password == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_PASS_PARAM)
		return msg, x.XErrStr("password not exist")
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
	return editUser(gname, gpassword)
}
func editUser(name, gpassword string) ([]byte, *x.Error) {
	for _, value := range db.TbUserMap {
		if value.Name == name {
			updataUser(value, gpassword)
			msg, _ := CreateManageEditCmdMsg()
			return msg, x.XErrStr("succeed")
		}
	}
	msg, _ := CreateErrMsg(sys.CODE_NAME_NOT_EXIST)
	return msg, x.XErrStr("name not exit")

}
func updataUser(v *db.SUser, password string) *x.Error {
	p := v.Pwd
	v.Pwd = password
	_, err := db.G_OrmDB.Update(v)
	if err != nil {
		v.Pwd = p
		return x.XErr(err)
	}
	return nil

}

//cmd=8
func CreateManageEditCmdMsg() ([]byte, *x.Error) {
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_EDIT_USER
	return CreateAnyMsg(msgMap)

}
