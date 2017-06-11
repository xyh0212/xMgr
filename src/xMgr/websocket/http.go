package websocket

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"xMgr/db"
	"xMgr/sys"
	"xMgr/user"
	"xsw/go_pub/x"

	"github.com/astaxie/beego/orm"
)

/////////////////////////游戏数据_game_data//////////////////////////////////////

//玩家解封
func PlayerResolveHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("PlayerResolveHandle")
	r.ParseForm()

	gsStr := r.PostFormValue("gs")
	gs, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
		return
	}
	gsInfo, ok := db.G_GSInfoMap[gs]
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_S_GS_ID_NO_EXIST))
		return
	}
	role := r.PostFormValue("role")
	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user where name='%v'`
	sql = fmt.Sprintf(sql, role)
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}

	sql = `insert into cmd_d (type,user_id,state,start_time,txt)values('7','%v','1','0','');`
	resultMap := result.ToMap()
	if resultMap == nil {
		w.Write(CreateErrMsg(sys.CODE_ROLE_NOT_EXIST))
		return
	}
	//sql = fmt.Sprintf(sql, resultMap["id"][0])
	err1 = gsInfo.ExecSql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	s := result.ToTableJson()
	w.Write(CreateSucceedMsg(*s))
}

//玩家封号
func PlayerCloseHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("PlayerCloseHandle")
	r.ParseForm()

	gsStr := r.PostFormValue("gs")
	gs, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
		return
	}
	gsInfo, ok := db.G_GSInfoMap[gs]
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_S_GS_ID_NO_EXIST))
		return
	}
	role := r.PostFormValue("role")
	close_time := r.PostFormValue("close_time")
	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user where name='%v'`
	sql = fmt.Sprintf(sql, role)
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}

	sql = `insert into cmd_d (type,user_id,state,start_time,txt,eff_seconds)values('6','%v','1','0','','%v');`
	resultMap := result.ToMap()
	if resultMap == nil {
		w.Write(CreateErrMsg(sys.CODE_ROLE_NOT_EXIST))
		return
	}
	sql = close_time
	//	sql = fmt.Sprintf(sql, resultMap["id"][0], close_time)
	err1 = gsInfo.ExecSql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	s := result.ToTableJson()
	w.Write(CreateSucceedMsg(*s))
}

//渠道账号
//select * from d_user where channel_accountname='渠道账号';
func PlayerChannelHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("PlayerListHandle")
	r.ParseForm()

	gsStr := r.PostFormValue("gs")
	gs, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
		return
	}
	gsInfo, ok := db.G_GSInfoMap[gs]
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_S_GS_ID_NO_EXIST))
		return
	}
	channel := r.PostFormValue("channel")
	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user where channel_accountname='%v'`
	sql = fmt.Sprintf(sql, channel)
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	w.Write(CreateSucceedMsg(*result.ToTableJson()))
}

//玩家角色
//select * from d_user where name='角色名';
func PlayerRoleHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("PlayerListHandle")
	r.ParseForm()

	gsStr := r.PostFormValue("gs")
	gs, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
		return
	}
	gsInfo, ok := db.G_GSInfoMap[gs]
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_S_GS_ID_NO_EXIST))
		return
	}
	role := r.PostFormValue("role")
	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user where name='%v'`
	sql = fmt.Sprintf(sql, role)
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	w.Write(CreateSucceedMsg(*result.ToTableJson()))
}

//玩家列表
//select * from d_user where channel=渠道类型;
//arrChan=`1,4,4,3,7`
func PlayerListHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("PlayerListHandle")
	r.ParseForm()

	gsStr := r.PostFormValue("gs")
	gs, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
		return
	}

	pageStr := r.PostFormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
		return
	}

	chanStr := r.PostFormValue("channel")
	arrChan := strings.Split(chanStr, ",")
	whereStr := " where "
	for _, v := range arrChan {
		whereStr += fmt.Sprintf(` channel='%v' or`, v)
	}
	whereStr = strings.TrimRight(whereStr, "or")

	limitStr := ` limit %v,50`
	limitStr = fmt.Sprintf(limitStr, page*50)

	sql := `select id,name,level,channel,emoney_total,login_time,create_time from d_user ` + whereStr + limitStr
	gsInfo, ok := db.G_GSInfoMap[gs]
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_S_GS_ID_NO_EXIST))
		return
	}
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	sql = `SELECT count(0) from d_user ` + whereStr
	resultPage, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	//x.PrintDbg(*result, *resultPage)
	w.Write(CreateSucceedMultiMsg(*result.ToTableJson(), *resultPage.ToJson()))
}

