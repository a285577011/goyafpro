package mysql

import (
	"git.oschina.net/pbaapp/goyaf/db"
)

type Sync struct {
	Base
}

//实例化支持事务的表对象
func NewTxSync(tx *db.Transaction) *Sync {
	syncMysql := &Sync{}
	syncMysql.Table = SyncMysql.Table
	syncMysql.PrimaryKey = SyncMysql.PrimaryKey
	syncMysql.Tx = tx

	return syncMysql
}

var SyncMysql *Sync

func init() {
	SyncMysql = &Sync{}
	SyncMysql.Table = "sync"
	SyncMysql.PrimaryKey = "sync_id"
}
