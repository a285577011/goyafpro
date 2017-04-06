package mysql

import ()

type AnalyseConfig struct {
	Base
}

var AnalyseConfigMysql *AnalyseConfig

func init() {
	AnalyseConfigMysql = &AnalyseConfig{}
	AnalyseConfigMysql.Table = "analyse_config"
	AnalyseConfigMysql.PrimaryKey = "config_id"
}
