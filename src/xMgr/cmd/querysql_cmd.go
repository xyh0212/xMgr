package cmd

import (
	//"database/sql"
	_ "github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	//	"net/url"
	"strconv"
	"strings"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

// {"cmd":3,"param":{"arry_s_gs_id" :["1","1","1"] ,"sql":"xxxxx"}}
func DoQuerySqlCmd(cmd CmdInfo, cid int) ([]byte, *x.Error) {
	x.PrintInfo("DoQuerySqlCmd")
	sql, _ := cmd.jsParam.CheckGet("sql")
	if sql == nil {
		msg, _ := CreateErrMsg(sys.CODE_SQL_NO_EXIST)
		return msg, x.XErrStr("sql not exist")
	}
	arry_s_gs_id, _ := cmd.jsParam.CheckGet("arry_s_gs_id")
	if arry_s_gs_id == nil {
		msg, _ := CreateErrMsg(sys.CODE_SQL_NO_EXIST)
		return msg, x.XErrStr("arry_s_gs_id not exist")
	}
	psql, err := sql.String()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	//x.PrintInfo(psql)
	parry_s_gs_id, err := arry_s_gs_id.StringArray()
	if err != nil {
		msg, _ := CreateErrMsg(sys.CODE_ERR)
		return msg, x.XErrStr("erron")
	}
	//查询数据库
	arrystr, _ := ExcelSqlByGsIds(psql, parry_s_gs_id)
	msg, _ := CreateClientSqlCmdMsg(arrystr)
	return msg, x.XErrStr("success")
}

//查询客服端传来的多条sql，对应多个数据库，并返回查找的数据
// select * from cmd_d1 , select * from cmd_d2 ...
//db1，db2.....
func ExcelSqlByGsIds(pSql string, arry_s_gs_id []string) ([]string, *x.Error) {
	intArry := arrStrToarrInt(arry_s_gs_id)
	arrysql := parserSql(pSql)
	var strs []string
	strs = make([]string, len(arry_s_gs_id))
	for i, inta := range intArry {
		str, err := ExcelSqlByGsId(arrysql, inta)
		strs[i] = str
		if err != nil {
			return strs, err
		}
	}
	return strs, nil
}

//查找多条sql，对应的1个数据库，
// select * from cmd_d1 ,select * from cmd_d2  对应数据库db1
//返回 一个数据库所对应的多sql数据
func ExcelSqlByGsId(arrSql []string, s_gs_id int) (string, *x.Error) {
	var numStr string
	sGs, ok := db.TbGseMap[s_gs_id]
	if !ok {
		return "", x.XErrStr(`db.TbGseMap[s_gs_id]的key不存在`)
	}
	numStr = `area：` + sGs.Name + `</br>`
	for _, sql := range arrSql {
		for _, value := range db.G_GSInfoMap {
			if s_gs_id == value.S_GS_ID {
				s, err := value.QuerySqlTable(sql)
				if err != nil {
					numStr += sql + `</br>` + err.GetStr()
					return numStr, err
				} else {
					numStr += *s
				}

			}
		}
	}
	return numStr, nil
}

//cmd=3
func CreateClientSqlCmdMsg(arrystr []string) ([]byte, *x.Error) {
	msgMap := make(map[string]interface{})
	msgMap["cmd"] = sys.CMD_QUERY_SQL
	msgMap["sql"] = arrystr
	return CreateAnyMsg(msgMap)
}

//string[]转换成int[]
func arrStrToarrInt(arry_s_gs_id []string) []int {
	var intArry []int
	intArry = make([]int, len(arry_s_gs_id))
	for s, str := range arry_s_gs_id {
		b, _ := strconv.Atoi(str)
		intArry[s] = b
	}
	return intArry
}

//格式化字符串为多条sql
//如：str = select * from cmd_d1
//    select * from cmd_d2
//转为字符串数组 {"select * from cmd_d1","select * from cmd_d2"}
func parserSql(sql string) []string {
	//str, err := url.QueryUnescape(sql) //urldecode
	//url.QueryEscape()
	//x.PrintDbg(err)
	//str = strings.TrimSpace(str) //去除前后的转义字符
	str := sql
	//str := x.UrlDecode(sql)
	x.PrintInfo(str)
	for {
		if strings.Contains(str, "\n\n") || strings.Contains(str, "  ") {
			str = strings.Replace(str, "  ", " ", -1)
			str = strings.Replace(str, "\n\n", "\n", -1)

		} else {
			break
		}
	}
	strs := strings.Split(str, "\n")
	return strs
}
