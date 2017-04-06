package dao

import (
	"../mysql"
	"fmt"
	"git.oschina.net/pbaapp/goyaf/db"
	"strconv"
)

type Record struct {
	Base
}

func (this *Record) Find(id string, fields ...string) map[string]string {
	result := this.Base.Find(id)
	rw, _ := strconv.ParseFloat(result["record_weight"], 64)
	result["record_weight"] = strconv.FormatFloat(rw/100, 'f', -1, 64)

	rbmi, _ := strconv.ParseFloat(result["record_bmi"], 64)
	result["record_bmi"] = strconv.FormatFloat(rbmi/100, 'f', -1, 64)

	rbmr, _ := strconv.ParseFloat(result["record_bmr"], 64)
	result["record_bmr"] = strconv.FormatFloat(rbmr/100, 'f', -1, 64)

	rwater, _ := strconv.ParseFloat(result["record_water"], 64)
	result["record_water"] = strconv.FormatFloat(rwater/100, 'f', -1, 64)

	rfat, _ := strconv.ParseFloat(result["record_fat"], 64)
	result["record_fat"] = strconv.FormatFloat(rfat/100, 'f', -1, 64)

	rmuscle, _ := strconv.ParseFloat(result["record_muscle"], 64)
	result["record_muscle"] = strconv.FormatFloat(rmuscle/100, 'f', -1, 64)

	rbone, _ := strconv.ParseFloat(result["record_bone"], 64)
	result["record_bone"] = strconv.FormatFloat(rbone/100, 'f', -1, 64)

	return result
}

func (this *Record) Insert(data map[string]string) int64 {
	rw, _ := strconv.ParseFloat(data["record_weight"], 64)
	data["record_weight"] = strconv.FormatFloat(rw*100, 'f', -1, 64)

	rbmi, _ := strconv.ParseFloat(data["record_bmi"], 64)
	data["record_bmi"] = strconv.FormatFloat(rbmi*100, 'f', -1, 64)

	rbmr, _ := strconv.ParseFloat(data["record_bmr"], 64)
	data["record_bmr"] = strconv.FormatFloat(rbmr*100, 'f', -1, 64)

	rwater, _ := strconv.ParseFloat(data["record_water"], 64)
	data["record_water"] = strconv.FormatFloat(rwater*100, 'f', -1, 64)

	rfat, _ := strconv.ParseFloat(data["record_fat"], 64)
	data["record_fat"] = strconv.FormatFloat(rfat*100, 'f', -1, 64)

	rmuscle, _ := strconv.ParseFloat(data["record_muscle"], 64)
	data["record_muscle"] = strconv.FormatFloat(rmuscle*100, 'f', -1, 64)

	rbone, _ := strconv.ParseFloat(data["record_bone"], 64)
	data["record_bone"] = strconv.FormatFloat(rbone*100, 'f', -1, 64)

	return this.Base.Insert(data)
}

func (this *Record) FetchAll(slt db.Select) []map[string]string {
	result := this.Base.FetchAll(slt)
	for k, v := range result {
		rw, _ := strconv.ParseFloat(v["record_weight"], 64)
		result[k]["record_weight"] = strconv.FormatFloat(rw/100, 'f', -1, 64)

		rbmi, _ := strconv.ParseFloat(v["record_bmi"], 64)
		result[k]["record_bmi"] = strconv.FormatFloat(rbmi/100, 'f', -1, 64)

		rbmr, _ := strconv.ParseFloat(v["record_bmr"], 64)
		result[k]["record_bmr"] = strconv.FormatFloat(rbmr/100, 'f', -1, 64)

		rwater, _ := strconv.ParseFloat(v["record_water"], 64)
		result[k]["record_water"] = strconv.FormatFloat(rwater/100, 'f', -1, 64)

		rfat, _ := strconv.ParseFloat(v["record_fat"], 64)
		result[k]["record_fat"] = strconv.FormatFloat(rfat/100, 'f', -1, 64)

		rmuscle, _ := strconv.ParseFloat(v["record_muscle"], 64)
		result[k]["record_muscle"] = strconv.FormatFloat(rmuscle/100, 'f', -1, 64)

		rbone, _ := strconv.ParseFloat(v["record_bone"], 64)
		result[k]["record_bone"] = strconv.FormatFloat(rbone/100, 'f', -1, 64)
	}
	return result
}

var RecordDAO *Record

func init() {
	fmt.Println("init dao record")
	RecordDAO = &Record{}
	RecordDAO.mysql = mysql.RecordMysql
}
