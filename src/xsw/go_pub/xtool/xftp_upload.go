package xtool

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dutchcoders/goftp"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"xsw/go_pub/x"
)

//ftp连接参数
type FtpParam struct {
	User string
	Psw  string
	Ip   string
	Port int

	Debug      bool //显示调试信息
	Crc        bool //文件上传后是否再下载一次，并判断crc
	Thread_num int
	Retry      int //失败后重试
}

const (
	myUPLOAD_MSG = 0
	myNULL_STR   = string("")
)

type myMsg struct {
	msg_type          int
	strLocalFilePath  *string
	strRemoteFilePath *string
	retry             int
}

//ftp上传参数
type FtpUploadParam struct {
	Local_path  string
	Remote_path string
}

type XFtpInfo struct {
	chanLog   chan myMsg
	wait_quit sync.WaitGroup
	ftp_param FtpParam
	quit      bool
}

/////////////////////////////////////////////////////////////////////////////////////
//外部函数
func InitXFtp(ftp_info FtpParam, pXFtpInfo *XFtpInfo) error {
	if len(ftp_info.User) <= 0 {
		return errors.New("need user")
	}
	if len(ftp_info.Psw) <= 0 {
		return errors.New("need Psw")
	}

	pXFtpInfo.chanLog = make(chan myMsg, 10000)
	pXFtpInfo.ftp_param = ftp_info
	pXFtpInfo.quit = false

	return nil
}

func FiniXFtp(pXFtpInfo *XFtpInfo) {
	pXFtpInfo.quit = true
	pXFtpInfo.wait_quit.Wait()
}

func XFtpParseRemoteStr(ftp_info *FtpParam, upload_param *FtpUploadParam, cmd string) error {
	// cmd := "admin:khgka*&^663@125.212.251.188:32021/d/ff.log --thread=8"
	pos := -1

	//user:
	pos = strings.IndexAny(cmd, ":")
	if pos <= 0 {
		return errors.New("don't find user:")
	}
	ftp_info.User = cmd[0:pos]
	cmd = cmd[pos+1:]

	//password@
	pos = strings.IndexAny(cmd, "@")
	if pos <= 0 {
		return errors.New("don't find password@")
	}
	ftp_info.Psw = cmd[0:pos]
	cmd = cmd[pos+1:]

	//ip:
	pos = strings.IndexAny(cmd, ":")
	if pos <= 0 {
		return errors.New("don't find ip:")
	}
	ftp_info.Ip = cmd[0:pos]
	cmd = cmd[pos+1:]

	//port/
	pos = strings.IndexAny(cmd, "/")
	if pos <= 0 {
		return errors.New("don't find port/")
	}
	ftp_info.Port, _ = strconv.Atoi(cmd[0:pos])
	cmd = cmd[pos+1:]

	//path
	upload_param.Remote_path = "/" + cmd
	x.AddPathEndFlag(&upload_param.Remote_path)
	upload_param.Remote_path += x.GetParentName(upload_param.Local_path)
	x.AddPathEndFlag(&upload_param.Remote_path)

	return nil
}

func DoUpload(upload_param FtpUploadParam, pXFtpInfo *XFtpInfo) error {
	var err error
	if pXFtpInfo == nil {
		return errors.New("must call InitXFtp first")
	}

	if !x.IsPath(upload_param.Local_path) {
		var ftp *goftp.FTP
		//open ftp
		if ftp, err = connectFtp(pXFtpInfo.ftp_param); err != nil {
			if ftp != nil {
				ftp.Close()
			}
			return err
		}
		defer ftp.Close()

		strLocalFile := x.GetFileName(upload_param.Local_path)
		strRemoteFile := x.GetFileName(upload_param.Remote_path)
		if strLocalFile != strRemoteFile {
			upload_param.Remote_path = x.GetDirName(upload_param.Remote_path) + strLocalFile
		}

		err = uploadFile(ftp, &upload_param.Local_path, &upload_param.Remote_path, true, pXFtpInfo.ftp_param.Crc)

	} else {
		for i := 0; i < pXFtpInfo.ftp_param.Thread_num; i++ {
			pXFtpInfo.wait_quit.Add(1)
			go ftpRoutine(pXFtpInfo, strconv.Itoa(i))
		}

		err = filepath.Walk(upload_param.Local_path, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if f == nil {
				return errors.New("nil os.FileInfo")
			}
			if f.IsDir() {
				return nil
			}

			str := path[len(upload_param.Local_path):]

			local_file := path
			remote_file := upload_param.Remote_path + str
			local_file = strings.Replace(local_file, "\\", "/", -1)
			remote_file = strings.Replace(remote_file, "\\", "/", -1)
			xftp_sync_msg(pXFtpInfo, myUPLOAD_MSG, &local_file, &remote_file, pXFtpInfo.ftp_param.Retry)
			return nil
		})
	}

	return err
}

/////////////////////////////////////////////////////////////////////////////////////
//内部函数

