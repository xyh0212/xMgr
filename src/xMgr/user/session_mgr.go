package user

import (
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

type SessionInfo struct {
	UserRec *db.SUser
	//xyh:*s_user
}
type SessionMgr struct {
	g_SessionMgrMap map[int]SessionInfo //key : cid
}

// 新增类SessionMgr，对应一个全局变量，所有对本文件的接口访问都以这个类为入口，里面管理一个map[]SessionInfo
// 处理登录cmd或者http请求时，调用SessionMgr.AddSession(cid, strSession),遍历s_user添加到g_SessionMgrMap
// 所有cmd和http请求，都调用func SessionMgr.CheckPrivilige(cid, nCmd int)

//var g_SessionMgrMap map[string]SessionInfo //key:session
var GSessionMgr SessionMgr

func init() {
	//SessionMgr.g_SessionMgrMap
	GSessionMgr.g_SessionMgrMap = make(map[int]SessionInfo)
}

//添加g_SessionMgrMap
func (self *SessionMgr) AddSession(cid int, strSession string) {
	for _, value := range db.TbUserMap {
		if strSession == value.Session {
			var s SessionInfo
			s.UserRec = value
			self.g_SessionMgrMap[cid] = s
		}
	}
}

//通过cid检验权限
func (self *SessionMgr) CheckPriviligeByCid(cid, nCmd int) bool {
	if nCmd == sys.CMD_LOGIN || nCmd == sys.CMD_CHECK_SESSION {
		return true
	}
	if nCmd == sys.CMD_QUERY_GS || nCmd == sys.CMD_QUERY_SQL {
		return true
	}
	sessionInfo, ok := self.g_SessionMgrMap[cid]
	if !ok {
		return false
	}
	return self.CheckPriviligeBySession(sessionInfo.UserRec.Session, nCmd)
}

//通过session检验权限
func (self *SessionMgr) CheckPriviligeBySession(strSession string, nCmd int) bool {
	if strSession == "" {
		return false
	}
	// nCmd<1001 is websocket; nCmd >1001 is http
	if nCmd < sys.CMD_HTTP_UPLOAD {
		for _, value := range self.g_SessionMgrMap {
			if value.UserRec.Session == strSession && value.UserRec.UserGroupId == sys.USER_GTOUP_ID {
				return true
			}
		}
		return false
	}
	return true
}

//cid id key
func (self *SessionMgr) DelSession(cid int) {
	delete(GSessionMgr.g_SessionMgrMap, cid)

}

//
func (self *SessionMgr) CreateSession() string {
	return x.RandStr(32)
}
