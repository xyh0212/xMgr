package cmd

import (
	//"github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"xMgr/db"
	"xMgr/sys"
	"xMgr/user"
	"xsw/go_pub/x"
)

// {"cmd":5,"param":{"session":"xxxxx"}}
func DoCheckSessionCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoCheckSessionCmd")
	///////////////////////////////////
	session, _ := cmd.jsParam.CheckGet("session")
	if session == nil {
		msg, _ := CreateErrMsg(sys.CODE_SESSION_NO_EXIST)
		return msg, x.XErrStr("session not exist")
	}
	psession, err := session.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	////////////////////////////////////////
	//

	if !checkSession(psession, cid) {
		msg, _ := CreateErrMsg(sys.CODE_SESSION_ERR)
		return msg, x.XErrStr("session erron")
	}

	return CreateSucceedMsg(sys.CMD_CHECK_SESSION)

}

func checkSession(session string, cid int) bool {
	x.PrintDbg(session)
	for _, value := range db.TbUserMap {
		x.PrintDbg(value)
		if session == value.Session {
			user.GSessionMgr.AddSession(cid, session)
			return true
		}
	}
	return false
}