func connectFtp(ftp_info FtpParam) (*goftp.FTP, error) {
	var err error
	var ftp *goftp.FTP

	if ftp_info.Debug {
		if ftp, err = goftp.ConnectDbg(fmt.Sprintf("%s:%d", ftp_info.Ip, ftp_info.Port)); err != nil {
			return nil, err
		}
	} else {
		if ftp, err = goftp.Connect(fmt.Sprintf("%s:%d", ftp_info.Ip, ftp_info.Port)); err != nil {
			return nil, err
		}
	}

	if err = ftp.Login(ftp_info.User, ftp_info.Psw); err != nil {
		return ftp, err
	}

	return ftp, nil
}

func xftp_sync_msg(pXFtpInfo *XFtpInfo, _msg_type int, _strLocalFilePath *string, _strRemoteFilePath *string, _retry int) error {
	msg_data := myMsg{
		msg_type:          _msg_type,
		strLocalFilePath:  _strLocalFilePath,
		strRemoteFilePath: _strRemoteFilePath,
		retry:             _retry,
	}
	//写消息到channel
	pXFtpInfo.chanLog <- msg_data
	return nil
}

func ftpRoutine(pXFtpInfo *XFtpInfo, sn string) {
	var err error
	stop_flag := false
	defer pXFtpInfo.wait_quit.Done()
	x.LogDbgF("[" + sn + "] routine begin")

	var ftp *goftp.FTP
	//open ftp
	if ftp, err = connectFtp(pXFtpInfo.ftp_param); err != nil {
		if ftp != nil {
			ftp.Close()
		}
		x.LogDbgF("["+sn+"] connectFtp err:%v", err)
	}

	var msg myMsg
	for {
		if stop_flag {
			break
		}

		if ftp == nil {
			time.Sleep(3 * time.Second)
			if ftp, err = connectFtp(pXFtpInfo.ftp_param); err != nil {
				if ftp != nil {
					ftp.Close()
					ftp = nil
				}
				x.LogInfo("[" + sn + "] connectFtp fail:" + err.Error())
			}
		}
		if ftp == nil {
			continue
		}

		select {
		//有日志
		case msg = <-pXFtpInfo.chanLog:
			{
				x.PrintInfoF("[%s][%s:%d/%s] ==> remote[%s]", sn, pXFtpInfo.ftp_param.Ip, pXFtpInfo.ftp_param.Port, *msg.strLocalFilePath, *msg.strRemoteFilePath)
				err = uploadFile(ftp, msg.strLocalFilePath, msg.strRemoteFilePath, true, pXFtpInfo.ftp_param.Crc)
				if err != nil {
					x.LogInfoF("[%v] uploadFile[%s] remote[%s] fail[%v]", sn, *msg.strLocalFilePath, *msg.strRemoteFilePath, err.Error())
					retry := msg.retry - 1

					{
						xftp_sync_msg(pXFtpInfo, msg.msg_type, msg.strLocalFilePath, msg.strRemoteFilePath, retry)
						ftp.Close()
						ftp = nil
					}
				} else {
					x.LogDbgF("[%v] upload file[%s] suc", sn, *msg.strLocalFilePath)
				}
			}
			break
			//超时，避免过忙
		case <-time.After(100 * time.Millisecond):
			if pXFtpInfo.quit {
				stop_flag = true
				break
			}
			break
		}
	}
	x.LogDbgF("[" + sn + "] routine quit")
	ftp.Close()
}

//上传单个文件
func uploadFile(ftp *goftp.FTP, local_file *string, remote_file *string, create_dir bool, crc bool) error {
	var err error
	var oFile *os.File
	var strCur string

	//make remote dir
	if create_dir {
		remote_dirs := x.GetDirNames(*remote_file)
		n := len(*remote_dirs)
		var strRmote string
		for i := n - 1; i >= 0; i-- {
			strRmote = (*remote_dirs)[i]
			if (len(strRmote)) <= 0 {
				continue
			}
			ftp.Cwd(strRmote)
			strCur, err = ftp.Pwd()
			x.AddPathEndFlag(&strCur)
			x.AddPathEndFlag(&strRmote)
			if strCur != strRmote {
				if err = ftp.Mkd(strRmote); err != nil {

				}
			}
		}
	}

	//open local file
	if oFile, err = os.Open(*local_file); err != nil {
		return err
	}
	defer oFile.Close()
	local_hash := x.Md5FromIO(oFile)
	oFile.Seek(0, 0)

	//upload
	if err := ftp.Stor(*remote_file, oFile); err != nil {
		return err
	}

	if crc {
		x.LogDbgF("file[%v]check hash .....", *remote_file)
		var remote_hash []byte
		//Download and check
		_, err = ftp.Retr(*remote_file, func(r io.Reader) error {
			remote_hash = x.Md5FromIO(r)
			return nil
		})
		if bytes.Compare(remote_hash, local_hash) != 0 {
			return errors.New("hash not match")
		}
		x.PrintInfoF("[%s]check hash suc", *remote_file)
	}

	return nil
}