//// insert into d_package_item(`user_id`,`package`,`title_txt`,`desc_txt`,`itemtype_main`,`itemtype_param`,`num`,`time_t`)
// values ("UserID","1","标题文字","内容文字","物品主类型","物品参数","个数",UNIX_TIMESTAMP())
//发送单个的邮件
func MailOneHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("EmailSkipHandle")
	r.ParseForm()
	gs := r.PostFormValue("gs")
	id, err := strconv.Atoi(gs)
	if err != nil {
		x.LogErr(x.XErr(err))
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}

	user_id := r.PostFormValue("user_id")
	title_txt := r.PostFormValue("title_txt")
	desc_txt := r.PostFormValue("desc_txt")
	item_type_main := r.PostFormValue("item_type_main")
	item_type_param := r.PostFormValue("item_type_param")
	num := r.PostFormValue("num")

	sql := `insert into d_package_item(user_id,package,title_txt,desc_txt,itemtype_main,itemtype_param,num,time_t)`
	sql += `values('%v','%v','%v','%v','%v','%v','%v',%v)`

	sql = fmt.Sprintf(sql, user_id, "1", title_txt, desc_txt, item_type_main, item_type_param, num, "UNIX_TIMESTAMP()")

	gsInfo, ok := db.G_GSInfoMap[id]
	if !ok {
		x.LogErr(`db.G_GSInfoMap[id]的key不存在`)
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}
	err1 := gsInfo.ExecSql(sql) //  Sprintf(item_type_param)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	w.Write(CreateSucceedMsg("suc"))
}

//接收多个游服 gs=`1,5,5,5`
func MailGroupHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("EmailGroupHandle")
	r.ParseForm()
	gsStr := r.PostFormValue("gs")
	//x.PrintInfo("rec:", gsStr)
	arrGsID, err := sys.String2ArryInt(gsStr, ",")
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "错误"))
	}

	//user_id := r.PostFormValue("user_id")
	title_txt := r.PostFormValue("title_txt")
	desc_txt := r.PostFormValue("desc_txt")
	item_type_main := r.PostFormValue("item_type_main")
	item_type_param := r.PostFormValue("item_type_param")
	num := r.PostFormValue("num")

	sql := `insert into d_package_item(user_id,package,title_id,desc_id,  itemtype_main,itemtype_param,num,time_t,title_txt,desc_txt)`
	sql += ` SELECT 11,1,0,0,'%v','%v','%v',%v,'%v','%v'`
	sql += ` from d_user where login_time+30*24*60*60>UNIX_TIMESTAMP()`
	sql = fmt.Sprintf(sql, item_type_main, item_type_param, num, "UNIX_TIMESTAMP()", title_txt, desc_txt)
	for _, v := range arrGsID {
		gsInfo, ok := db.G_GSInfoMap[v]
		if !ok {
			x.LogErr(`db.G_GSInfoMap[v]的key不存在`)
			w.Write(CreateErrMsg(sys.CODE_ERR))
			return
		}
		err1 := gsInfo.ExecSql(sql) //  Sprintf(item_type_param)
		if err1 != nil {
			x.LogErr(err1)
			w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
			return
		}
	}
	w.Write(CreateSucceedMsg("suc"))
}

func EmailHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("EmailHandle")
	r.ParseForm()
	gsStr := r.PostFormValue("gs_id")
	id, _ := strconv.Atoi(gsStr)
	gsInfo, ok := db.G_GSInfoMap[id]
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_GS_NOT_EXIST))
		return
	}
	err := gsInfo.InitPDB()
	if err != nil {
		w.Write(CreateErrInfoMsg(sys.CODE_DB_GROUP_ID_NO_EXIST, err.GetStr()))
		return
	}

	sql := `select id,name,level from d_user where name='` + r.PostFormValue("name") + `'`
	rows, err := x.CreateRecordSet(gsInfo.PDB, sql)
	if err != nil {
		x.PrintDbg(err)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err.GetStr()))
		return
	}
	s := rows.ToJson()
	if s == nil {
		w.Write(CreateErrMsg(sys.CODE_NAME_NOT_EXIST))
		return
	}
	x.PrintDbg(*s)
	w.Write(CreateSucceedMsg(*s))
}

/////////////////////////其他////////////////////////////////////////////////////////////////////////
//登入网址
func LoginHandle(w http.ResponseWriter, r *http.Request) {
	var uploadTemplate = template.Must(template.ParseFiles("www/project/login.html"))
	if err := uploadTemplate.Execute(w, nil); err != nil {
		x.LogErrF("Execute: ", err.Error())
		return
	}
}

