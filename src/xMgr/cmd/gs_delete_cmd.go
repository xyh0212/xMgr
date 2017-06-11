package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":13,"param":{"id":12}}
func DoDeleteGSCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoDeleteGSCmd")
	id, _ := cmd.jsParam.CheckGet("id")
	if id == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_INFO_PARAM)
		return msg, x.XErrStr("id not exist")
	}

	pId, err := id.Int()
	if err != nil {
		x.PrintDbg(err)
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	obj := new(db.SGs)
	obj.Id = pId
	msg, err1 := Delete(obj)
	if err1 == nil {
		delete(db.TbGseMap, pId)
		return CreateSucceedMsg(sys.CMD_DELETE_GS)
	}
	return msg, err1
}
