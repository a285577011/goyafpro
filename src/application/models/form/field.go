//字段文件
package form

import (
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/form"
)

//存储所有字段的定义
var fields map[string]form.Field

//从fields中的复制一个出来
func GetField(name string) *form.Field {
	field, ok := fields[name]
	if !ok {
		panic("field " + name + " is not exist.")
	}
	return &field
}

func init() {
	goyaf.Debug("init form field")

	fields = map[string]form.Field{
		"user_id": form.Field{
			Name:     "user_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "user_id不正确",
			}},
		},
		"people_id": form.Field{
			Name:     "people_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "用户编号不正确",
			}},
		},
		"people_sex": form.Field{
			Name:     "people_sex",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Max:    map[string]int{"isset": 1, "value": 2},
				Errmsg: "性别不正确",
			}},
		},
		"people_birthday": form.Field{
			Name:     "people_birthday",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Max:    map[string]int{"isset": 1, "value": 10},
				Errmsg: "生日不正确",
			}},
		},
		"people_height": form.Field{
			Name:     "people_height",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "身高不正确",
			}},
		},
		"people_name": form.Field{
			Name:     "people_name",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Min:    map[string]int{"isset": 1, "value": 1},
				Max:    map[string]int{"isset": 1, "value": 45},
				Errmsg: "people_name错误",
			}},
		},
		"people_avatar": form.Field{
			Name:     "people_avatar",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "string",
				Min:    map[string]int{"isset": 1, "value": 1},
				Max:    map[string]int{"isset": 1, "value": 1000},
				Errmsg: "people_avatar错误",
			}},
		},
		"is_share": form.Field{
			Name:     "is_share",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 0},
				Max:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "is_share错误",
			}},
		},
		"page": form.Field{
			Name:     "page",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "page错误",
			}},
		},
		"count": form.Field{
			Name:     "count",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "count错误",
			}},
		},
		"start_time": form.Field{
			Name:     "start_time",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "开始时间不正确",
			}},
		},
		"end_time": form.Field{
			Name:     "end_time",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "结束时间不正确",
			}},
		},

		//任务相关
		"mission_id": form.Field{
			Name:     "mission_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "mission_id不正确",
			}},
		},
		"log_id": form.Field{
			Name:     "log_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "log_id不正确",
			}},
		},
		//分享相关
		"share_id": form.Field{
			Name:     "share_id",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Errmsg: "share_id不正确",
			}},
		},
	}

}
