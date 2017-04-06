package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type Target struct {
	Base
}

//实例化支持事务的表对象
func NewTxTarget(tx *db.Transaction) *Target {
	targetMysql := &Target{}
	targetMysql.Table = TargetMysql.Table
	targetMysql.PrimaryKey = TargetMysql.PrimaryKey
	targetMysql.Tx = tx

	return targetMysql
}

var TargetMysql *Target

func init() {
	TargetMysql = &Target{}
	TargetMysql.Table = "target"
	TargetMysql.PrimaryKey = "target_id"
}
