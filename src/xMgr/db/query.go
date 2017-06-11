package db

import (
	"strconv"
	//"xMgr/user"
	"database/sql"
	"strings"
	"xMgr/sys"
	"xsw/go_pub/x"

	"github.com/astaxie/beego/orm"
)

var (
	//key:s_xxxx.id
	TbGsGroupMap   map[int]*SGsGroup
	TbGseMap       map[int]*SGs        //
	TbUserMap      map[int64]*SUser    //
	TbDbGroupMap   map[int]*SDbGroup   //
	TbUserGroupMap map[int]*SUserGroup //

	TbCodeStringMap x.DBDataMap
	G_GSInfoMap     map[int]*GSInfo //key:idGS
)

type GSInfo struct {
	S_GS_ID       int
	S_DB_GROUP_ID int
	PDB           *sql.DB
}

func (gs *GSInfo) InitPDB() *x.Error {
	if gs.PDB == nil {
		var err error
		oTbGse, ok := TbGseMap[gs.S_GS_ID]
		if !ok {
			return x.XErrStr("TbGseMap[gs.S_GS_ID] 的key不存在")
		} //tb_gs
		tbDbGroup, ok := TbDbGroupMap[oTbGse.DbGroupId] //s_db_group
		if !ok {
			return x.XErrStr("oTbGse.DbGroupId is not exit")
		}
		var url string
		url = tbDbGroup.DbUserName + ":" + tbDbGroup.DbPwd + "@tcp(" + oTbGse.GetDBIp() + ":" + strconv.Itoa(tbDbGroup.DbPort) + ")/" + oTbGse.GetDBName() + "?charset=utf8"
		gs.PDB, err = sql.Open("mysql", url)
		if err != nil {
			return x.XErr(err)
		}
		x.SetDBCharacter(gs.PDB)
	}
	return nil
}
func (gs *GSInfo) QuerySqlTable(sql string) (*string, *x.Error) {
	err := gs.InitPDB()
	if err != nil {
		return nil, err
	}
	sGs, ok := TbGseMap[gs.S_GS_ID]
	if !ok {
		return nil, x.XErrStr("TbGseMap[gs.S_GS_ID] 的key不存在")
	}
	x.LogInfoF("%s,%s,%s", sGs.GetDBIp(), sGs.Name, sql)
	rows, xerr2 := x.CreateRecordSet(gs.PDB, sql)
	if xerr2 != nil {
		return nil, xerr2
	}
	return rowToStr(rows, gs.S_GS_ID, sql), nil
}

func (gs *GSInfo) QuerySql(sql string) (*x.DBRowSet, *x.Error) {
	err := gs.InitPDB()
	if err != nil {
		return nil, err
	}
	sGs, ok := TbGseMap[gs.S_GS_ID]
	if !ok {
		return nil, x.XErrStr("TbGseMap[gs.S_GS_ID] 的key不存在")
	}
	x.LogInfoF("%s,%s,%s", sGs.GetDBIp(), sGs.Name, sql)
	rows, xerr2 := x.CreateRecordSet(gs.PDB, sql)
	if xerr2 != nil {
		return nil, xerr2
	}
	return rows, nil
}

func (gs *GSInfo) ExecSql(sqll string) *x.Error {
	err1 := gs.InitPDB()
	if err1 != nil {
		return err1
	}
	sGs, ok := TbGseMap[gs.S_GS_ID]
	if !ok {
		return x.XErrStr("TbGseMap[gs.S_GS_ID] 的key不存在")
	}
	x.LogInfoF("%s,%s,%s", sGs.GetDBIp(), sGs.Name, sqll)
	_, err := gs.PDB.Exec(sqll)
	if err != nil {
		return x.XErr(err)
	}
	return nil
}
func LoadTable() *x.Error {
	var err error
	var xerr *x.Error

	db, err := orm.GetDB("default")

	//load tb_code_string
	TbCodeStringMap, xerr = x.CreateDBDataMap(db, "SELECT * from s_code_string")
	if xerr != nil {
		return xerr
	}

	///load tb_gs_group
	if TbGsGroupMap == nil {
		var g_arryTbGsGroup []*SGsGroup
		qs1 := G_OrmDB.QueryTable("s_gs_group")
		_, err = qs1.All(&g_arryTbGsGroup)
		if err != nil {
			return x.XErr(err)
		}
		num := len(g_arryTbGsGroup)
		TbGsGroupMap = make(map[int]*SGsGroup, num)
		for _, s := range g_arryTbGsGroup {
			TbGsGroupMap[s.Id] = s
		}
	}

	///load tb_gs
	if TbGseMap == nil {
		var g_arryTbGse []*SGs
		qs1 := G_OrmDB.QueryTable("s_gs")
		_, err = qs1.All(&g_arryTbGse)
		if err != nil {
			return x.XErr(err)
		}
		num := len(g_arryTbGse)
		TbGseMap = make(map[int]*SGs, num)
		G_GSInfoMap = make(map[int]*GSInfo, num)
		for _, s := range g_arryTbGse {
			var gSInfo GSInfo
			//gSInfo = new(GSInfo)
			gSInfo.S_GS_ID = s.Id
			gSInfo.S_DB_GROUP_ID = s.DbGroupId
			TbGseMap[s.Id] = s
			G_GSInfoMap[s.Id] = &gSInfo
		}
	}

	//load tb_user
	if TbUserMap == nil {
		var G_arryTbUser []*SUser
		qs1 := G_OrmDB.QueryTable("s_user")
		_, err = qs1.All(&G_arryTbUser)
		if err != nil {
			return x.XErr(err)
		}
		num := len(G_arryTbUser)
		TbUserMap = make(map[int64]*SUser, num)
		for _, s := range G_arryTbUser {
			TbUserMap[s.Id] = s
		}
	}

	//load tb_db_group
	if TbDbGroupMap == nil {
		var g_arryTbDbGroup []*SDbGroup
		qs1 := G_OrmDB.QueryTable("s_db_group")
		_, err = qs1.All(&g_arryTbDbGroup)
		if err != nil {
			return x.XErr(err)
		}
		num := len(g_arryTbDbGroup)
		TbDbGroupMap = make(map[int]*SDbGroup, num)
		for _, s := range g_arryTbDbGroup {
			TbDbGroupMap[s.Id] = s
		}
	}
	//load tb_user_group
	if TbUserGroupMap == nil {
		var g_arryTbDbGroup []*SUserGroup
		qs1 := G_OrmDB.QueryTable("s_user_group")
		_, err = qs1.All(&g_arryTbDbGroup)
		if err != nil {
			return x.XErr(err)
		}
		num := len(g_arryTbDbGroup)
		TbUserGroupMap = make(map[int]*SUserGroup, num)
		for _, s := range g_arryTbDbGroup {
			TbUserGroupMap[s.Id] = s
		}
	}

	x.PrintInfo("LoadTable suc")
	return nil
}

