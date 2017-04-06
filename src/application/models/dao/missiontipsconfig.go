package dao

import (
	"../mysql"
)

type MissionTipsConfig struct {
	Base
}

var MissionTipsConfigDAO *MissionTipsConfig

func init() {
	MissionTipsConfigDAO = &MissionTipsConfig{}
	MissionTipsConfigDAO.mysql = mysql.MissionTipsConfigMysql
}
