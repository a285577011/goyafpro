package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type Migrate struct {
	Base
}

//实例化支持事务的表对象
func NewTxMigrate(tx *db.Transaction) *Migrate {
	mysql := &Migrate{}
	mysql.Table = MigrateMysql.Table
	mysql.PrimaryKey = MigrateMysql.PrimaryKey
	mysql.Tx = tx

	return mysql
}

var MigrateMysql *Migrate

func init() {
	MigrateMysql = &Migrate{}
	MigrateMysql.Table = "migrate"
}
