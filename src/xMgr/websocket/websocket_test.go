package websocket

import (
	"fmt"
	//	"github.com/astaxie/beego/orm"
	"io"
	"os"
	"testing"
	"xMgr/db"
	"xsw/go_pub/x"
)

func Test_Insert(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	db.LoadTable()
	//str := `upload/2016-02/1-02-29_20-00-05-test.txt`
	/*
		x.PrintInfo(CreateSavePath(str))
		x.ExcuSqlFile(db.CreateDBXmlNode(2), "../../../bin/upload/2016-02/xyh.txt")
		str = "dalfg?sdkg=upload/2016-02/xyh.txt&id=lglsdl&name=gjlg&session=N0xvMphFwTW6qCHLS3e2I9YgCcgCzAQN"
		pmap, err1 := UrlByMap(str)
		if err1 != nil {
			x.PrintDbg(err1)
		}
		x.PrintDbg(pmap)
		x.PrintDbg(GetPath("xyh"))
	*/
}

func copyFile(dst io.Reader, src string) (w int64, err error) {

	dstFile, err := os.Create(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	return io.Copy(dstFile, dst)
}
func Test_HttpCmd(t *testing.T) {
	sql := `insert into d_package_item(user_id,package,title_txt,desc_txt,itemtype_main,itemtype_param,num,time_t)`
	sql += ` values(1233, '1', '标题文字', '内容文字', '65665', '466', 123, UNIX_TIMESTAMP())`
	x.PrintDbg(sql)
	sql = `insert into d_package_item(user_id,package,title_id,desc_id,  itemtype_main,itemtype_param,num,time_t,title_txt,desc_txt)`
	sql += `SELECT 11,1,0,0,'%v','%v','%v',%v,'%v','%v'`
	sql = fmt.Sprintf(sql, "1111", "1111", 132, "UNIX_TIMESTAMP()", "pk", "1313")
	gsinfo, ok := db.G_GSInfoMap[1]
	if !ok {
		x.LogInfo(`db.G_GSInfoMap[1]的key不存在`)
	}
	err := gsinfo.ExecSql(sql)
	x.PrintDbg(err)

}
func Test_PlayerList(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	db.LoadTable()
	GSInfo, ok := db.G_GSInfoMap[1]
	if !ok {
		x.PrintDbg("key bu cun zai")
		return
	}
	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user where channel_accountname='%v'`
	sql = fmt.Sprintf(sql, "xyh")
	GSInfo.InitPDB()
	rows, xerr2 := x.CreateRecordSet(GSInfo.PDB, sql)
	//rows, err1 := GSInfo.QuerySql(sql)
	if xerr2 != nil {
		x.PrintDbg(xerr2)
	}

	sql = `SELECT count(0) from d_user where channel=104 limit 0,50`
	rows1, _ := GSInfo.QuerySql(sql)
	x.PrintDbg(*rows1.ToJson())
	x.PrintInfo(rows.ToMap())
	sql = `SELECT d.id,d.producer_order,u.name,FROM_UNIXTIME(d.create_time) from d_moneycard d LEFT JOIN d_user u on d.role_id=u.id `
	sql += `where d.sdk_rmb=100 and d.producer_order='1' and d.create_time='1455930097' and u.name='蛟龙传奇'`
	rows, xerr2 = GSInfo.QuerySql(sql)
	x.PrintInfo(*rows.ToTableJson())
	sql = `abc`
	x.PrintInfo(len(sql))
}
