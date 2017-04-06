package formMigrate

import (
	//myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Migrate struct {
	form.Base
}

//实例化表单
func NewMigrate(data url.Values) *Migrate {
	accept := &Migrate{}
	accept.Fields = map[string]*form.Field{
		"pba_user_id": &form.Field{
			Name:     "pba_user_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "pba_user_id错误",
			}},
		},
		"mushu_account": &form.Field{
			Name:     "mushu_account",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "mushu_account错误",
			}},
		},
		"mushu_password": &form.Field{
			Name:     "mushu_password",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "mushu_password错误",
			}},
		},
	}

	accept.SetFieldsValues(data)

	return accept
}
