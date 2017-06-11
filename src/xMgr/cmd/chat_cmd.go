package cmd

//file:///C:/Users/Administrator/Desktop/wfile:///C:/Users/Administrator/Desktop/workplace/xMgr/src/xMgr/cmd/querygs_cmd.goorkplace/xMgr/src/xMgr/cmd/mgr_cmd.go

import (
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":4,"param":{"chat":"聊天内容"}}
func DoChatCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintDbg("chat")
	chat, _ := cmd.jsParam.CheckGet("chat")
	if chat == nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, nil
	}

	strChat, _ := chat.String()

	return ([]byte)(strChat), nil
}
