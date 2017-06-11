package x

import (
	"database/sql"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

////////////////////////////////////////////
//DBRowSet
type DBRowSet struct {
	RowArray    []*DBRow //字段值
	ColumnsName []string //字段名
}

func (self *DBRowSet) init(columnsName []string) {
	// nColumnNum := len(columnsName)
	// self.RowArray = make([]DBRow, nRowsNum)
	// for i := 0; i < nRowsNum; i++ {
	// 	r := self.RowArray[i]
	// 	r.init(nColumnNum)
	// }
	self.ColumnsName = columnsName
}

// 字段名1 字段名2
// 值1 值2
func (self *DBRowSet) ToString() *string {
	nRowNum := len(self.RowArray)
	nColumn := len(self.ColumnsName)
	var str string
	var str2 string

	for i := 0; i < nColumn; i++ {
		str2 += self.ColumnsName[i] + "\t"
	}
	str = str2 + "\n"

	for i := 0; i < nRowNum; i++ {
		r := self.RowArray[i]
		str2 = ""
		for j := 0; j < nColumn; j++ {
			str2 += r.ColumnArray[j].Value + "\t"
		}
		str = str + str2 + "\n"
	}
	return &str
}

//[值1,值2],[值1,值2]
func (self *DBRowSet) ToAryString() *string {

	nRowNum := len(self.RowArray)
	nColumn := len(self.ColumnsName)
	var str string
	var str2 string

	for i := 0; i < nRowNum; i++ {
		r := self.RowArray[i]
		str2 = "["
		for j := 0; j < nColumn; j++ {
			if j > 1 {
				str2 += ","
			}
			str2 += r.ColumnArray[j].Value
		}
		str = str + str2 + "]"
	}
	return &str
}

//[值1,值2],[值1,值2]
func (self *DBRowSet) ToTableJson() *string {

	nRowNum := len(self.RowArray)
	nColumn := len(self.ColumnsName)
	var str string
	var str2 string

	for i := 0; i < nRowNum; i++ {
		r := self.RowArray[i]
		str2 = "["
		for j := 0; j < nColumn; j++ {
			if j > 1 {
				str2 += ","
			}
			str2 += r.ColumnArray[j].Value
		}
		str = str + str2 + "]"
	}
	return &str
}

func (self *DBRowSet) ToMap() *string {
	return nil
}

func (self *DBRowSet) ToJson() *string {
	return nil
}

////////////////////////////////////////////

func GetResultString(pDB *sql.DB, strSql string) (*string, *Error) {
	res, x := CreateRecordSet(pDB, strSql)
	if x != nil {
		return nil, x
	}
	return res.ToString(), nil
}

func CreateRecordSet(pDB *sql.DB, strSql string) (pRecordSet *DBRowSet, xerr *Error) {
	if pDB == nil {
		return nil, XErrStr("pDB==nil")
	}

	// Execute the query
	rows, err := pDB.Query(strSql)
	if err != nil {
		return nil, XErr(err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, XErr(err)
	}

	nColumnNum := len(columns)
	pRecordSet = &DBRowSet{}
	pRecordSet.init(columns)

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
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			r.setColumnVal(i, value)
		}
		pRecordSet.RowArray = append(pRecordSet.RowArray, r)
		nRow = nRow + 1
	}
	if err = rows.Err(); err != nil {
		return nil, XErr(err)
	}
	return pRecordSet, nil
}
