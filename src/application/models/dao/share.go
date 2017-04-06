package dao

import (
	"../mysql"
)

type Share struct {
	Base
}

var ShareDAO *Share

func init() {
	ShareDAO = &Share{}
	ShareDAO.mysql = mysql.ShareMysql
}
