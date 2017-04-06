package business

import (
	"../dao"
	"../error"
	"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"git.oschina.net/pbaapp/goyaf/lib"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Record struct {
	Base
}

//增加记录
func (this *Record) Add(params map[string]string, httpLang string) string {
	now := time.Now()
	_, ok := params["add_time"]
	if !ok {
		params["add_time"] = strconv.FormatInt(now.Unix(), 10)
	}
	//计算当天凌晨0点的时间戳
	t, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02")+" 00:00:00")
	params["day_time"] = strconv.FormatInt(t.Unix()-8*3600, 10)

	//因为用户身高会改变，所以在保持记录的时候需要存储用户身高
	people := dao.PeopleDAO.Find(params["people_id"])
	params["people_height"] = people["people_height"]
	if len(params["people_height"]) == 0 {
		params["people_height"] = "0"
	}

	recordId := dao.RecordDAO.Insert(params)
	//检测目标是否达成
	TargetBN.CheckTarget(map[string]string{
		"record_id": strconv.FormatInt(recordId, 10),
		"people_id": params["people_id"],
	})

	//返回打败了多少人
	percent := lib.Round(this.calcBeatPercent(strconv.FormatInt(recordId, 10)), 0)
        if httpLang == "zh-tw" {
            return "打敗全國" + strconv.FormatFloat(percent, 'f', -1, 64) + "%的人"
        } else if httpLang == "en" {
            return "You have defeated the country " + strconv.FormatFloat(percent, 'f', -1, 64) + "% of people"
        } else {
            return "打败全国" + strconv.FormatFloat(percent, 'f', -1, 64) + "%的人"
        }
}

//数据列表
func (this *Record) List(params map[string]string) []map[string]string {
	this.CheckPageAndCount(&params)
	offset, _ := strconv.Atoi(params["offset"])
	count, _ := strconv.Atoi(params["count"])

	where := map[string]string{
		"people_id=?": params["people_id"],
	}
	if _, ok := params["start_time"]; ok && len(params["start_time"]) > 0 {
		where["add_time>=?"] = params["start_time"]
	}
	if _, ok := params["end_time"]; ok && len(params["end_time"]) > 0 {
		where["add_time<=?"] = params["end_time"]
	}

	idResult := dao.RecordDAO.FetchAll(db.Select{
		Columns: "max(record_id) as record_id",
		Where:   where,
		Count:   count,
		Offset:  offset,
		Group:   "day_time",
		Order:   "record_id desc",
	})

	result := make([]map[string]string, len(idResult))
	i := 0
	for _, v := range idResult {
		record := dao.RecordDAO.Find(v["record_id"])
		result[i] = record
		i++
	}

	return result
}

//记录分析
func (this *Record) Analyse(params map[string]string, httpLang string) string {
	//如果type没有值则为1
	//1表示详细描述
	//2表示简短描述
	_, ok := params["type"]
	if !ok || len(params["type"]) == 0 {
		params["type"] = "1"
	}

	people := dao.PeopleDAO.Find(params["people_id"])
	if len(people) == 0 {
		ourError.PanicError(ourError.PeopleIsNotExists)
	}

	age := this.calcAge(people["people_birthday"])

	analyseConfig := dao.AnalyseConfigDAO.FetchAll(db.Select{
		Where: map[string]string{
			"sex=?":                     people["people_sex"],
			"min_age<=? and max_age>=?": strconv.Itoa(age),
			"min_fat<=? and max_fat>=?": params["record_fat"],
			"min_bmi<=? and max_bmi>=?": params["record_bmi"],
		},
		Count: 1,
	})
	if !(len(analyseConfig) > 0) {
            if httpLang == "zh-tw" {
		return "暫無分析報告"
            } else if httpLang == "en" {
		return "No Reports"
            } else {
		return "暂无分析报告"
            }        
	}

	tipsKey := ""
        if params["type"] == "1" {
            if httpLang == "zh-tw" {
                tipsKey = "detail_tips_tw"
            } else if httpLang == "en" {
                tipsKey = "detail_tips_en"
            } else {
                tipsKey = "detail_tips"
            }
        } else {
            tipsKey = "simple_tips"
        }
        
	configTips := strings.Split(strings.Trim(analyseConfig[0][tipsKey], "\r\n"), "\r\n\r\n")
	if !(len(configTips) > 0) {
            if httpLang == "zh-tw" {
		return "暫無分析報告"
            } else if httpLang == "en" {
		return "No Reports"
            } else {
		return "暂无分析报告"
            }        
	}

	rand.Seed(time.Now().UnixNano())
	tips := configTips[rand.Intn(len(configTips))]
	return tips
}

//根据生日计算年龄
//生日格式：1990-01-01
func (this *Record) calcAge(birthday string) (age int) {
	t, _ := time.Parse("2006-01-02", birthday)

	birthdayYear := t.Year()
	nowYear := time.Now().Year()

	age = nowYear - birthdayYear

	return
}

//计算打败了多少人
func (this *Record) calcBeatPercent(recordId string) float64 {
	record := dao.RecordDAO.Find(recordId)
	goyaf.Debug(record)
	people := dao.PeopleDAO.Find(record["people_id"])
	goyaf.Debug(people)

	result := dao.RecordDAO.FetchAll(db.Select{
		Columns: "record_id, record_weight",
		Where: map[string]string{
			"people_id in (select people_id from people where people_sex=" + people["people_sex"] + ")": "db_expression:nil",
		},
		Order: "add_time desc",
		Count: 1000,
	})
	num := 0.0
	CurWeight, _ := strconv.ParseFloat(record["record_weight"], 64)
	for _, row := range result {
		recordWeight, _ := strconv.ParseFloat(row["record_weight"], 64)
		if recordWeight >= CurWeight {
			num = num + 1
		}
	}
	total := strconv.Itoa(len(result) + 1)
	totalF, _ := strconv.ParseFloat(total, 64)

	goyaf.Debug(num)
	goyaf.Debug(totalF)
	goyaf.Debug(num / totalF)
	return num * 100.0 / totalF
}

var RecordBN *Record

func init() {
	fmt.Println("init business record")
	RecordBN = &Record{}
}
