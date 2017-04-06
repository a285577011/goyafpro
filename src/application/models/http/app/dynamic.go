package ourHttpUser

import (
	"git.oschina.net/pbaapp/goyaf"
	"reflect"
)

type Dynamic struct {
	Base
}

//解析sso
func (this *Dynamic) Add(params map[string]string) (id string, err error) {
	this.Method = "POST"
	this.Uri = "/iapi/dynamic/add/"
	this.Params = params

	var data interface{}
	data, err = this.Base.Request()
	if err != nil {
		goyaf.ErrorLog.Println("add dynamic error:", err, ",url:", this.Url, ",params:", params)
		return
	}

	v := reflect.ValueOf(data)
	i := v.Interface()
	a := i.(map[string]interface{})

	id = reflect.ValueOf(a["dynamic_id"]).String()

	return
}

func init() {
	goyaf.Log("init http app dynamic")
}
