package formMission

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Accept struct {
	form.Base
}

//实例化表单
func NewAccept(data url.Values) *Accept {
	accept := &Accept{}
	accept.Fields = map[string]*form.Field{
		"mission_id": myform.GetField("mission_id"),
	}

	accept.SetFieldsValues(data)

	return accept
}
