package x

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

//flag
const (
	FLAG_TIME          = 1 << iota // 2015-04-04/14:01:01 文件内容
	FLAG_FILENAME_TIME             // 2015-04-04 文件名
)

const (
	NULL_STR = string("")
)

type LogMsg struct {
	pLogObj  *logObj
	strLog   *string
	cur_time time.Time
}

type LogMgr struct {
	nLogTypeSeed int
	chanLog      chan LogMsg
	mapLogObj    map[int]*Logger //key:nLogType
	//日志chan队列的buffer长度，建议不要少于1024，不多于102400，最长：2147483648
	chanLogBuffSize int
}

//日志全局变量
//全局once
var g_Once_V sync.Once
var g_LogMgr *LogMgr
var g_wait_quit sync.WaitGroup
var g_bNagle bool

type Logger struct {
	file_path string
	Log_type  int
	Prefix    string //前缀
	Flag      int
	w         *bufio.Writer
}

/////////////////////////////////////////////////////////////////////////
//内部函数

func SetNagle(bNagle bool) {
	g_bNagle = bNagle
}

func Log_Run() {
	//初始化全局变量
	if g_LogMgr == nil {
		g_LogMgr = new(LogMgr)
	}
	//调用初始化操作，全局只运行一次
	g_Once_V.Do(log_Init)
	//内部日志工作协程
	go func() {
		var log_msg LogMsg
		stop_flag := false
		defer g_wait_quit.Done()

		for {
			if stop_flag {
				break
			}

			select {
			//有日志
			case log_msg = <-g_LogMgr.chanLog:
				{
					if log_msg.pLogObj == nil {
						stop_flag = true
						break
					}
					log_Write(&log_msg, false)
				}
				break
				//超时，避免过忙
			case <-time.After(100 * time.Millisecond):
				break
			}
		}
	}()

}

func NewLogger(file_name string, flag int, prefix string) (int, error) {
	if g_LogMgr == nil {
		return -1, errors.New("g_LogMgr is nil")
	}

	if flag == 0 {
		flag = FLAG_TIME | FLAG_FILENAME_TIME
	}

	if flag&FLAG_FILENAME_TIME != 0 {
		file_name = file_name + " " + time.Now().Format("2006-01-02") + ".log"
	} else {
		file_name = file_name + ".log"
	}

	//创建&打开新日志文件
	Log_type := g_LogMgr.createLogType()

	if g_LogMgr.mapLogObj[Log_type] != nil {
		return -1, errors.New(fmt.Sprintf("NewLogger %s repeat", file_name))
	}

	myLog := &Logger{}
	myLog.Log_type = Log_type
	myLog.Flag = flag
	myLog.Prefix = prefix
	myLog.file_path = file_name
	g_LogMgr.mapLogObj[myLog.Log_type] = myLog
	return Log_type, nil
}

// func WriteLogFmt(p *logObj, format string, a ...interface{}) error {
// 	str := fmt.Sprintf(format, a...)
// 	return sync_msg(log_type, &str)
// }

func WriteLog(p *logObj, str *string) error {
	return sync_msg(p, str)
}

func DumpLog(p *logObj, str *string) error {
	file_name, file_line, err := GetSourceFileLine(3)
	if err != nil {
		return err
	}
	str2 := file_name + ":" + strconv.Itoa(file_line) + "," + *str

	return sync_msg(p, &str2)
}

/////////////////////////////////////////////////////////////////////////
//内部函数

func (l *LogMgr) createLogType() int {
	l.nLogTypeSeed += 1
	return l.nLogTypeSeed
}

func log_Write(log_msg *LogMsg, bFlush bool) error {
	if log_msg == nil {
		panic("nil log_msg")
	} else if log_msg.pLogObj == nil {
		panic("nil pLogObj")
	}

	if g_bLog2DB && log_msg.pLogObj.bIsWriteDB {
		return log_Write_DB(log_msg, bFlush)
	}
	return log_Write_File(log_msg, bFlush)
}

func log_Write_File(log_msg *LogMsg, bFlush bool) error {
	if log_msg == nil {
		panic("nil log_msg")
	} else if log_msg.pLogObj == nil {
		panic("nil pLogObj")
	}
	nLogType := log_msg.pLogObj.myLogLev

	var err error

	log_obj := g_LogMgr.mapLogObj[nLogType]
	if log_obj == nil {
		return errors.New(fmt.Sprintf("log type error:%d", nLogType))
	}

	if log_obj.w == nil {
		var Log_fd *os.File
		os.Mkdir("log", os.ModeDir)
		Log_fd, err = os.OpenFile(log_obj.file_path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return errors.New(fmt.Sprintf("Open log file %s fail", log_obj.file_path))
		}
		Log_fd.Seek(0, os.SEEK_END)
		//创建bufio的Writer
		log_obj.w = bufio.NewWriter(Log_fd)
	}

	var strLog string
	if log_obj.Flag&FLAG_TIME != 0 {
		strLog = strLog + log_msg.cur_time.Format("2006-01-02 15:04:05,")
	}

	strLog = strLog + *log_msg.strLog

	_, err = log_obj.w.WriteString(strLog)
	if err != nil {
		return err
	}

	if bFlush {
		log_obj.w.Flush()
	}

	return nil
}

func log_Init() {
	//初始化全局变量
	if g_LogMgr == nil {
		g_LogMgr = new(LogMgr)
		g_wait_quit.Add(1)
	}

	//设置日志channel buffer
	if g_LogMgr.chanLogBuffSize <= 0 {
		g_LogMgr.chanLogBuffSize = 1024
	}
	g_LogMgr.chanLog = make(chan LogMsg, g_LogMgr.chanLogBuffSize)
	g_LogMgr.mapLogObj = make(map[int]*Logger, 3)
}

func log_Fini() {
	var str string
	str = ""
	sync_msg(nil, &str)
	g_wait_quit.Wait()

	//flush file
	for _, value := range g_LogMgr.mapLogObj {
		if value.w != nil {
			value.w.Flush()
		}
	}

}

//------------------------
//   logger内部使用方法

// 写入日志到channel
func sync_msg(p *logObj, log_str *string) error {
	log_msg_data := LogMsg{
		pLogObj:  p,
		strLog:   log_str,
		cur_time: time.Now(),
	}
	if g_bNagle {
		//写消息到channel
		g_LogMgr.chanLog <- log_msg_data
	} else {
		log_Write(&log_msg_data, true)
	}
	return nil
}
