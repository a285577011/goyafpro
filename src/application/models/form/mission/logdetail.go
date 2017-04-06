package formMission

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type LogDetail struct {
	form.Base
}

//实例化表单
func NewLogDetail(data url.Values) *LogDetail {
	logDetail := &LogDetail{}
	logDetail.Fields = map[string]*form.Field{
		"log_id": myform.GetField("log_id"),
	}

	logDetail.SetFieldsValues(data)

	return logDetail
}
