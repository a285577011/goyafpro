package formPeople

import (
	myform "../../form"
	"fmt"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Info struct {
	form.Base
}

//实例化表单
func NewInfo(data url.Values) *Info {
	info := &Info{}
	info.Fields = map[string]*form.Field{
		"people_id": myform.GetField("people_id"),
	}

	info.SetFieldsValues(data)

	return info
}

func init() {
	fmt.Println("init form people info")
}
