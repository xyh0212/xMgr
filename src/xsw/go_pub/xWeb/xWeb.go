// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xWeb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"xsw/go_pub/x"
)

const (
	HOME_DIR = "www"
)

var (
	g_funcsCmd map[string]interface{}
)

func init() {
	g_funcsCmd = make(map[string]interface{})
}

func myHttpErr(w http.ResponseWriter, code int) { http.Error(w, http.StatusText(code), code) }

// func getFileContent(path string) []byte {
// 	path = HOME_DIR + "/" + path
// 	fi, err := os.Open(path)
// 	if err != nil {
// 		return nil
// 	}
// 	defer fi.Close()
// 	fd, err := ioutil.ReadAll(fi)
// 	return fd
// }

func myServeFile(w http.ResponseWriter, r *http.Request, strPath string) int {
	//h5 需要跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

	file_path := HOME_DIR + strPath
	fileinfo, err := os.Stat(file_path)
	if err != nil {
		if os.IsNotExist(err) {
			return http.StatusNotFound
		}

		if !os.IsPermission(err) {
			return http.StatusForbidden
		}
		return http.StatusForbidden
	}

	if fileinfo.IsDir() {
		return http.StatusNotImplemented
	}

	http.ServeFile(w, r, HOME_DIR+"/"+strPath)
	return http.StatusOK
}

func DefaultServeHome(w http.ResponseWriter, r *http.Request) {
	//h5 需要跨域访问
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

	defer r.Body.Close()
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		x.LogErr(x.XErr(err))
	}

	var strBody string
	if len(bBody) > 0 {
		strBody = string(bBody)
	}

	nStatusCode := myServeFile(w, r, r.URL.Path)

	if nStatusCode != http.StatusOK {
		myHttpErr(w, nStatusCode)
	}
	if len(strBody) > 0 {
		LogAccess(r, &strBody, nStatusCode)
	} else {
		LogAccess(r, nil, nStatusCode)
	}
}

func LogAccess(r *http.Request, strBody *string, nStatusCode int) {
	if r.Method == "GET" || strBody == nil {
		str := fmt.Sprintf("%s,%s,%s,%d", r.RemoteAddr, r.Method, r.RequestURI, nStatusCode)
		x.LogAccess(str)
	} else {

		str := fmt.Sprintf("%s,%s,%s?%s,%d", r.RemoteAddr, r.Method, r.RequestURI, *strBody, nStatusCode)
		x.LogAccess(str)
	}
}

func AddHttpHandle(strUrl string, f interface{}) {
	g_funcsCmd[strUrl] = f
}

// /根目录由此函数进
func DoHttpHandle(w http.ResponseWriter, r *http.Request) {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		x.LogErr(x.XErr(err))
	}
	r.Body.Close()
	r.ParseForm()

	var strBody string
	if len(bBody) > 0 {
		strBody = string(bBody)
	}
	strUrl := r.URL.Path
	oFunc := g_funcsCmd[strUrl]

	nStatusCode := http.StatusOK

	if oFunc == nil {
		nStatusCode = myServeFile(w, r, strUrl)
	} else {
		//h5 需要跨域访问
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型

		_, err = callReflect(oFunc, strUrl, w, r, &strBody)
		if err != nil {
			xerr := x.XErr(err)
			x.LogErr(xerr)
		}
	}

	if nStatusCode != http.StatusOK {
		myHttpErr(w, nStatusCode)
	}
	if len(strBody) > 0 {
		LogAccess(r, &strBody, nStatusCode)
	} else {
		LogAccess(r, nil, nStatusCode)
	}
}

// func handleUrl(w http.ResponseWriter, r *http.Request) {}

func callReflect(f2 interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(f2)
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return nil, err
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return result, err
}
