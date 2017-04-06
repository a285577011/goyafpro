package mysql

import ()

type People struct {
	Base
}

var PeopleMysql *People

func init() {
	PeopleMysql = &People{}
	PeopleMysql.Table = "people"
	PeopleMysql.PrimaryKey = "people_id"
}
