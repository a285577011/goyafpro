package dao

import (
	"../mysql"
)

type AnalyseConfig struct {
	Base
}

var AnalyseConfigDAO *AnalyseConfig

func init() {
	AnalyseConfigDAO = &AnalyseConfig{}
	AnalyseConfigDAO.mysql = mysql.AnalyseConfigMysql
}
