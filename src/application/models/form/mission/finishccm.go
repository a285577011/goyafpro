package formMission

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type FinishCCM struct {
	form.Base
}

//实例化表单
func NewFinishCCM(data url.Values) *FinishCCM {
	finish := &FinishCCM{}
	finish.Fields = map[string]*form.Field{
		"mission_id": myform.GetField("mission_id"),
		"log_id":     myform.GetField("log_id"),
		"day": &form.Field{
			Name:     "day",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "day不正确",
			}},
		},
	}

	finish.SetFieldsValues(data)

	return finish
}