//得到有产品的字符串
func GetAllProductJstr() string {
	var sumStr string
	var str string
	var arrInt []int
	arrInt = make([]int, len(TbGsGroupMap))
	sumStr = `{"cmd":2, "num":` + strconv.Itoa(len(TbGsGroupMap)) + `,"p":[`
	i := 0
	for _, value := range TbGsGroupMap {
		arrInt[i] = value.Id
		i++
	}
	arrInt = sys.SortArryInt(arrInt) //排序数组

	for _, value := range arrInt {
		str += getProductJStr(TbGsGroupMap[value]) + `,`
	}
	str = strings.TrimRight(str, ",")
	sumStr = sumStr + str + `]}`
	return sumStr
}

//tb_gs_group的id相同的数据为s key为tb_gs_group表对应的数组下标
func getProductJStr(product *SGsGroup) string {
	var str, sumStr string
	ss := selectTbGsGroupIdMap(product.Id)
	ss = sortArrySGS(ss)
	sumStr = `{"name":"` + product.Name + `",` + `"gs_num":` + strconv.Itoa(len(ss)) + `,"gs_group_id":` + strconv.Itoa(product.Id) + `,"gs":[`
	for _, value := range ss {
		str += getGsJStr(ss, value.GsId) + `,`
	}
	str = strings.TrimRight(str, ",")
	sumStr += str + `]}`
	return sumStr
}

//排序SGS数组
func sortArrySGS(arr []*SGs) []*SGs {
	var temp *SGs
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i].Id < arr[j].Id {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

//查询与tb_gs_group的id相同的数据 key为tb_gs_group表的id
func selectTbGsGroupIdMap(tb_gs_group_id int) []*SGs {
	var pTbGses []*SGs
	var arrInt []int
	var i int = 0
	for _, s := range TbGseMap {
		if s.GsGroupId == tb_gs_group_id {
			i++
		}
	}
	pTbGses = make([]*SGs, i)
	arrInt = make([]int, i)
	i = 0
	for _, s := range TbGseMap {
		if s.GsGroupId == tb_gs_group_id {
			arrInt[i] = s.Id
			i++
		}
	}

	arrInt = sys.SortArryInt(arrInt)
	for i, s := range arrInt {
		pTbGses[i] = TbGseMap[s]
	}
	return pTbGses
}

//查询用gs_id来tb_gs相同的gs_group_id的行得到部分json字符串
func getGsJStr(s []*SGs, gs_id int) string {
	var str string
	for _, u := range s {
		if u.GsId == gs_id {
			str = `{"name":"` + u.Name + `","gs_id":` + strconv.Itoa(u.GsId) + `,"id":` + strconv.Itoa(u.Id) + `}`
			return str
		}
	}
	return str
}

//把row转换为string
func rowToStr(self *x.DBRowSet, s_gs_id int, sqll string) *string {
	nRowNum := len(self.RowArray)
	nColumn := len(self.ColumnsName)
	var numstr string
	var str string
	numstr = `sql：` + sqll + `</br>`
	numstr += `<table  border="1">`
	str = `<tr>`
	for i := 0; i < nColumn; i++ {
		str += `<td>` + self.ColumnsName[i] + `</td>`
		//str += self.ColumnsName[i] + "\t"
	}
	str += `</tr>`
	numstr += str

	for i := 0; i < nRowNum; i++ {
		r := self.RowArray[i]
		str = `<tr>`
		for j := 0; j < nColumn; j++ {
			str += `<td>` + r.ColumnArray[j].Value + `</td>`
		}
		str += `</tr>`
		numstr = numstr + str
	}
	numstr += `</table>`
	return &numstr
}
