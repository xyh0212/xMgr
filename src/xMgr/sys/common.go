package sys

import (
	//"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"xsw/go_pub/x"
)

type IDBAryToString interface {
	ToString() string
}

//排序int[]
func SortArryInt(arr []int) []int {
	var temp int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

//排序int64[]
func SortArryInt64(arr []int) []int {
	var temp int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

//返回值为字符串["",""]
func ObjToString(pObj interface{}) string {
	var sumstr string
	sumstr = `[`
	object := reflect.ValueOf(pObj)
	myref := object.Elem()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		str := fmt.Sprintf("%v", field.Interface())
		sumstr += `"` + str + `",`
	}
	sumstr = strings.TrimRight(sumstr, ",")
	sumstr += `]`
	return sumstr
}

//结构体 返回值为[]string
//type User struct {
//Name string
//Age  int
//}
//转为 []string{"Name","age"}
func StructToArrStr(pObj interface{}) []string {
	object := reflect.ValueOf(pObj)
	myref := object.Elem()
	typeOfType := myref.Type()
	var arr []string
	arr = make([]string, myref.NumField())
	for i := 0; i < myref.NumField(); i++ {
		str := fmt.Sprintf("%v", typeOfType.Field(i).Name)
		arr[i] = x.SnakeString(str)

		//x.PrintDbg(x.SnakeString(str))
	}
	return arr
}

//表字段的字符串 ["","",""]
func StructToString(pObj interface{}) string {
	object := reflect.ValueOf(pObj)
	myref := object.Elem()
	typeOfType := myref.Type()
	//var str string
	str := `[`
	for i := 0; i < myref.NumField(); i++ {
		s := fmt.Sprintf("%v", typeOfType.Field(i).Name)
		str += `"` + x.SnakeString(s) + `",`
	}
	str = strings.TrimRight(str, ",") + `]`
	return str
}

//////////////////////////////////////////////////////////
func DBAryToString(pAry IDBAryToString) string {
	var sumstr string
	sumstr = `{"userjson":[`
	sumstr += pAry.ToString()
	sumstr = strings.TrimRight(sumstr, ",")
	sumstr += `]}`
	fmt.Println(sumstr)
	return sumstr
}

//"id:56,name:dfg,name2:dlf,gs_group_id:14"
func String2Map(str string) (map[string]string, *x.Error) {
	arrStr := strings.Split(str, "&")
	return ArrayString2Map(arrStr, "=")
}
func ArrString2Map(arrStr []string) (map[string]string, *x.Error) {
	return ArrayString2Map(arrStr, ":")
}

//str的格式：id:56name:dfgname2:dlid:14
func String2ArryInt(str, sign string) ([]int, *x.Error) {
	var arrInt []int
	arrStr := strings.Split(str, sign)
	arrInt = make([]int, len(arrStr))
	for i, value := range arrStr {
		j, err := strconv.Atoi(value)
		if err != nil {
			//x.PrintInfo(err)
			return nil, x.XErr(err)
		}
		arrInt[i] = j
	}
	return arrInt, nil
}

//str = aksdgjlskd:dfkl:djfgda:ldg:kjg
// GetRightStr(str,":")  return kjg
//str = dfgdsi
//GetRightStr(str,"ds")  return i
func GetRightStr(str, sign string) string {
	arr := String2ArrStr(str, sign)
	return arr[len(arr)-1]

}

//str = 123:dfkl:djfgda:ldg:kjg
// GetRightStr(str,":")  return 123
func GetLeftStr(str, sign string) string {
	arr := String2ArrStr(str, sign)
	return arr[0]

}

//str的格式：id:56name:dfgname2:dlid:14
func String2ArrStr(str, sign string) []string {
	return strings.Split(str, sign)
}

//把字符串数组转换成map
//例如：数组{"id:56","name:dfg","name2:dlf","gs_group_id:14"}
//转成：map[id]="56" ;map[name]="dfg"
func ArrayString2Map(arrStr []string, sign string) (map[string]string, *x.Error) {
	pMap := make(map[string]string)
	for _, value := range arrStr {
		Strs := strings.Split(value, sign)
		if len(Strs) == 1 {
			return nil, x.XErrStr(value)
		}
		str := ""
		for i := 1; i < len(Strs); i++ {
			str += Strs[i] + sign
		}
		str = strings.TrimRight(str, sign)
		pMap[Strs[0]] = str
	}
	return pMap, nil
}

// Map2Struct(m, &TB_s_gs)
func Map2Struct(m map[string]string, obj interface{}) {
	object := reflect.ValueOf(obj)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		value := fmt.Sprintf("%v", typeOfType.Field(i).Name) //typeOfType.Field(i).Name
		value = x.SnakeString(value)
		str := m[value]
		//fmt.Println(str)
		str1 := fmt.Sprintf("%v", myref.Field(i).Type())
		switch str1 {
		case "int":
			s, _ := strconv.Atoi(str)
			myref.Field(i).SetInt(int64(s))
		case "string":
			myref.Field(i).SetString(str)
		case "int64":
			s, _ := strconv.Atoi(str)
			myref.Field(i).SetInt(int64(s))
		case "time.Time":
			s, _ := time.Parse("2006-01-02 15:04:05", str)
			myref.Field(i).Set(reflect.ValueOf(s))
		default:
			x.LogErr("没有指定的类型")
		}
	}

}

//html的空格
func GetBlankString(nun int) string {
	str := "&nbsp;"
	for i := 1; i < nun; i++ {
		str += str
	}
	return str
}
