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
		"note_content": &form.Field{
			Name:     "note_content",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Min:    map[string]int{"isset": 1, "value": 1},
				Max:    map[string]int{"isset": 1, "value": 140},
				Errmsg: "笔记内容长度不对",
			}},
		},
		"is_share": myform.GetField("is_share"),
	}

	add.SetFieldsValues(data)

	return add
}
