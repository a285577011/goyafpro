package mysqlPbayouxing

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type People struct {
	Base
}

//实例化支持事务的表对象
func NewTxPeople(tx *db.Transaction) *People {
	mysql := &People{}
	mysql.Table = mysql.Table
	mysql.PrimaryKey = mysql.PrimaryKey
	mysql.Tx = tx

	return mysql
}

var PeopleMysql *People

func init() {
	PeopleMysql = &People{}
	PeopleMysql.Table = "people"
	PeopleMysql.PrimaryKey = "people_id"
}
