package dao

import (
	"../mysql"
)

type Note struct {
	Base
}

var NoteDAO *Note

func init() {
	NoteDAO = &Note{}
	NoteDAO.mysql = mysql.NoteMysql
}
