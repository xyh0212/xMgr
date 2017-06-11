package x

import (
	"fmt"
	"os"
	"time"
)

const (
	LOGLEV_DBG = 1
	LOGLEV_ACCESS
	LOGLEV_INFO
	LOGLEV_WARN
	LOGLEV_ERR
)

var (
	dbg_log  logObj
	info_log logObj
	err_log  logObj

	access_log  logObj
	db_info_log logObj
	db_err_log  logObj

	log_lev int
)

func init_log() {
	log_lev = LOGLEV_INFO
	Log_Run()
	var err error

	info_log.bIsWriteDB = true
	info_log.bIsWriteFile = true
	info_log.myLogLev = LOGLEV_INFO
	info_log.myLogHandle, err = NewLogger("log/info", 0, "info")
	if err != nil {
		fmt.Println(err)
	}

	db_info_log.bIsWriteDB = true
	db_info_log.bIsWriteFile = true
	db_info_log.myLogLev = LOGLEV_INFO
	db_info_log.myLogHandle, err = NewLogger("log/db_info", 0, "db_info")
	if err != nil {
		fmt.Println(err)
	}

	db_err_log.bIsWriteFile = true
	db_err_log.myLogLev = LOGLEV_ERR
	db_err_log.myLogHandle, err = NewLogger("log/db_err", 0, "db_err")
	if err != nil {
		fmt.Println(err)
	}

	err_log.bIsWriteDB = true
	err_log.bIsWriteFile = true
	err_log.myLogLev = LOGLEV_ERR
	err_log.myLogHandle, err = NewLogger("log/err", 0, "err")
	if err != nil {
		fmt.Println(err)
	}

	access_log.bIsWriteDB = true
	access_log.bIsWriteFile = true
	access_log.myLogLev = LOGLEV_ACCESS
	access_log.myLogHandle, err = NewLogger("log/access", 0, "access")
	if err != nil {
		fmt.Println(err)
	}

	dbg_log.myLogLev = LOGLEV_DBG
	dbg_log.myLogHandle, err = NewLogger("log/dbg", 0, "dbg")
	if err != nil {
		fmt.Println(err)
	}

	InitLogDB()
}

/////////////////////////////////////////////
type logObj struct {
	myLogLev     int
	myLogHandle  int
	bIsWriteFile bool
	bIsWriteDB   bool
}

func (self *logObj) isAllowLog() bool {
	return log_lev <= self.myLogLev
}

func (self *logObj) Print(args ...interface{}) {
	if !self.isAllowLog() {
		return
	}

	str := fmt.Sprint(args...)
	print_str(str)
	if self.bIsWriteFile {
		self.LogString(str)
	}
}

func (self *logObj) LogString(args ...interface{}) {
	if !self.isAllowLog() {
		return
	}

	str := fmt.Sprintln(args...)
	WriteLog(self, &str)
}

func (self *logObj) LogF(format string, a ...interface{}) {
	if !self.isAllowLog() {
		return
	}

	str := fmt.Sprintf(format, a...)
	WriteLog(self, &str)
}

func (self *logObj) PrintF(format string, a ...interface{}) {
	if !self.isAllowLog() {
		return
	}

	str := fmt.Sprintf(format, a...)
	print_str(str)
	if self.bIsWriteFile {
		self.LogString(str)
	}
}

func (self *logObj) DumpF(format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	DumpLog(self, &str)
}

/////////////////////////////////////////////

func SetLogLev(nLev int) {
	log_lev = nLev
}

func FiniX() {
	log_Fini()
}

func print_str(str string) {
	fmt.Println(time.Now().Format("15:04:05  ") + str)
}

func PrintInfo(args ...interface{}) {
	info_log.Print(args...)
}

func PrintErr(args ...interface{}) {
	err_log.Print(args...)
}

func PrintInfoF(format string, a ...interface{}) {
	info_log.PrintF(format, a)
}

func PrintDbg(args ...interface{}) {
	dbg_log.Print(args...)
}

func PrintDbgF(format string, a ...interface{}) {
	dbg_log.PrintF(format, a...)
}

func PrintErrF(format string, a ...interface{}) {
	err_log.PrintF(format, a...)
}

func LogInfo(args ...interface{}) {
	info_log.LogString(args...)
}

func LogAccess(args ...interface{}) {
	access_log.LogString(args...)
}

func LogDbg(args ...interface{}) {
	dbg_log.LogString(args...)
}

func LogInfoF(format string, a ...interface{}) {
	info_log.PrintF(format, a...)
}

func LogDbgF(format string, a ...interface{}) {
	dbg_log.PrintF(format, a...)
}

func LogErr(args ...interface{}) {
	err_log.LogString(args...)
}

func LogErrF(format string, a ...interface{}) {
	err_log.LogF(format, a...)
}

func DumpErrF(format string, a ...interface{}) {
	err_log.DumpF(format, a...)
}

func CheckErr(err error) {
	if err != nil {
		str := string(err.Error())
		LogErr(str)
		os.Exit(1)
	}
}
