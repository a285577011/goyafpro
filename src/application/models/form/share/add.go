package formShare

import (
	//myform "../../form"
	"encoding/json"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/form"
	"git.oschina.net/pbaapp/goyaf/validate"
	"net/url"
)

type Add struct {
	form.Base
}

//实例化表单
func NewAdd(data url.Values) *Add {
	add := &Add{}
	add.Fields = map[string]*form.Field{
		"share_type": &form.Field{
			Name:     "share_type",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "share_type不正确",
			}},
		},
		"share_content": &form.Field{
			Name:     "share_content",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "share_content不能为空",
			}, form.Validate{
				Type: "func",
				Func: func(value string) bool {
					var j map[string]string
					err := json.Unmarshal([]byte(value), &j)
					if err != nil {
						return false
					}
					goyaf.Debug(j)
					if j["type"] != "1" && j["type"] != "2" {
						return false
					}
					v := validate.NewFloat(map[string]interface{}{
						"value": j["num"],
						"min":   0.1,
					})
					if !v.Validate() {
						return false
					}

					return true
				},
				Errmsg: "share_content不是正确的json格式数据",
			}},
		},
	}
	//设置值
	add.SetFieldsValues(data)

	return add
}

func init() {
	goyaf.Debug("init form share add")
}
