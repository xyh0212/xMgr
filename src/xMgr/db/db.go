package db

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"xsw/go_pub/x"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	G_OrmDB orm.Ormer
)

func InitDB() *x.Error {
	err2 := ConnectOrmDB("config/db.xml")
	if err2 != nil {
		return err2
	}

	err2 = LoadTable()
	if err2 != nil {
		return err2
	}

	x.PrintInfo("InitDB suc")
	fmt.Println(x.Any2JsonArrayString(G_GSInfoMap))

	return nil
}

////////////////////////////////////////////////////////////////////
//
func CountTable(strTable string, db *sql.DB) (n int, err error) {
	// 从数据库加载
	if db == nil {
		x.LogErr("db is nil")
		return 0, nil
	}
	sql := fmt.Sprintf("select count(0) from %s limit 1", strTable)
	row := db.QueryRow(sql)

	err = row.Scan(&n)
	return n, err
}

////////////////////////////////////////////////////////////////////
//
func GetIDAndName(strTable string, db *sql.DB, filedName ...string) (*string, *x.Error) {
	if db == nil {
		x.LogErr("db is nil")
		return nil, x.XErrStr("db is nil")
	}
	var str string
	for _, v := range filedName {
		str += v + `,`
	}
	str = strings.TrimRight(str, ",")
	sql := `select ` + str + ` from ` + strTable
	//	rows, err := x.CreateRecordSet(db, sql)
	//if err != nil {
	//	return nil, err
	//}
	//	pstr := rows.ToJson()
	return &sql, nil
}

////////////////////////////////////////////////////////////////////

func getDBConnectString(strXmlPath string) (string, x.DBXmlNode, *x.Error) {
	var xmlDB x.DBXmlNode
	var err error

	// db.xml
	var strXml []byte
	if strXml, err = ioutil.ReadFile(strXmlPath); err != nil {
		return "", xmlDB, x.XErr(err)
	}
	if err = xml.Unmarshal(strXml, &xmlDB); err != nil {
		return "", xmlDB, x.XErr(err)
	}

	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		xmlDB.User, xmlDB.Pass, xmlDB.Host, xmlDB.Port, xmlDB.Db, xmlDB.Extra)
	return str, xmlDB, nil
}

func ConnectOrmDB(strXmlPath string) *x.Error {
	if G_OrmDB != nil {
		return nil
	}
	var err error

	str, xmlDB, xerr := getDBConnectString(strXmlPath)
	if xerr != nil {
		return xerr
	}

	orm.RegisterModel(new(SDbGroup))
	//x.PrintDbg("/////////////////////////////////////////////////////")
	orm.RegisterModel(new(SGsGroup))
	orm.RegisterModel(new(SCodeString))
	orm.RegisterModel(new(SGs))
	orm.RegisterModel(new(SUser))
	orm.RegisterModel(new(SUserGroup))
	orm.RegisterModel(new(SSdkType))
	orm.RegisterModel(new(SItemType))
	//orm.RegisterModel(new(TbCodeString2))

	err = orm.RegisterDataBase("default", "mysql", str) //设置conn中的数据库为默认使用数据库
	if err != nil {
		return x.XErr(err)
	}

	err = orm.RunSyncdb("default", false, true) //第二个true:强制创建表 第三个true:显示详细信息
	if err != nil {
		return x.XErr(err)
	}

	G_OrmDB = orm.NewOrm()
	xerr = SetDBCharacter(G_OrmDB)
	if xerr != nil {
		return xerr
	}
	err1 := initDBData(&xmlDB)
	if err1 != nil {
		return err1
	}
	return nil
}
func initDBData(xmlDB *x.DBXmlNode) *x.Error {
	db, err := orm.GetDB()
	if err != nil {
		return x.XErr(err)
	}
	userCount, err1 := x.CountTable(db, `s_user`)
	if err1 != nil {
		return err1
	}
	if userCount > 0 {
		return nil
	}
	_, err5 := x.ExcuSqlFile(xmlDB, `config/install.sql`) //config/install.sql
	if err5 != nil {
		return err5
	}
	return nil

}
func SetDBCharacter(pOrm orm.Ormer) *x.Error {
	var err error
	// _, err = pOrm.Raw("SET NAMES UTF8").Exec()
	// if err != nil {
	// 	return x.XErr(err)
	// }

	// _, err = pOrm.Raw("set character_set_results=latin1").Exec()
	// if err != nil {
	// 	return x.XErr(err)
	// }

	_, err = pOrm.Raw("SET character_set_client = latin1").Exec()
	if err != nil {
		return x.XErr(err)
	}

	_, err = pOrm.Raw("SET character_set_connection = UTF8").Exec()
	if err != nil {
		return x.XErr(err)
	}

	_, err = pOrm.Raw("SET character_set_database = UTF8").Exec()
	if err != nil {
		return x.XErr(err)
	}

	_, err = pOrm.Raw("SET character_set_results = latin1").Exec()
	if err != nil {
		return x.XErr(err)
	}

	_, err = pOrm.Raw("SET character_set_server = UTF8").Exec()
	if err != nil {
		return x.XErr(err)
	}

	return nil
}

//id 为 gs的id
func CreateDBXmlNode(s_gs_id int) *x.DBXmlNode {
	var e x.DBXmlNode
	sGs, ok := TbGseMap[s_gs_id]
	if !ok {
		x.LogErr(`TbGseMap[id]的key找不到`)
		return nil
	}
	value, ok := TbDbGroupMap[sGs.DbGroupId]
	if !ok {
		x.LogErr(`TbDbGroupMap[sGs.DbGroupId]的key找不到`)
		return nil
	}
	e.Db = sGs.DbName
	e.Extra = value.GetExtra()
	e.Pass = value.DbPwd // "skiK@983.ff"
	e.Port = value.GetDBPort()
	e.User = value.DbUserName // "kofront"
	e.Host = value.DbIp       //"192.168.0.200"
	return &e
}

func SessionByUserName(session string) (string, *x.Error) {
	for _, value := range TbUserMap {
		if value.Session == session {
			return value.Name, nil
		}
	}
	return "", x.XErrStr("The user name has expired")
}
