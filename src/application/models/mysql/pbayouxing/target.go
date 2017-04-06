package mysqlPbayouxing

import ()

type Target struct {
	Base
}

var TargetMysql *Target

func init() {
	TargetMysql = &Target{}
	TargetMysql.Table = "target"
	TargetMysql.PrimaryKey = "target_id"
}
