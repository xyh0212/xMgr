package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":2}
func DoQueryGSCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoQueryGSCmd")
	msg, err := CreateInitCmdMsg1()
	if err != nil {
		//x.PrintInfo("我是err", err)
		return nil, err
	}
	//x.PrintInfo("hello", string(msg))
	return msg, nil

}

// {"cmd":2}
func CreateInitCmdMsg1() ([]byte, *x.Error) {
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_QUERY_GS
	msgMap["info"] = db.GetAllProductJstr()
	//x.PrintDbg("我啥都啥都分", db.GetAllProductJstr())
	return CreateAnyMsg(msgMap)
}
