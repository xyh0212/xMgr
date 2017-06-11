package x

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

var (
	g_Orm          orm.Ormer
	g_db           *sql.DB
	g_bLog2DB      bool
	g_bDelayInsert bool
)

func InitLogDB() {
	orm.RegisterModel(new(TB_d_log_str))
	g_bDelayInsert = true
}

func SetLog2DB(o orm.Ormer) *Error {
	var err error
	g_db, err = orm.GetDB()
	if err != nil {
		x := XErr(err)
		LogErr(x)
		return x
	}
	g_Orm = o
	g_bLog2DB = (g_Orm != nil)
	return nil
}

////////////////////////////////////////////////////////////////////
type TB_d_log_str struct {
	Id   int64 `pk`
	Type int
	T    time.Time ``
	Str  string    `orm:"size(500)"`
}

func (self *TB_d_log_str) TableName() string {

	return "d_log_str"
}

////////////////////////////////////////////////////////////////////
//insert log
func log_Write_DB(log_msg *LogMsg, bFlush bool) error {
	if log_msg == nil {
		panic("nil log_msg")
	} else if log_msg.pLogObj == nil {
		panic("nil pLogObj")
	}
	nLogType := log_msg.pLogObj.myLogLev

	var err error
	if g_Orm == nil {
		panic(errors.New("g_Orm==nil"))
	}

	// obj := new(TB_d_log_str)
	// obj.Type = log_msg.nLogType
	// obj.T = log_msg.cur_time
	// obj.Str = *log_msg.strLog

	// _, err = g_Orm.Insert(obj)
	var b []byte
	str1 := string(EscapeSql(b, *(log_msg.strLog)))
	if g_bDelayInsert {
		str1 = fmt.Sprintf("INSERT delayed INTO d_log_str (type, t, str) VALUES (%d, \"%s\", \"%s\")",
			nLogType, log_msg.cur_time.Format("2006-01-02 15:04:05"), str1)
	} else {
		str1 = fmt.Sprintf("INSERT INTO d_log_str (type, t, str) VALUES (%d, \"%s\", \"%s\")",
			nLogType, log_msg.cur_time.Format("2006-01-02 15:04:05"), str1)
	}
	_, err = g_db.Exec(str1)
	if err != nil {
		if g_bDelayInsert {
			g_bDelayInsert = false
		}
		db_err_log.LogString(fmt.Sprintln("Insert:%v", err))
	}

	return nil
}
