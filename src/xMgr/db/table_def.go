package db

import (
	"fmt"
	"strconv"
	"time"
	"xMgr/sys"
	"xsw/go_pub/x"
)

//////////////////////////////////////////////////////////
//放置表的字段名
var TableFiledMap map[string][]string

//////////////////////////////////////////////////////////
// SUserGroup
type SUserGroup struct {
	Id        int    `"k"`
	Name      string ``
	Privilege string ``
}

//////////////////////////////////////////////////////////
// SDbGroup
type SDbGroup struct {
	Id         int    `pk`
	Name       string ``
	DbUserName string ``
	DbPwd      string ``
	DbIp       string ``
	DbPort     int    ``
}

func (self *SDbGroup) TableName() string {

	return "s_db_group"
}

func (self *SDbGroup) GetDBIp(gs *SGs) string {
	if len(self.DbIp) <= 0 {
		return gs.Ip
	}
	return self.DbIp
}

func (self *SDbGroup) GetDBPort() int {
	if self.DbPort <= 0 {
		return 3306
	}
	return self.DbPort
}

func (self *SDbGroup) GetExtra() string {
	return "charset=utf8"
}

//////////////////////////////////////////////////////////
// TbGsGroup
type SGsGroup struct {
	Id   int    `"pk"`
	Name string `orm:"size(256)"`
}

//////////////////////////////////////////////////////////
// s_code_string
type SCodeString struct {
	Id  int    `pk`
	Str string ``
}

// func (self *SCodeString) TableName() string {

// 	return "s_code_string"
// }

//////////////////////////////////////////////////////////
//s_gs表结构体
type SGs struct {
	Id        int       ` index:"pk"`
	Name      string    `orm:"size(256)"`
	Name2     string    `orm:"size(256)"`
	GsGroupId int       ``
	GsId      int       ``
	Ip        string    `orm:"size(256)"`
	GsPort    int       ``
	DbGroupId int       ``
	DbName    string    `orm:"size(256)"`
	StartTime time.Time ``
}

func (self *SGs) GetDBIp() string {
	if len(self.Ip) > 0 {
		return self.Ip
	}
	dbGroup, ok := TbDbGroupMap[self.DbGroupId]
	if !ok {
		x.LogErr("TbDbGroupMap[self.DbGroupId]的key不存在")
		return ""
	}
	return dbGroup.DbIp
}
func (self *SGs) GetDBName() string {
	if len(self.DbName) <= 0 {
		fmt.Println("dbname:", self.DbName, "  id:", self.Id)
		return fmt.Sprintf("hotgame_hl%d", self.GsId)
	}
	fmt.Println("dbname2:", self.DbName, "  id:", self.Id)
	return self.DbName
}

func (self *SGs) GetDBGS2String(sign string) string {
	str := "area:" + self.Name + sign
	str += "id:" + strconv.Itoa(self.Id) + sys.GetBlankString(3) + "name:" + self.Name + sys.GetBlankString(3) + "db_name:" + self.DbName + sys.GetBlankString(3) + `ip:` + self.Ip
	str += sys.GetBlankString(3) + "gs_id:" + strconv.Itoa(self.GsId) + sign
	return str
}

//////////////////////////////////////////////////////////
//s_user表结构体
type SUser struct {
	Id          int64  `pk`
	Name        string ``
	Pwd         string ``
	UserGroupId int    ``
	Session     string ``
}

type SUserAry struct {
	sys.IDBAryToString
	aryData []*SUser
}

func (self *SUserAry) ToString() (str string) {
	for _, v := range self.aryData {
		str += sys.ObjToString(v)
	}
	return str
}

//////////////////////////////////////////////////////////
//s_sdk_type
type SSdkType struct {
	Id   int    ``
	Name string ``
}

func (u *SSdkType) TableUnique() [][]string {
	return [][]string{
		{"id"},
	}
}

//////////////////////////////////////////////////////////
// s_item_type
type SItemType struct {
	Id                 int    `pk`
	Name               string ``
	money_val          int    ``
	package_id         int    ``
	equip_create_point int    ``
}

func (self *SItemType) TableName() string {

	return "s_item_type"
}
