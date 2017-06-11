package websocket

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

/////////////////////////充值管理//////////////////////////////////////
//补单
func ResupplyOrderHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("ResupplyOrderHandle")
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
		x.LogErr("gsInfo, ok := db.G_GSInfoMap[gs] 找不到gs")
		w.Write(CreateErrMsg(sys.CODE_S_GS_ID_NO_EXIST))
		return
	}
	order := r.PostFormValue("order")
	sql := `select producer_order from d_moneycard_web where producer_order='` + order + `'`
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	if len(result.RowArray) > 0 {
		w.Write(CreateSucceedMsg("已补单"))
		return
	}
	sql = `insert into d_moneycard_web (producer_order,sdk_type,sdk_rmb,sdk_ret_time,sdk_ret,sdk_goods,sdk_count) SELECT  producer_order,sdktype,rmb,UNIX_TIMESTAMP(),1,goods,count from d_moneycard d WHERE d.producer_order='%v';`
	sql = fmt.Sprintf(sql, order)
	err1 = gsInfo.ExecSql(sql)
	if err1 != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}
	w.Write(CreateSucceedMsg("成功补单"))
}

//订单查询
func QueryOrderHandle(w http.ResponseWriter, r *http.Request) {
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
	start_time := r.PostFormValue("start_time")
	end_time := r.PostFormValue("end_time")

	start_money := r.PostFormValue("start_money")
	end_money := r.PostFormValue("end_money")

	order_number := r.PostFormValue("order_number")

	user_name := r.PostFormValue("user_name")
	where := " where "
	if start_time != "" && end_time != "" {
		where += `d.create_time between '%v' and '%v' and `
		where = fmt.Sprintf(where, start_time, end_time)
	}
	if start_money != "" && end_money != "" {
		where += `d.sdk_rmb >= '%v' and d.sdk_rmb <= '%v' and `
		where = fmt.Sprintf(where, start_money, end_money)
	}
	if start_money == "" && end_money != "" {
		where += `d.sdk_rmb <='%v' and `
		where = fmt.Sprintf(where, end_money)
	}
	if start_money != "" && end_money == "" {
		where += `d.sdk_rmb >='%v' and `
		where = fmt.Sprintf(where, start_money)
	}
	if order_number != "" {
		where += `d.producer_order='%v' or d.sdk_order='%v' and `
		where = fmt.Sprintf(where, order_number, order_number)
	}
	if user_name != "" {
		where += `u.name='%v' and `
		where = fmt.Sprintf(where, user_name)
	}
	if len(where) <= 7 {
		where = ""
	}
	where = strings.TrimRight(where, "and ")

	pageStr := r.PostFormValue("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		x.LogErr(err)
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "page:参数不能为空"))
		return
	}
	if page <= 0 {
		x.LogErr("page:参数小于等于0")
		w.Write(CreateErrInfoMsg(sys.CODE_PARAM_ERR, "page:参数小于等于0"))
	}
	page = page - 1
	limitStr := ` limit %v,50`
	limitStr = fmt.Sprintf(limitStr, page*50)
	sql := `SELECT d.id,d.producer_order,u.name,u.channel,FROM_UNIXTIME(d.create_time),d.sdk_order,u.channel_accountname,d.sdk_rmb,d.status from d_moneycard d LEFT JOIN d_user u on d.role_id=u.id `
	sql += where + limitStr
	result, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}

	sql = `SELECT count(0) from d_moneycard d LEFT JOIN d_user u on d.role_id=u.id ` + where
	resultPage, err1 := gsInfo.QuerySql(sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}

	info := result.ToTableJson()
	if info == nil {
		x.LogErr(info)
		return
	}
	info2 := resultPage.ToJson()
	if info2 == nil {
		x.LogErr(info2)
		return
	}
	w.Write(CreateSucceedMultiMsg(*info, *info2))
}
