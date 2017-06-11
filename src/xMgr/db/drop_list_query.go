package db

import (
	"xMgr/sys"
)

type DropList struct {
	TableName string
	Id        string
	Name      string
}

var DropListMap map[int]*DropList

func init() {
	DropListMap = make(map[int]*DropList, 1)

	dropList := DropList{"s_item_type", "id", "name"}
	DropListMap[sys.TABLE_S_ITEM_TYPE] = &dropList

	dropList1 := DropList{"s_gs", "id", "name"}
	DropListMap[sys.TABLE_S_GS] = &dropList1 ///TABLE_S_SDK_TYPE

	dropList3 := DropList{"s_sdk_type", "id", "name"}
	DropListMap[sys.TABLE_S_SDK_TYPE] = &dropList3

	dropList5 := DropList{"s_gs_group", "id", "name"}
	DropListMap[sys.TABLE_S_GS_GROUP] = &dropList5

	//TABLE_All_S_GS
	dropList4 := DropList{"s_gs", "", ""}
	DropListMap[sys.TABLE_All_S_GS] = &dropList4

}
