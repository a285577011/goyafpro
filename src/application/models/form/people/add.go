package formPeople

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Add struct {
	form.Base
}

//实例化表单
func NewAdd(data url.Values) *Add {
	add := &Add{}
	add.Fields = map[string]*form.Field{
		"people_type": &form.Field{
			Name:     "people_type",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Max:    map[string]int{"isset": 1, "value": 2},
				Errmsg: "类型不正确",
			}},
		},
		"people_sex":      myform.GetField("people_sex"),
		"people_birthday": myform.GetField("people_birthday"),
		"people_height":   myform.GetField("people_height"),
		"people_name":     myform.GetField("people_name"),
		"people_avatar":   myform.GetField("people_avatar"),
	}
	//设置值
	add.SetFieldsValues(data)

	return add
}
