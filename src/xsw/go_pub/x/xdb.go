package x

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//	"io/ioutil"
	"bytes"
	"os/exec"
	"strconv"
)

type DBColumn struct {
	Value string //字段值
}

type DBXmlNode struct {
	Host  string `xml:"host"`
	Port  int    `xml:"port"`
	Db    string `xml:"db"`
	User  string `xml:"user"`
	Pass  string `xml:"pass"`
	Extra string `xml:"extra"`
}

////////////////////////////////////////////
//DBRow
type DBRow struct {
	ColumnArray []*DBColumn //各个字段的值
}

func (self *DBRow) init(nColumnNum int) {
	self.ColumnArray = make([]*DBColumn, nColumnNum)
}

func (self *DBRow) setColumnVal(nCol int, val string) *Error {
	if len(self.ColumnArray) < (nCol + 1) {
		return XErrStr("column out of index")
	}
	col := new(DBColumn)
	col.Value = val
	self.ColumnArray[nCol] = col
	return nil
}

func (self *DBRow) GetString(nFieldIndex int) (*string, *Error) {
	if len(self.ColumnArray) <= nFieldIndex {
		xerr := XErrStr("index much more")
		LogErr(xerr)
		return nil, xerr
	}
	return &self.ColumnArray[nFieldIndex].Value, nil
}

func (self *DBRow) GetInt(nFieldIndex int) int64 {
	if len(self.ColumnArray) <= nFieldIndex {
		xerr := XErrStr("index much more")
		LogErr(xerr)
		return 0
	}
	n, err := AtoI64(self.ColumnArray[nFieldIndex].Value)
	if err != nil {
		xerr := XErr(err)
		LogErr(xerr)
		return 0
	}
	return n
}

// [1,2,3]
func (self *DBRow) ToJsonString() string {
	nColNum := len(self.ColumnArray)
	var str string
	str += "["
	for i := 0; i < nColNum; i++ {
		if i > 0 {
			str += ","
		}
		// 有些字符需要转义
		str = str + "\"" + string(EscapeString([]byte{}, self.ColumnArray[i].Value)) + "\""
	}
	str += "]"
	return str
}

//////////////////////////////////////////////
//
func ExcuSqlFile(pDBInfo *DBXmlNode, strFile string) (str string, xerr *Error) {
	cmd := exec.Command(
		"cmd",
		"/C",
		"mysql",
		"-h"+pDBInfo.Host,
		"-u"+pDBInfo.User,
		"-p"+pDBInfo.Pass,
		"-P"+strconv.Itoa(pDBInfo.Port),
		pDBInfo.Db,
		"<",
		strFile)

	var out bytes.Buffer
	cmd.Stdout = &out //输出
	cmd.Stderr = &out

	err := cmd.Run()
	str = out.String()
	if err != nil {
		xerr = XErrStr(str)
		db_err_log.LogString(err.Error(), "\n", str)
		return str, xerr
	}
	db_info_log.LogString(str)
	return str, nil
}

func ExcuSqlFileWithLog(pDBInfo *DBXmlNode, strFile string) (xerr *Error) {
	cmd := exec.Command(
		"cmd",
		"/C",
		"mysql",
		"-h"+pDBInfo.Host,
		"-u"+pDBInfo.User,
		"-p"+pDBInfo.Pass,
		"-P"+strconv.Itoa(pDBInfo.Port),
		pDBInfo.Db,
		"<",
		strFile)

	var out bytes.Buffer
	cmd.Stdout = &out //输出
	cmd.Stderr = &out

	err := cmd.Run()
	str := out.String()
	if err != nil {
		xerr = XErrStr(str)
		db_err_log.LogString(err.Error(), "\n", str)
		return xerr
	}
	db_info_log.LogString(str)
	return nil
}

func CountTable(db *sql.DB, strTable string) (n int, xerr *Error) {
	var err error
	// 从数据库加载
	if db == nil {
		LogErr("db is nil")
		return 0, nil
	}

	sql := fmt.Sprintf("select count(0) from %s limit 1", strTable)
	row := db.QueryRow(sql)

	err = row.Scan(&n)
	if err != nil {
		xerr = XErr(err)
		LogErr(xerr)
		return 0, xerr
	}
	return n, nil
}

func SetDBCharacter(db *sql.DB) *Error {
	var err error
	// _, err = db.("SET NAMES UTF8")
	// if err != nil {
	// 	return XErr(err)
	// }

	// _, err = db.("set character_set_results=latin1")
	// if err != nil {
	// 	return XErr(err)
	// }

	_, err = db.Exec("SET character_set_client = latin1")
	if err != nil {
		return XErr(err)
	}

	_, err = db.Exec("SET character_set_connection = UTF8")
	if err != nil {
		return XErr(err)
	}

	_, err = db.Exec("SET character_set_database = UTF8")
	if err != nil {
		return XErr(err)
	}

	_, err = db.Exec("SET character_set_results = latin1")
	if err != nil {
		return XErr(err)
	}

	_, err = db.Exec("SET character_set_server = UTF8")
	if err != nil {
		return XErr(err)
	}

	return nil
}
