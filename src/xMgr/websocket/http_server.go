package websocket

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
	"time"
	"xMgr/sys"
	"xsw/go_pub/x"
)

//http://localhost:8888/project/user.html?id=1323&name=xyh&session=jsdfk&path=sdf:kufd:dfsd:jdxf
//当session为空时返回错误
//当session验证失败时，返回user已过期
//返回值 pmap, username, err
//pmap["id"]为123；pmap["name"]为xyh
func UrlByMap(url string) (map[string]string, *x.Error) {
	arrStr := strings.Split(url, "?")
	if len(arrStr) == 1 {
		return nil, x.XErrStr(string(CreateErrMsg(sys.CODE_SESSION_NO_EXIST)))
	}
	pmap, _ := sys.String2Map(arrStr[1])
	return pmap, nil
}

func CreatePath(pathHead, userName, fileName string) (*string, *x.Error) {
	path := GetPath(pathHead)
	//err := x.MkdirAll(path)
	//	if err != nil {
	//		x.LogInfo(err)
	//		return nil, err
	//	}
	path = path + "/" + GetfileName(userName, fileName)
	return &path, nil
}
func StringByMap(str string) []string {
	str = strings.Replace(str, "undefined:", "", -1)
	arrstr := sys.String2ArrStr(str, ":")
	return arrstr
}
func GetfileName(userName, fileName string) string {
	numstr := ""
	t := time.Now().Unix()
	str := time.Unix(t, 0).String()
	arrstr := strings.Split(str, "+")
	arrstr = strings.Split(arrstr[0], " ")
	numstr += userName + "-" + arrstr[0][5:10] + "_"
	ss := strings.Replace(str[11:19], ":", "-", -1)
	numstr += ss + "-" + fileName
	return numstr
}
func GetPath(s string) string {
	//x.GetDirName()
	t := time.Now().Unix()
	str := time.Unix(t, 0).String()
	return s + "/" + str[0:7]
}

func CopyFile(src string, file multipart.File) (w int64, err error) {
	dstFile, err := os.Create(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, file)
}
func CreateSavePath(path string) string {
	path = sys.GetRightStr(path, "upload/")
	str := sys.GetRightStr(path, ".")
	path = sys.GetLeftStr(path, "."+str)
	path = `log/sqlresult/` + path
	str = sys.GetRightStr(path, "/")
	//	err := x.MkdirAll(sys.GetLeftStr(path, "/"+str))
	//	if err != nil {
	//		x.LogInfo(err)
	//		return ""
	//	}
	path = path + ".log"
	return path
}
