package x

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StringResInfo struct {
	strContent string
}

//key:file path
var g_StringResMgr map[string]*StringResInfo

///////////////////////////////////////////////////////////////////////////
//对外接口

//遍历res_dir，加载所有文件
func LoadStringRes(res_dir string) error {
	g_StringResMgr = make(map[string]*StringResInfo, 3)

	err := filepath.Walk(res_dir, walkFunc)
	if err != nil {
		return nil
	}

	return nil
}

func GetStringRes(path string) (string, error) {
	objRes := g_StringResMgr[path]
	if objRes == nil {
		loadResFile(path)
		objRes = g_StringResMgr[path]
	}

	if objRes == nil {
		return "", errors.New(fmt.Sprintf("%s is not exist", path))
	}
	return objRes.strContent, nil
}

///////////////////////////////////////////////////////////////////////////
//内部函数

func loadResFile(res_file_path string) error {
	content, err := ReadFile(res_file_path)
	if err != nil {
		return err
	}
	g_StringResMgr[res_file_path] = &StringResInfo{strContent: string(content)}
	return nil
}

func walkFunc(path string, f os.FileInfo, err error) error {
	if f == nil {
		return errors.New("nil os.FileInfo")
	}
	if f.IsDir() {
		return nil
	}

	if err != nil {
		return err
	}

	err = loadResFile(path)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	return fd, nil
}
