// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"xMgr/sys"
	"xsw/go_pub/x"
	"xsw/go_pub/xWeb"
)

func InitWeb(port int) *x.Error {
	defer sys.DelRoutine()
	////////////////game_data/////////////////////////////////////
	http.HandleFunc("/", xWeb.DoHttpHandle)
	http.HandleFunc("/upload", uploadFileHandle)
	http.HandleFunc("/post", execFileHandle)
	http.HandleFunc("/project", LoginHandle)
	http.HandleFunc("/e_mail", EmailHandle)
	http.HandleFunc("/e_mail_skip", MailOneHandle)
	http.HandleFunc("/id_name", IdAndNameHandle)
	http.HandleFunc("/e_mail_group", MailGroupHandle)
	http.HandleFunc("/player_list", PlayerListHandle)
	http.HandleFunc("/player_role", PlayerRoleHandle) //PlayerResolveHandle
	http.HandleFunc("/player_channel", PlayerChannelHandle)
	http.HandleFunc("/player_close", PlayerCloseHandle)
	http.HandleFunc("/player_resolve", PlayerResolveHandle)
	////////////////recharge/////////////////////////////////////
	http.HandleFunc("/query_order", QueryOrderHandle)
	http.HandleFunc("/resupply_order", ResupplyOrderHandle)
	////////////////gs/////////////////////////////////////
	http.HandleFunc("/table", TableHandle) //TableHandle
	x.LogInfoF("listen port:%d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		return x.XErr(err)
	}
	return nil
}

func InitWebsocket() *x.Error {
	go h.run()
	http.HandleFunc("/ws", serveWs)
	pResult := new(Result)
	x.PrintInfo(pResult)
	//SetConfig("config/xMgr.xml", pResult)
	sys.AddRoutine()
	//go InitWeb(pResult.Port)
	go InitWeb(8888)
	return nil
}

//结构体与xml的布局相对应
type Result struct {
	Port int `xml:"Port"`
}

func SetConfig(pXmlUrl string, pStruct interface{}) {

	pContent, err := ioutil.ReadFile(pXmlUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(pContent, &pStruct)
	if err != nil {
		log.Fatal(err)
	}
}
