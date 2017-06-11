package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	//"testing"
	//"fmt"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":2} 发送{json：[["",""],["",""]]}
func DoQueryGSManageCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoQueryGSManageCmd")

	msg, err := CreateInitCmdMsg2()
	if err != nil {
		return nil, err
	}
	return msg, nil
}

//得到排序的SGs的数组
func getArrGS() []*db.SGs {
	var arr []int
	var sgs []*db.SGs
	arr = make([]int, len(db.TbGseMap))
	sgs = make([]*db.SGs, len(db.TbGseMap))
	i := 0
	for key, _ := range db.TbGseMap {
		arr[i] = key
		i++
	}
	arr = sys.SortArryInt(arr)
	for index, value := range arr {
		sgs[index] = db.TbGseMap[value]
	}
	return sgs
}

// {"cmd":2,"param":{"tag":2}}
func CreateInitCmdMsg2() ([]byte, *x.Error) {
	str := x.Any2JsonArrayString(getArrGS())
	str = `{"json":` + str + `}`
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_QUERY_MANAGE_GS
	msgMap["info"] = str
	msgMap["field"] = sys.StructToString(new(db.SGs))
	return CreateAnyMsg(msgMap)
}
