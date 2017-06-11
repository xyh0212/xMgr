package x

import (
	"io/ioutil"
	"os"
	"strings"
)

//获取文件名
func GetFileName(file_name string) string {
	var name string
	file_len := len(file_name)
	for i := file_len - 1; i > 0; i-- {
		c := file_name[i]
		if c == '\\' || c == '/' {
			name = file_name[i+1:]
			break
		}
	}
	return name
}

//获取目录名
func GetDirName(file_name string) string {
	var name string
	file_len := len(file_name)
	for i := file_len - 1; i > 0; i-- {
		c := file_name[i]
		if c == '\\' || c == '/' {
			name = file_name[0 : i+1]
			name = strings.TrimSpace(name)
			break
		}
	}
	name = strings.Replace(name, "\\", "/", -1)
	return name
}

//获取各个目录名
//c:/d1/d2/f1.ext 输出 [   c:/d1/d2/ c:/d1/ c:/]
func GetDirNames(file_name string) *[]string {
	var names = make([]string, 2)
	file_len := len(file_name)
	for i := file_len - 1; i > 0; i-- {
		c := file_name[i]
		if c == '\\' || c == '/' {
			// fmt.Println(i)
			str := file_name[0 : i+1]
			str = strings.TrimSpace(str)
			if len(str) <= 0 {
				continue
			}
			str = strings.Replace(str, "\\", "/", -1)
			names = append(names, str)
		}
	}
	return &names
}

//获取父目录及文件名
func GetParentFileName(file_name string) string {
	var short_name string
	file_len := len(file_name)
	j := 0
	for i := file_len - 1; i > 0; i-- {
		c := file_name[i]
		if c == '\\' || c == '/' {
			j += 1
			if j > 1 {
				short_name = file_name[i+1:]
				break
			}
		}
	}
	return short_name
}

//获取父目录名
func GetParentName(file_name string) string {
	if len(file_name) <= 0 {
		return ""
	}
	file_name = strings.Replace(file_name, "\\", "/", -1)
	end := strings.LastIndex(file_name, "/")
	start := strings.LastIndex(file_name[0:end], "/")
	return file_name[start+1 : end]
}

// 目录后加/或者\\
func AddPathEndFlag(path *string) {
	l := len(*path) - 1
	if l > 0 {
		c := (*path)[l]
		if c != '\\' && c != '/' {
			*path += "/"
		}
	} else {
		*path += "/"
	}
	*path = strings.Replace(*path, "\\", "/", -1)
}

func IsPath(path string) bool {
	file_len := len(path)
	if file_len <= 0 {
		return false
	}
	c := path[file_len-1]
	if c == '\\' || c == '/' {
		return true
	}
	return false
}

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func ReplaceFileString(strFilePath, strOld string, strNew string) (xerr *Error) {
	buf, err := ioutil.ReadFile(strFilePath)
	if err != nil {
		return XErr(err)
	}
	content := string(buf)

	//替换
	newContent := strings.Replace(content, strOld, strNew, -1)

	//重新写入
	err = ioutil.WriteFile(strFilePath, []byte(newContent), 0)
	if err != nil {
		return XErr(err)
	}

	return nil
}
