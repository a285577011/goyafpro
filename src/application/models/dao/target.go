package dao

import (
	"../mysql"
)

type Target struct {
	Base
}

var TargetDAO *Target

func init() {
	TargetDAO = &Target{}
	TargetDAO.mysql = mysql.TargetMysql
}
