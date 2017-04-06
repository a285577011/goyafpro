package formPeople

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type AddTarget struct {
	form.Base
}

//实例化表单
func NewAddTarget(data url.Values) *AddTarget {
	addTarget := &AddTarget{}
	addTarget.Fields = map[string]*form.Field{
		"people_id": myform.GetField("people_id"),
		"target_weight": &form.Field{
			Name:     "target_weight",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "target_weight不正确",
			}},
		},
		"is_share": myform.GetField("is_share"),
	}
	//设置值
	addTarget.SetFieldsValues(data)

	return addTarget
}
