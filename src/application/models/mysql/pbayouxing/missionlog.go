package mysqlPbayouxing

import ()

type MissionLog struct {
	Base
}

var MissionLogMysql *MissionLog

func init() {
	MissionLogMysql = &MissionLog{}
	MissionLogMysql.Table = "mission_log"
	MissionLogMysql.PrimaryKey = "log_id"
}
