package formIapiRecord

import (
	myform "../../../form"
	"git.oschina.net/pbaapp/goyaf"
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
		"user_id":    myform.GetField("user_id"),
		"page":       myform.GetField("page"),
		"count":      myform.GetField("count"),
		"start_time": myform.GetField("start_time"),
		"end_time":   myform.GetField("end_time"),
	}

	list.SetFieldsValues(data)

	return list
}

func init() {
	goyaf.Debug("init form iapi record list")
}
