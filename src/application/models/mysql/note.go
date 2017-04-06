package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type Note struct {
	Base
}

//实例化支持事务的表对象
func NewTxNote(tx *db.Transaction) *Note {
	noteMysql := &Note{}
	noteMysql.Table = NoteMysql.Table
	noteMysql.PrimaryKey = NoteMysql.PrimaryKey
	noteMysql.Tx = tx

	return noteMysql
}

var NoteMysql *Note

func init() {
	NoteMysql = &Note{}
	NoteMysql.Table = "note"
	NoteMysql.PrimaryKey = "note_id"
}
