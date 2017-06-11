package cmd

import (
	//"github.com/server-nado/orm"
	//"xMgr/db"
	//"net/url"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strings"
	//"strconv"
	//"encoding/json"
	//"reflect"
	//"strconv"
	//"time"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

//删除
//obj := new(db.SGs)
//obj.Id = 1
//Delete(obj)
func Delete(obj interface{}) ([]byte, *x.Error) {
	_, err1 := db.G_OrmDB.Delete(obj)
	if err1 != nil {
		return CreateSqlErrMsg(x.XErr(err1))
	}
	return nil, nil
}

//修改
//obj := new(db.SGs)
//obj.Name = "xyh"
//Updata(obj)
func Updata(arrStr []string, obj interface{}) ([]byte, *x.Error) {
	value, err := sys.ArrString2Map(arrStr)
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	sys.Map2Struct(value, obj)
	_, err1 := db.G_OrmDB.Update(obj)
	if err1 != nil {
		return CreateSqlErrMsg(x.XErr(err1))
	}
	return nil, nil
}

//添加
//obj := new(db.SGs)
//var arr = []string{"name:dfg","name2:dlf","gs_group_id:14","gs_id:12","ip:6354613"}
//Insert(arr,obj)
func Insert(arrStr []string, obj interface{}) ([]byte, *x.Error) {
	value, err := sys.ArrString2Map(arrStr)
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	sys.Map2Struct(value, obj)
	_, err1 := db.G_OrmDB.Insert(obj)
	if err1 != nil {
		return CreateSqlErrMsg(x.XErr(err1))
	}
	//db.TbGseMap[]obj.(db.SGs)
	return nil, nil
}
func CreateSqlErrMsg(perr *x.Error) ([]byte, *x.Error) {
	var js *simplejson.Json
	var err error
	js, err = simplejson.NewJson([]byte(`{}`))
	if err != nil {
		return nil, x.XErr(err)
	}
	js.Set("cmd", 10)
	js.Set("info", perr.GetStr())
	var strMsg []byte
	strMsg, err = js.Encode()
	return strMsg, x.XErr(err)
}

//返回值字符串类型为{"userjson":[["",""],["",""]]}
func getGsTableJStr() string {
	var sumstr string
	sumstr = `{"userjson":[`
	for _, value := range db.TbGseMap {
		sumstr += sys.ObjToString(value) + `,`
	}
	sumstr = strings.TrimRight(sumstr, ",")
	sumstr += `]}`
	return sumstr
}

//返回值字符串类型为{"userjson":[["",""],["",""]]}
func TbGseMapToString() string {
	var sumstr string
	sumstr = `{"userjson":[`
	for _, value := range db.TbGseMap {
		sumstr += sys.ObjToString(value) + `,`
	}
	sumstr = strings.TrimRight(sumstr, ",")
	sumstr += `]`
	fmt.Println(sumstr)
	return sumstr
}
func checkParam(cmd CmdInfo, str string) ([]byte, *x.Error, *string) {
	id, _ := cmd.jsParam.CheckGet(str)
	if id == nil {
		msg, _ := CreateErrMsg(sys.CODE_NEED_ID_PARAM)
		return msg, x.XErrStr("name not exist"), nil
	}
	pid, err := id.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron"), nil
	}
	return nil, nil, &pid
}
