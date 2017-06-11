package x

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

////////////////////////////////////////////
//dBDataMap
type dBDataMap struct {
	DBDataMap
	rowMap      map[int64]*DBRow //字段值
	columnsName []string         //字段名
}

func (self *dBDataMap) GetObj(nKey int64) *DBRow {
	if self.rowMap == nil {
		return nil
	}
	return self.rowMap[nKey]
}

func (self *dBDataMap) AddObj(nKey int64, row *DBRow) bool {
	if self.rowMap == nil {
		return false
	}
	self.rowMap[nKey] = row
	return true
}

func (self *dBDataMap) GetAmount() int {
	if self.rowMap == nil {
		return 0
	}
	return len(self.rowMap)
}

func (self *dBDataMap) Clear() {
	if self.rowMap == nil {
		return
	}
	self.rowMap = make(map[int64]*DBRow)
}

func (self *dBDataMap) init(columnsName []string) {
	self.columnsName = columnsName
	self.rowMap = make(map[int64]*DBRow)
}

// 字段名1 字段名2
// 值1 值2
func (self *dBDataMap) ToString() *string {
	nColumn := len(self.columnsName)
	var str string
	var str2 string

	for i := 0; i < nColumn; i++ {
		str2 += self.columnsName[i] + "\t"
	}
	str = str2 + "\n"

	for _, r := range self.rowMap {
		str2 = ""
		for j := 0; j < nColumn; j++ {
			str2 += r.ColumnArray[j].Value + "\t"
		}
		str = str + str2 + "\n"
	}
	return &str
}

////////////////////////////////////////////
func CreateDBDataMap(pDB *sql.DB, strSql string) (pDataMap2 DBDataMap, xerr *Error) {
	if pDB == nil {
		return nil, XErrStr("pDB==nil")
	}

	// Execute the query
	rows, err := pDB.Query(strSql)
	if err != nil {
		return nil, XErr(err)
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, XErr(err)
	}

	nColumnNum := len(columns)
	pDataMap := new(dBDataMap)
	pDataMap.init(columns)

	// Make a slice for the values
	values := make([]sql.RawBytes, nColumnNum)
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	nRow := 0
	var value string
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, XErr(err)
		}
		// Now do something with the data.
		// Here we just print each column as a string.
		r := new(DBRow)
		r.init(nColumnNum)
		var nKey int64
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = ""
			} else {
				value = string(col)
				if i == 0 {
					nKey, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						return nil, XErr(err)
					}
				}
			}
			r.setColumnVal(i, value)
		}
		if pDataMap.rowMap[nKey] != nil {
			return nil, XErrStr("id重复")
		}
		pDataMap.rowMap[nKey] = r
		nRow = nRow + 1
	}
	if err = rows.Err(); err != nil {
		return nil, XErr(err)
	}
	return pDataMap, nil
}
