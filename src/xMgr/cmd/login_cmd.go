package cmd

//file:///C:/Users/Administrator/Desktop/wfile:///C:/Users/Administrator/Desktop/workplace/xMgr/src/xMgr/cmd/querygs_cmd.goorkplace/xMgr/src/xMgr/cmd/mgr_cmd.go

import (
	//"github.com/bitly/go-simplejson"
	"xMgr/db"
	"xMgr/user"
	//"xMgr/websocket"
	//"fmt"
	"xMgr/sys"

	"xsw/go_pub/x"
)

// {"cmd":1,"param":{"name":"xsw","pass":"xxxxx"}}
func DoLoginCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	pass, _ := cmd.jsParam.CheckGet("pass")
	if pass == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_PASS_PARAM)
		return msg, x.XErrStr("pass not exist")
	}
	// todo:检查
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
	gpass, err := pass.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}

	msg, err3 := CheckPassAndUser(cid, gname, gpass)
	return msg, err3

}

//check passweord and user
func CheckPassAndUser(cid int, name, pass string) ([]byte, *x.Error) {
	for _, value := range db.TbUserMap {
		if name == value.Name && pass == value.Pwd {
			//密码正确的事件
			strSession := user.GSessionMgr.CreateSession()
			value.Session = strSession
			updataSession(value)
			msg, _ := CreateUserLoginCmdMsg(strSession, value.UserGroupId)
			return msg, x.XErrStr("user succeed")

		}

	}
	msg, _ := CreateErrMsg(sys.CODE_PWD_ERR)
	return msg, x.XErrStr("password or username is erron")
}

// {'cmd':1,session:""}user成功登入
func CreateUserLoginCmdMsg(strSession string, userGroupId int) ([]byte, *x.Error) {
	s, ok := db.TbUserGroupMap[userGroupId]
	if !ok {
		return nil, x.XErrStr("db.TbUserGroupMap[userGroupId]的key不存在")
	}
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_LOGIN
	msgMap["session"] = strSession
	msgMap["privilige"] = s.Privilege
	return CreateAnyMsg(msgMap)
}

func updataSession(v *db.SUser) {
	_, err := db.G_OrmDB.Update(v, "session")
	if err != nil {
		x.PrintErr(err)
	}
}
