package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":14,"param":{"id":"xxxxx"}}
func DoEditGSCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoEditGSCmd")
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
	strMap, _ := sys.ArrString2Map(infoArrStr)
	id, _ := strconv.Atoi(strMap["id"])
	msg, err1 := Updata(infoArrStr, db.TbGseMap[id])
	if err1 == nil {
		return CreateSucceedMsg(sys.CMD_EDIT_GS)
	}
	return msg, err1
}
