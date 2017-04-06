package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type MissionTipsConfig struct {
	Base
}

//实例化支持事务的表对象
func NewTxMissionTipsConfig(tx *db.Transaction) *MissionTipsConfig {
	mysql := &MissionTipsConfig{}
	mysql.Table = MissionTipsConfigMysql.Table
	mysql.PrimaryKey = MissionTipsConfigMysql.PrimaryKey
	mysql.Tx = tx

	return mysql
}

var MissionTipsConfigMysql *MissionTipsConfig

func init() {
	MissionTipsConfigMysql = &MissionTipsConfig{}
	MissionTipsConfigMysql.Table = "mission_tips_config"
	MissionTipsConfigMysql.PrimaryKey = "config_id"
}
