package formRecord

import (
	//myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type List struct {
	form.Base
}

//实例化表单
func NewList(data url.Values) *List {
	list := &List{}
	list.Fields = map[string]*form.Field{
		"people_id": &form.Field{
			Name:     "people_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "用户编号不正确",
			}},
		},
		"start_time": &form.Field{
			Name:     "start_time",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "开始时间不正确",
			}},
		},
		"end_time": &form.Field{
			Name:     "end_time",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "结束时间不正确",
			}},
		},
		"page": &form.Field{
			Name:     "page",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "page错误",
			}},
		},
		"count": &form.Field{
			Name:     "count",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "count错误",
			}},
		},
	}

	list.SetFieldsValues(data)

	return list
}
