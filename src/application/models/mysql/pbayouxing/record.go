package mysqlPbayouxing

import ()

type Record struct {
	Base
}

var RecordMysql *Record

func init() {
	RecordMysql = &Record{}
	RecordMysql.Table = "record"
	RecordMysql.PrimaryKey = "record_id"
}
