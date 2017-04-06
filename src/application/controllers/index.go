package controllers

import (
	"../models/http/app"
	"../models/mysql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Index struct {
	Base
}

func (this *Index) IndexAction() {
	//ourError.PanicError(ourError.RequestAppError)
	//panic("hehe")
	this.GetResponse().AppendBody("index index index 13")
}

func (this *Index) SetcookieAction() {
	cookie := &http.Cookie{}

	cookie.Path = "/"
	cookie.Name = "sso"
	cookie.Value = "2105%2Fk%2F7KYEBJ49k0cfe6QX2uuiZXOKREx6mhx6Uvw" //65编号的数据

	this.GetResponse().SetCookie(cookie)
}

func (this *Index) TestAction() {
	//测试发送动态
	var params map[string]string
	str := `{"dynamic_content":"{\"target\":\"我要增重3公斤\"}","member_id":"8386725","source_id":"155","type":"25"}`
	fmt.Println(json.Unmarshal([]byte(str), &params))
	fmt.Println(params)
	dh := &ourHttpUser.Dynamic{}
	fmt.Println(dh.Add(params))
}

//插入测试记录数据
func (this *Index) RecordtestAction() {
	rand.Seed(time.Now().UnixNano())
	time := 1415203200

	for i := 0; i < 30; i++ {
		time += 86400
		data := map[string]string{
			"people_id":     "29",
			"people_height": "170",
			"record_weight": strconv.Itoa((500 + rand.Intn(33)) * 10),
			"record_bmi":    "1840",
			"record_bmr":    "139500",
			"record_water":  "6500",
			"record_fat":    strconv.Itoa((100 + rand.Intn(10)) * 10),
			"record_age":    "22",
			"record_muscle": "8560",
			"record_bone":   "460",
			"add_time":      strconv.Itoa(time),
			"day_time":      strconv.Itoa(time),
		}
		mysql.RecordMysql.Insert(data)
	}

}

//当前时间戳
func (this *Index) NowunixAction() {
	fmt.Println(time.Now().Unix())
}

func (this *Index) Test2Action() {
	fmt.Println(this.GetRequest().GetQuerys())
}

func init() {
	fmt.Println("init controller index")
}
