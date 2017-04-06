package dao

import (
	"../mysql"
)

type People struct {
	Base
}

var PeopleDAO *People

func init() {
	PeopleDAO = &People{}
	PeopleDAO.mysql = mysql.PeopleMysql
}
