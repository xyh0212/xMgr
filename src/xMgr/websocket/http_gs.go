package websocket

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"net/http"
	"strconv"
	"xMgr/db"
	"xMgr/sys"
	"xsw/go_pub/x"
)

/////////////////////////游服//////////////////////////////////////
//获取表的所有数据
func TableHandle(w http.ResponseWriter, r *http.Request) {
	x.PrintInfo("TableHandle")
	r.ParseForm()
	gsStr := r.PostFormValue("table_id") //String2ArryInt

	id, err := strconv.Atoi(gsStr)
	if err != nil {
		x.LogErr(x.XErr(err))
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}
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

	dropList, ok := db.DropListMap[id]
	if !ok {
		x.LogErr(`db.DropListMap[id]的key不存在`)
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}
	sql := `SELECT * from ` + dropList.TableName + limitStr
	pdb, err := orm.GetDB()
	x.SetDBCharacter(pdb)
	if err != nil {
		x.LogErr(x.XErr(err))
		w.Write(CreateErrMsg(sys.CODE_ERR))
		return
	}
	rows, err1 := x.CreateRecordSet(pdb, sql)
	if err != nil {
		return
	}

	sql = `SELECT count(0) from ` + dropList.TableName
	rows2, err1 := x.CreateRecordSet(pdb, sql)
	if err1 != nil {
		x.LogErr(err1)
		w.Write(CreateErrInfoMsg(sys.CODE_SQL_SYNTXT_ERR, err1.GetStr()))
		return
	}

	w.Write(CreateSucceedMultiMsg(*rows.ToTableJson(), *rows2.ToJson()))
}