//上传文件
///post url=upload/2016-02/xyh.txt&id=skdf&name=xyh
func uploadFileHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("uploadFileHandle start")
	pmap, err := UrlByMap(r.URL.String())
	if err != nil {
		w.Write(CreateErrMsg(sys.CODE_ERR))
		x.PrintDbg(err)
		return
	}
	ok := user.GSessionMgr.CheckPriviligeBySession(pmap["session"], sys.CMD_HTTP_UPLOAD)
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_NO_PRIVILIGE))
		return
	}
	name, err := db.SessionByUserName(pmap["session"])
	if err != nil {
		w.Write(CreateErrMsg(sys.CODE_SESSION_NO_EXIST))
		return
	}
	strs := StringByMap(pmap["filekey"])
	var str string
	for _, v := range strs {
		file, fileHead, err := r.FormFile(v) //文件file  //文件名fileHead.name
		if err != nil {
			w.Write([]byte(CreateErrMsg(sys.CODE_ERR)))
			x.LogErrF("FormFile: ", err.Error())
			return
		}
		//x.PrintDbg("文件名：", fileHead.Filename)
		s, _ := CreatePath("upload", name, fileHead.Filename)
		_, err1 := CopyFile(*s, file)
		if err1 != nil {
			x.PrintInfo("copy fail", err1)
		}
		str += *s + ":"
	}
	str = strings.TrimRight(str, ":")
	//x.PrintDbg(str)
	w.Write([]byte(CreateSucceedMsg(str)))
}

//执行文件
///post url="upload/2016-02/xyh.txt&id=skdf&name=xyh
func execFileHandle(w http.ResponseWriter, r *http.Request) {
	pmap, err := UrlByMap(r.URL.String())
	if err != nil {
		w.Write([]byte(CreateErrMsg(sys.CODE_ERR)))
		x.LogInfo(err)
		return
	}
	ok := user.GSessionMgr.CheckPriviligeBySession(pmap["session"], sys.CMD_HTTP_EXEC_FILE)
	if !ok {
		w.Write(CreateErrMsg(sys.CODE_NO_PRIVILIGE))
		return
	}

	name, err := db.SessionByUserName(pmap["session"])
	if err != nil {
		w.Write(CreateErrMsg(sys.CODE_SESSION_NO_EXIST))
		return
	}
	arrInt, err := sys.String2ArryInt(pmap["info"], ":")

	arrPath := sys.String2ArrStr(pmap["path"], ":")
	for i, path := range arrPath {
		x.PrintErr(i)
		x.PrintErr(path)
		//arrPath[i] = x.UrlDecode(path)
	}
	var clientStr string
	var serveStr string
	for _, value := range arrInt {
		var str string
		var strTmpSvr string
		for i, path := range arrPath {
			nameofPath := sys.GetLeftStr(sys.GetRightStr(path, `/`), "-")
			//serveStr += path
			if name != nameofPath {

				w.Write(CreateErrMsg(sys.CODE_NAME_NOT_EXIST))
				return
			}
			b, err := x.ExcuSqlFile(db.CreateDBXmlNode(value), path) // "upload/2016-02/xyh.txt"
			if err != nil {
				x.LogDbg(err)
			}

			strTmpSvr += `File` + strconv.Itoa(i+1) + `:` + path + "\n" + b
			b = strings.Replace(b, "\r\n", "<br>", -1)
			b = strings.Replace(b, "\r", "<br>", -1)
			b = strings.Replace(b, "\n", "<br>", -1)
			path = sys.GetRightStr(sys.GetRightStr(path, `/`), "-")

			str += `File` + strconv.Itoa(i+1) + `:` + path + `</br>` + b + `</br><hr></hr>`
		}
		sGs, ok := db.TbGseMap[value]
		if !ok {
			x.LogErr(`db.TbGseMap[value]的key不存在`)
			w.Write(CreateErrMsg(sys.CODE_ERR))
		}
		serveStr += sGs.GetDBGS2String("") + strTmpSvr
		clientStr += sGs.GetDBGS2String(`</br>`) + str + `</br>`
	}

	err2 := ioutil.WriteFile(CreateSavePath(arrPath[0]), []byte(serveStr), 0666)
	if err2 != nil {
		x.LogErr(err2)
	}
	w.Write([]byte(clientStr))
}

//获取表的两个字段 例如：id与姓名
func IdAndNameHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("IdAndNameHandle")
	r.ParseForm()
	gsStr := r.PostFormValue("table_id") //String2ArryInt
	pdb, err1 := orm.GetDB()
	if err1 != nil {
		x.LogErr(x.XErr(err1))
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}
	x.SetDBCharacter(pdb)
	id, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(x.XErr(err))
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}
	dropList := db.DropListMap[id]
	str, err2 := db.GetIDAndName(dropList.TableName, pdb, dropList.Id, dropList.Name)
	if err2 != nil {
		x.LogErr(err2)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err2.GetStr()))
		return
	}

	w.Write(CreateSucceedMsg(*str))
}
