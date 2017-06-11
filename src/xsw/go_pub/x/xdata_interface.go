package x

import ()

type Obj interface {
	GetID() int64
}

/////////////////////////////////////////////////////////
//map
type ObjMap interface {
	GetObj(int64) Obj
	AddObj(int64, Obj) bool
	GetAmount() int
}

type DBDataMap interface {
	GetObj(int64) *DBRow
	AddObj(int64, *DBRow) bool
	GetAmount() int
	Clear()
}

/////////////////////////////////////////////////////////
//Vector
