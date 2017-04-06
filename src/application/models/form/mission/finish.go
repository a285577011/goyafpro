package formMission

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Finish struct {
	form.Base
}

//实例化表单
func NewFinish(data url.Values) *Finish {
	finish := &Finish{}
	finish.Fields = map[string]*form.Field{
		"log_id": myform.GetField("log_id"),
	}

	finish.SetFieldsValues(data)

	return finish
}
