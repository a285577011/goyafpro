package formShare

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Content struct {
	form.Base
}

//实例化表单
func NewContent(data url.Values) *Content {
	content := &Content{}
	content.Fields = map[string]*form.Field{
		"share_id": myform.GetField("share_id"),
	}
	//设置值
	content.SetFieldsValues(data)

	return content
}

func init() {
	goyaf.Debug("init form share content")
}
