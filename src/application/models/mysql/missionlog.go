package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type MissionLog struct {
	Base
}

//实例化支持事务的表对象
func NewTxMissionLog(tx *db.Transaction) *MissionLog {
	noteMysql := &MissionLog{}
	noteMysql.Table = MissionLogMysql.Table
	noteMysql.PrimaryKey = MissionLogMysql.PrimaryKey
	noteMysql.Tx = tx

	return noteMysql
}

var MissionLogMysql *MissionLog

func init() {
	MissionLogMysql = &MissionLog{}
	MissionLogMysql.Table = "mission_log"
	MissionLogMysql.PrimaryKey = "log_id"
}
