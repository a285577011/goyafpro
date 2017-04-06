package dao

import (
	"../mysql"
)

type MissionLog struct {
	Base
}

var MissionLogDAO *MissionLog

func init() {
	MissionLogDAO = &MissionLog{}
	MissionLogDAO.mysql = mysql.MissionLogMysql
}
