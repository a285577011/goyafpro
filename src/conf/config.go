package conf

import (
	"git.oschina.net/pbaapp/goyaf"
)

var Config map[string]map[string]string

func init() {
	Config = make(map[string]map[string]string)

	//通用配置
	common := make(map[string]string)
	common["mysql.host"] = "192.168.3.233"
	common["mysql.port"] = "3306"
	common["mysql.database"] = "mushu-youxing"
	common["mysql.username"] = "root"
	common["mysql.password"] = "123456"
	common["mysql.charset"] = "utf8mb4"
	common["mysql.driver_options.1002"] = "SET NAMES utf8mb4"
	//common["mysql.maxconn"] = "1000"
	//common["mysql.maxidleconn"] = "50"

	common["mysql.pbayouxing.host"] = "192.168.3.233"
	common["mysql.pbayouxing.port"] = "3306"
	common["mysql.pbayouxing.database"] = "pba-fat-scale"
	common["mysql.pbayouxing.username"] = "root"
	common["mysql.pbayouxing.password"] = "123456"
	common["mysql.pbayouxing.charset"] = "utf8mb4"
	common["mysql.pbayouxing.driver_options.1002"] = "SET NAMES utf8mb4"
	//common["mysql.pbayouxing.maxconn"] = "1000"
	//common["mysql.pbayouxing.maxidleconn"] = "50"

	common["url-app.pba.cn"] = "http://192.168.3.233"
	common["url-user.mushu.cn"] = "http://192.168.3.233:9110"
	common["http-listen-port"] = "8005"
	common["debugmode"] = "1"

	common["stdout"] = "/Users/zenghui/tmp/youxing-log.log"
	common["stderr"] = "/Users/zenghui/tmp/youxing-error.log"

	Config["common"] = common

	//开发环境配置
	devel := make(map[string]string)
	Config["devel"] = devel

	//测试环境配置
	test := make(map[string]string)
	test["stdout"] = "/data/log/mushu-youxing.log"
	test["stderr"] = "/data/log/mushu-youxing-error.log"
	Config["test"] = test

	//生产环境配置
	product := make(map[string]string)
	product["mysql.host"] = "10.10.149.234"
	product["mysql.database"] = "youxing_new"
	product["mysql.username"] = "mushu"
	product["mysql.password"] = "rVT9d4oNl"

	product["mysql.pbayouxing.host"] = "10.10.149.234"
	product["mysql.pbayouxing.database"] = "youxing"
	product["mysql.pbayouxing.username"] = "mushu"
	product["mysql.pbayouxing.password"] = "rVT9d4oNl"

	product["url-app.pba.cn"] = "http://app.pba.cn"
	product["url-user.mushu.cn"] = "http://user.mushu.cn"

	product["stdout"] = "/data/log/go/log/mushu-youxing.log"
	product["stderr"] = "/data/log/go/log/mushu-youxing-error.log"
	Config["product"] = product

}

func init() {
	goyaf.Debug("init config")
}
