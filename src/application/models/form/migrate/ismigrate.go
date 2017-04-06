package formMigrate

import (
	//myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type IsMigrate struct {
	form.Base
}

//实例化表单
func NewIsMigrate(data url.Values) *IsMigrate {
	f := &IsMigrate{}
	f.Fields = map[string]*form.Field{
		"pba_user_id": &form.Field{
			Name:     "pba_user_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "pba_user_id错误",
			}},
		},
	}

	f.SetFieldsValues(data)
	return f
}
