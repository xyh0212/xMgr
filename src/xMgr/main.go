package main

import (
	//"sync"
	"fmt"
	"xMgr/cmd"
	//"xMgr/db"
	"xMgr/sys"
	"xMgr/websocket"
	"xsw/go_pub/x"
)

func RcvMsgRoutine() {
	defer sys.DelRoutine()
	var err error
	var xerr *x.Error
	x.PrintInfo("RcvMsgRoutine begin")
	for {
		msg, bStop := websocket.PeekRcvMsg()
		if bStop {
			x.PrintInfo("PeekRcvMsg stop")
			break
		}
		if msg == nil {
			continue
		}
		xerr = cmd.DoCmd(string(msg.ByteMsg), msg.ID)
		if err != nil {
			x.PrintErr(xerr)
		}
	}
	x.PrintInfo("RcvMsgRoutine exit")
}

func main() {
	fmt.Println("112222")
	defer x.FiniX()
	x.SetLogLev(x.LOGLEV_DBG)
	x.SetNagle(true)
	//	err := db.InitDB()
	//	if err != nil {
	//		x.PrintErr(err)
	//		return
	//	}
	x.PrintInfo("start")

	sys.AddRoutine()
	go RcvMsgRoutine()

	err := websocket.InitWebsocket()
	if err != nil {
		x.LogErr(err.GetStr())
	}
	sys.WaitAllRoutineQuit()
	x.PrintInfo("end")
}
