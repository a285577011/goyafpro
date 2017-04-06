//通用的表单

package form

import (
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Common struct {
	Base
}

//实例化表单
func NewCommon(data url.Values, fieldNames []string) *Common {
	common := &Common{}
	common.Fields = make(map[string]*form.Field)

	fieldsIni := map[string]*form.Field{
		"people_id": &form.Field{
			Name:     "people_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "用户编号不正确",
			}},
		},
		"page": &form.Field{
			Name:     "page",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "page错误",
			}},
		},
		"count": &form.Field{
			Name:     "count",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "count错误",
			}},
		},
	}

	for _, fieldName := range fieldNames {
		field, ok := fieldsIni[fieldName]
		if ok {
			common.Fields[fieldName] = field
		}
	}

	common.SetFieldsValues(data)

	return common
}
