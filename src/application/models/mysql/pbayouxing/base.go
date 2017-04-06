package mysqlPbayouxing

import (
	"../../mysql"
	"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
)

type Base struct {
	mysql.Base
}

//查询单条记录
func (this *Base) Find(id string, fields ...string) map[string]string {
	this.setOptions()
	if len(fields) > 0 {
		return this.Base.Find(id, fields[0])
	}
	return this.Base.Find(id)
}

//插入数据
func (this *Base) Insert(data map[string]string) int64 {
	this.setOptions()
	return this.Base.Insert(data)
}

//查询单条记录
func (this *Base) Update(data map[string]string, where map[string]string) int64 {
	this.setOptions()
	return this.Base.Update(data, where)
}

//查询列表数据
func (this *Base) FetchAll(slt db.Select) []map[string]string {
	this.setOptions()
	return this.Base.FetchAll(slt)
}

//设置数据库连接参数
func (this *Base) setOptions() {
	this.Options = map[string]string{
		"driver":      "mysql",
		"host":        goyaf.GetConfigByKey("mysql.pbayouxing.host"),
		"port":        goyaf.GetConfigByKey("mysql.pbayouxing.port"),
		"database":    goyaf.GetConfigByKey("mysql.pbayouxing.database"),
		"username":    goyaf.GetConfigByKey("mysql.pbayouxing.username"),
		"password":    goyaf.GetConfigByKey("mysql.pbayouxing.password"),
		"charset":     goyaf.GetConfigByKey("mysql.pbayouxing.charset"),
		"maxconn":     goyaf.GetConfigByKey("mysql.pbayouxing.maxconn"),
		"maxidleconn": goyaf.GetConfigByKey("mysql.pbayouxing.maxidleconn"),
	}
}

func init() {
	fmt.Println("init mysql pbayouxing base")
}
