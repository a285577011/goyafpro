package mysql

import ()

type Share struct {
	Base
}

var ShareMysql *Share

func init() {
	ShareMysql = &Share{}
	ShareMysql.Table = "share"
	ShareMysql.PrimaryKey = "share_id"
}
