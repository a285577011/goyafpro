package mysql

import (
	"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
)

func (this *Base) GetAdapter() *db.Adapter {
	if len(this.Options) == 0 {
		this.setOptions()
	}
	return db.NewAdapter(this.Options)
}

type Mysql interface {
	Find(id string, fields ...string) map[string]string
	Insert(data map[string]string) int64
	Update(data map[string]string, where map[string]string) int64
	FetchAll(slt db.Select) []map[string]string
}

type Base struct {
	Table      string
	PrimaryKey string
	Options    map[string]string
	Adapter    *db.Adapter
	Tx         *db.Transaction
}

//查询单条记录
func (this *Base) Find(id string, fields ...string) map[string]string {
	tableGateway := db.NewTable(this.Table, this.GetAdapter())

	slt := db.Select{Where: map[string]string{this.PrimaryKey: id}}
	if len(fields) > 0 {
		slt = db.Select{
			Columns: fields[0],
			Where:   map[string]string{this.PrimaryKey: id},
		}
	}

	result, err := tableGateway.Select(slt)
	if err != nil {
		goyaf.ErrorLog.Println("mysql find error:", err)
	}
	if len(result) == 0 {
		return make(map[string]string)
	}
	return result[0]
}

//插入数据
func (this *Base) Insert(data map[string]string) int64 {
	tableGateway := db.NewTable(this.Table, this.GetAdapter())
	tableGateway.SetTx(this.Tx)

	result, err := tableGateway.Insert(data)
	if err != nil {
		goyaf.ErrorLog.Println("mysql insert error:", err)
		panic(err)
	}
	return result
}

//查询单条记录
func (this *Base) Update(data map[string]string, where map[string]string) int64 {
	tableGateway := db.NewTable(this.Table, this.GetAdapter())
	result, err := tableGateway.Update(data, where)
	if err != nil {
		goyaf.ErrorLog.Println("mysql update error:", err)
	}
	return result
}

//查询列表数据
func (this *Base) FetchAll(slt db.Select) []map[string]string {
	tableGateway := db.NewTable(this.Table, this.GetAdapter())
	result, err := tableGateway.Select(slt)
	if err != nil {
		goyaf.ErrorLog.Println("mysql fetchAll error:", err)
	}
	return result
}

//设置数据库连接参数
func (this *Base) setOptions() {
	this.Options = map[string]string{
		"driver":      "mysql",
		"host":        goyaf.GetConfigByKey("mysql.host"),
		"port":        goyaf.GetConfigByKey("mysql.port"),
		"database":    goyaf.GetConfigByKey("mysql.database"),
		"username":    goyaf.GetConfigByKey("mysql.username"),
		"password":    goyaf.GetConfigByKey("mysql.password"),
		"charset":     goyaf.GetConfigByKey("mysql.charset"),
		"maxconn":     goyaf.GetConfigByKey("mysql.maxconn"),
		"maxidleconn": goyaf.GetConfigByKey("mysql.maxidleconn"),
	}
}
func init() {
	fmt.Println("init mysql base")
}
