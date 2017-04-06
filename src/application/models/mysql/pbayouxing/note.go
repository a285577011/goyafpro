package mysqlPbayouxing

import ()

type Note struct {
	Base
}

var NoteMysql *Note

func init() {
	NoteMysql = &Note{}
	NoteMysql.Table = "note"
	NoteMysql.PrimaryKey = "note_id"
}
