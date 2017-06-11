package xtool

import (
	"encoding/xml"
	"io/ioutil"
	"os/exec"
	"time"
	"xsw/go_pub/x"
)

type guardXmlNode struct {
	Exe      string `xml:"exe"`
	Param    string `xml:"param"`
	WorkPath string `xml:"work_path"`
}

var (
	g_bInit        bool
	g_GuardXmlNode guardXmlNode
)

func initCfg() *x.Error {
	if g_bInit {
		return nil
	}
	var err error

	var strXml []byte
	if strXml, err = ioutil.ReadFile("guard.xml"); err != nil {
		return x.XErr(err)
	}
	if err = xml.Unmarshal(strXml, &g_GuardXmlNode); err != nil {
		return x.XErr(err)
	}

	return nil
}

func GuardExe() *x.Error {
	var xerr *x.Error
	xerr = initCfg()
	if xerr != nil {
		return xerr
	}

	for {
		cmd := exec.Command(g_GuardXmlNode.Exe, g_GuardXmlNode.Param)
		cmd.Dir = g_GuardXmlNode.WorkPath
		err := cmd.Start()
		if err != nil {
			x.PrintErr("启动命令失败 ", err)
			time.Sleep(time.Second * 30)
			continue
		}

		x.PrintInfo("进程启动")
		err = cmd.Wait()
		x.PrintErr("进程退出 ", err)

		time.Sleep(time.Second * 1)
	}
}
