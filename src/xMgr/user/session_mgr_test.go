package user

import (
	// "fmt"
	"testing"
	"xMgr/db"
	"xsw/go_pub/x"
)

func Test_SessionMgr(t *testing.T) {
	x.PrintDbg("/////////////////////////////////////////////////////")
	var err *x.Error
	err = db.ConnectOrmDB(`../../../bin/config/db.xml`)
	if err != nil {
		t.Error("ConnectDB:", err)
	}
	db.LoadTable()
	//var err error
	//var str string
	//str, err = CreateSession(1, "许")
	//if err != nil {
	//	t.Errorf("CreateSession:%s", err.Error())
	//}
	//DelSession(str)
	//GSessionMgr.AddSession(1, "ldkfnl")
	//x.PrintDbg(GSessionMgr.g_SessionMgr)
	//x.PrintDbg(GSessionMgr.CheckPrivilige(1, "ldkfnl"))
	s := GSessionMgr.CreateSession()
	x.PrintDbg(s)
	x.PrintDbg(GSessionMgr.g_SessionMgrMap)

}
