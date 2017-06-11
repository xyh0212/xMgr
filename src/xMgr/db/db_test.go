package db

import (
	//"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"os"
	"strconv"
	"testing"
	"xsw/go_pub/x"
)

func Test_DB(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error

	err = ConnectOrmDB(`../../../bin/config/db.xml`)

	if err != nil {

		t.Error("ConnectDB:", err)
	}
	//加载数据
	LoadTable()
	//打印所有产品json字符串
	x.PrintInfo(GetAllProductJstr())
}

func Test_DB_QueryAny(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}
	db, _ := orm.GetDB("default")
	rows, xerr2 := x.CreateRecordSet(db, "SELECT * FROM s_gs")
	if xerr2 != nil {
		t.Error(xerr2)
		return
	}
	fmt.Println(rows.ToString())
}
func Test_DB_UpdataAny(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error

	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)

	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}
	// G_OrmDB.Raw("SET NAMES UTF8").Exec()
	// G_OrmDB.Raw("set character_set_results=latin1").Exec()

	// character_set_client
	// character_set_connection
	// character_set_database
	// character_set_results
	// character_set_server
	// character_set_system

	db, _ := orm.GetDB()
	// db.Exec("SET character_set_server = latin1")
	// db.Exec("SET character_set_client = latin1")
	// db.Exec("SET character_set_connection = UTF8")
	// db.Exec("SET character_set_database = UTF8")
	// db.Exec("SET character_set_results = latin1")
	// db.Exec("SET character_set_server = UTF8")
	// db.Exec("SET character_set_system = UTF8")

	db.Exec("update s_user set name = '1名3' where id = 2")

	// G_OrmDB.Raw("update s_user set name = '名2' where id = 1").Exec()
	// G_OrmDB.Raw("insert into d_session(id,session,user_id) values(1,sdfsaf,1)").Exec()
	// G_OrmDB.Raw("insert into s_user(id,name) values(2,'名')").Exec()
	// fmt.Println("hello")
}
func Test_RawDB(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	var err error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}

	var str string
	G_OrmDB.Raw("SELECT str from s_code_string where id=1000 limit 1").Exec()
	///row := g_OrmDB.QueryRow("SELECT str from s_code_string where id=1000 limit 1")
	//err = row.Scan(&str)

	file, err := os.Create("x.txt")
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(row)
	fmt.Println(str)
	file.Write(([]byte)(str))
}

func Test_UrlEncode(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	str := "a b+\""
	str1 := url.QueryEscape(str)
	str2, _ := url.QueryUnescape(str1)
	fmt.Println(str1)
	fmt.Println(str2)
}
func Test_Add_DB_User(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	user := new(SUser)
	user.Name = "名2"
	user.Pwd = "12223"
	//_, _, erre := user.Objects(user).Save()

}
func Test_GS(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	//getGsTableJStr()
}
func getGsTableJStr() string {
	var str string
	str = `{"userjson":[`
	lenth := len(TbGseMap)
	i := 0
	for _, value := range TbGseMap {
		if i == lenth-1 {
			str += getGsRowStr(value)
		} else {
			str += getGsRowStr(value) + ","
		}

		i++
	}
	str += `]}`
	x.PrintDbg(str)
	return str
}

func getGsRowStr(value *SGs) string {
	var str string
	str = `[`
	str += `"` + strconv.Itoa(value.Id) + `",`
	str += `"` + value.Name + `",`
	str += `"` + strconv.Itoa(value.GsGroupId) + `",`
	str += `"` + value.Ip + `",`
	str += `"` + strconv.Itoa(value.GsPort) + `",`
	str += `"` + strconv.Itoa(value.DbGroupId) + `",`
	str += `"` + value.GetDBName() + `",`
	str += `"` + strconv.Itoa(value.GsId) + `"`
	str += `]`
	return str
}
func Test_DB_Updata(t *testing.T) {

	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}
	//update
	testUser := new(SUser)
	testUser.Id = 2
	testUser.Name = "123"
	num, err := G_OrmDB.Update(testUser)
	if err == nil {
		x.PrintDbg(err)
	} else {
		x.PrintDbg(num)
	}
}
func Test_InitPDB(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	} /*
		LoadTable()
		gsInfo := G_GSInfoMap[1]
		gsInfo.InitPDB()
		sql := "select id,name,level from d_user where name='xyh'"
		//n, _ := CountTable("d_user", gsInfo.PDB)

		rows, xerr2 := x.CreateRecordSet(gsInfo.PDB, sql)
		if xerr2 != nil {
			x.PrintDbg(xerr2)
		}
		str := rows.ToJson()
		db, _ := orm.GetDB()
		j, _ := GetIDAndName("s_item_type", db, "id", "name")
		x.PrintDbg(*str)
		x.PrintDbg(*j)
		x.PrintDbg(DropListMap[1])
	*/

}
func Test_ExecSql(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}
	LoadTable()
	gsInfo := G_GSInfoMap[1]
	err := gsInfo.InitPDB()
	if err != nil {
		x.PrintDbg(err)
		return
	}
	sql := `insert into user(id,name,test)values('56','231','1')`
	err = gsInfo.ExecSql(sql)
	x.PrintDbg(err)

}
func Test_QuerySqlTable(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}
	LoadTable()
	gsInfo, ok := G_GSInfoMap[1]
	if !ok {
		x.PrintInfo(`G_GSInfoMap的key不存在`)
	}
	gsInfo.InitPDB() //xsw3
	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user where name='a1613dfh'`
	x.PrintDbg(sql)
	rows, err := x.CreateRecordSet(gsInfo.PDB, sql)
	if err != nil {
		x.PrintDbg(err)
		//w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err.GetStr()))
		return
	}
	s := rows.ToTableJson()
	x.PrintDbg(*s)

}

func Test_InitDbTable(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var xerr *x.Error
	xerr = ConnectOrmDB(`../../../bin/config/db.xml`)
	if xerr != nil {
		t.Error("connectRawDB:", xerr)
		return
	}
	LoadTable()

}
