package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type Mission struct {
	Base
}

//实例化支持事务的表对象
func NewTxMission(tx *db.Transaction) *Mission {
	noteMysql := &Mission{}
	noteMysql.Table = MissionMysql.Table
	noteMysql.PrimaryKey = MissionMysql.PrimaryKey
	noteMysql.Tx = tx

	return noteMysql
}

var MissionMysql *Mission

func init() {
	MissionMysql = &Mission{}
	MissionMysql.Table = "mission"
	MissionMysql.PrimaryKey = "mission_id"
}
