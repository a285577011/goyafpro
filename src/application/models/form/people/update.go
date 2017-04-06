package formPeople

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Update struct {
	form.Base
}

//实例化表单
func NewUpdate(data url.Values) *Update {
	update := &Update{}
	update.Fields = map[string]*form.Field{
		"people_id":       myform.GetField("people_id"),
		"people_sex":      myform.GetField("people_sex"),
		"people_birthday": myform.GetField("people_birthday"),
		"people_height":   myform.GetField("people_height"),
		"people_name":     myform.GetField("people_name"),
		"people_avatar":   myform.GetField("people_avatar"),
	}
	//设置值
	update.SetFieldsValues(data)

	return update
}
