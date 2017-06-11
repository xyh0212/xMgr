package cmd

import (
	//"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	//"strconv"
	//"strings"
	//"time"
	"xMgr/db"
	"xMgr/sys"
	//"fmt"
	//"reflect"
	"xsw/go_pub/x"
)

// {"cmd":12,"param":{"info":"["字段名:值","字段名:值"]"}}
func DoAddGSCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoAddGSCmd")
	info, _ := cmd.jsParam.CheckGet("info")
	if info == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_INFO_PARAM)
		return msg, x.XErrStr("info not exist")
	}

	infoArrStr, err := info.StringArray()
	if err != nil {
		x.PrintDbg(err)
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	sgs := new(db.SGs)
	//x.PrintDbg(infoArrStr)
	msg, err1 := Insert(infoArrStr, sgs)
	if err1 == nil {
		db.TbGseMap[sgs.Id] = sgs
		return CreateSucceedMsg(sys.CMD_ADD_GS)
	}
	return msg, err1
}
