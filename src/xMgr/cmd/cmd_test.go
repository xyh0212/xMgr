package cmd

import (
	//"github.com/server-nado/orm"
	"testing"
	"xMgr/db"
	//"net/url"
	//"fmt"
	"reflect"
	//"strings"

	"fmt"
	//	"xMgr/websocket"

	//"time"
	"io/ioutil"
	"xMgr/sys"
	"xsw/go_pub/x"
)

func wTest_CreateMsg(t *testing.T) {
	var str []byte
	var err error
	//str, err = x.CreateErrorMsg("错误")
	if err != nil {
		//t.Error("CreateErrorMsg fail", err)
	} else {
		x.PrintInfo(string(str))
	}
	//x.PrintInfo("dsoifho/\ndpfih")
	//parserSql(`oisr\nh g`)
}

type User struct {
	Name string
	Age  int
	Id   string
}

func Test_Insert(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	obj := new(db.SUser)
	obj.Name = "xyh"
	_, err1 := db.G_OrmDB.Update(obj)
	if err1 != nil {
		x.PrintDbg(err)
	}
}
func Test_updata(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	obj := new(db.SUser)
	obj.Id = 2
	_, err1 := db.G_OrmDB.Update(obj)
	if err1 != nil {
		x.PrintDbg(err)
	}
}
func Test_delete(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	obj := new(db.SUser)
	obj.Id = 2
	_, err1 := db.G_OrmDB.Delete(obj)
	if err1 != nil {
		x.PrintDbg(err)
	}
}
func Test_GS(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	db.LoadTable()
	//fmt.Sprintln("fmt::%v", db.G_GSInfoMap)
	fmt.Println(x.Any2JsonArrayString(getArrGS()))

}

func Test_DecodeMsg(t *testing.T) {
	tonydon := &User{"TangXiaodong", 100, "0000123"}
	fmt.Println(sys.ObjToString(tonydon))

}
func Test_ObjToString(t *testing.T) {

	GSIN := new(db.SGs)
	StructToArrStr(GSIN)
	//str := `["id:56","name:dfg","name2:dlf","gs_group_id:14","gs_id:12","ip:6354613","gs_port:45","db_group_id:2","db_name:dfjh","start_time:2014-06-15 08:37:18"]`
	//Insert(str, GSIN)
	//value, _ := String2Map(str)
	//Map2Struct(value, GSIN)
	//_, erre := db.G_OrmDB.Insert(GSIN)
	//x.PrintDbg(erre)
	x.PrintDbg(sys.StructToString(new(db.SGs)))

}

//返回值为字符串["",""]
func StructToArrStr(pObj interface{}) []string {
	object := reflect.ValueOf(pObj)
	myref := object.Elem()
	typeOfType := myref.Type()
	var arr []string
	arr = make([]string, myref.NumField())
	var s string
	for i := 0; i < myref.NumField(); i++ {
		str := fmt.Sprintf("%v", typeOfType.Field(i).Name)
		s += `"` + x.SnakeString(str) + `",`
		//x.PrintDbg(myref.Field(i).Type())
		arr[i] = x.SnakeString(str)

	}
	//x.PrintDbg(x.SnakeString(s))
	return arr

}
func Test_ReadFile(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	db.LoadTable()
	sdbGroup := new(db.SDbGroup)
	//str := `INSERT INTO d_cmd VALUES (5,3);`
	//x.PrintInfo(str)
	var e x.DBXmlNode
	//e := new(x.DBXmlNode)
	e.Db = "xmgr"
	e.Extra = sdbGroup.GetExtra()
	e.Pass = "skiK@983.ff"
	e.Port = sdbGroup.GetDBPort()
	e.User = "kofront"
	e.Host = "192.168.0.200"
	//C:/Users/Administrator/Desktop/workplace/xMgr/bin/upload/2016-02/xyh.txt
	x.ExcuSqlFile(&e, `../../../bin/upload/xyh.txt`) //`../../../bin/upload/xyh.txt`
	x.PrintDbg(e)
	//IORead()
	//websocket.InitHttpUpload()  11:07:50  &{192.168.0.200 3306 xmgr 通用数据库配置 skiK@983.ff charset=utf8}
}

func IORead() {
	var out = []byte("osdhgiodarhgoeira快速登录积分换卡死的减肥hgpierhghepri")
	err2 := ioutil.WriteFile("log/output2.txt", out, 0666) //写入文件(字节数组)
	fmt.Println(err2)
}
