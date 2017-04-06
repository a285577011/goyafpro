package formRecord

import (
	myform "../../form"
	"git.oschina.net/pbaapp/goyaf/form"
	"net/url"
)

type Analyse struct {
	form.Base
}

//实例化表单
func NewAnalyse(data url.Values) *Analyse {
	analyse := &Analyse{}
	analyse.Fields = map[string]*form.Field{
		"people_id": myform.GetField("people_id"),
		"record_weight": &form.Field{
			Name:     "record_weight",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 150.0},
				Errmsg:   "record_weight错误",
			}},
		},
		"record_bmi": &form.Field{
			Name:     "record_bmi",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 99.99},
				Errmsg:   "record_bmi错误",
			}},
		},
		"record_bmr": &form.Field{
			Name:     "record_bmr",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 6000.0},
				Errmsg:   "record_bmr错误",
			}},
		},
		"record_water": &form.Field{
			Name:     "record_water",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 99.99},
				Errmsg:   "record_water错误",
			}},
		},
		"record_fat": &form.Field{
			Name:     "record_fat",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 99.99},
				Errmsg:   "record_fat错误",
			}},
		},
		"record_age": &form.Field{
			Name:     "record_age",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:   "int",
				Min:    map[string]int{"isset": 1, "value": 1},
				Max:    map[string]int{"isset": 1, "value": 100},
				Errmsg: "record_age错误",
			}},
		},
		"record_muscle": &form.Field{
			Name:     "record_muscle",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 99.99},
				Errmsg:   "record_muscle错误",
			}},
		},
		"record_bone": &form.Field{
			Name:     "record_bone",
			Required: true,
			Validate: []form.Validate{form.Validate{
				Type:     "float",
				FloatMin: map[string]float64{"isset": 1.0, "value": 0.01},
				FloatMax: map[string]float64{"isset": 1.0, "value": 99.99},
				Errmsg:   "record_bone错误",
			}},
		},
		"type": &form.Field{
			Name:     "type",
			Required: false,
			Validate: []form.Validate{form.Validate{
				Type:     "int",
				FloatMin: map[string]float64{"isset": 1, "value": 1},
				FloatMax: map[string]float64{"isset": 1, "value": 2},
				Errmsg:   "type错误",
			}},
		},
	}

	analyse.SetFieldsValues(data)

	return analyse
}
