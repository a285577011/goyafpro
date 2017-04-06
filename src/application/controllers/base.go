//控制器基类
package controllers

import (
	"../models/dao"
	"../models/error"
	"../models/http/user"
	"encoding/json"
	"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/form"
	"reflect"
	"strconv"
)

type Base struct {
	goyaf.GoyafController
}

//输出成功信息
func (this *Base) PrintSuccessMessage(message, data interface{}) {
	//转换bool，int为字符串，统一输出
	switch reflect.TypeOf(data).Kind() {
	case reflect.Bool:
		if reflect.ValueOf(data).Bool() {
			data = "1"
		} else {
			data = "0"
		}
	case reflect.Int64:
		data = strconv.FormatInt(reflect.ValueOf(data).Int(), 10)
	case reflect.Float64:
		data = strconv.FormatFloat(reflect.ValueOf(data).Float(), 'f', -1, 64)
	case reflect.Slice:
		//如果是slice类型并且length为0，json编码后会为null，这里做一个转换，使其输出空数组
		if reflect.ValueOf(data).Len() == 0 {
			data = make([]string, 0)
		}
	}

	result := map[string]interface{}{
		"errno":  "0",
		"errmsg": "",
		"data":   data,
	}

	jsonResult, _ := json.Marshal(result)

	isJsonp := this.GetRequest().GetQuery("is_jsonp")
	if len(isJsonp) == 0 {
		this.GetResponse().SetHeader("Content-Type", "application/json; charset=utf-8")
		this.GetResponse().AppendBody(string(jsonResult))
	} else {
		callback := this.GetRequest().GetQuery("callback")
		this.GetResponse().AppendBody(callback + "(" + string(jsonResult) + ")")
	}
}

//输出错误信息
func (this *Base) PrintErrorMessage(err ourError.Error) {
	result := make(map[string]interface{})
	result["errno"] = err.Errno
	result["errmsg"] = err.Errmsg
	result["data"] = ""

	jsonResult, _ := json.Marshal(result)

	isJsonp := this.GetRequest().GetQuery("is_jsonp")
	if len(isJsonp) == 0 {
		this.GetResponse().SetHeader("Content-Type", "application/json; charset=utf-8")
		this.GetResponse().AppendBody(string(jsonResult))
	} else {
		callback := this.GetRequest().GetQuery("callback")
		this.GetResponse().AppendBody(callback + "(" + string(jsonResult) + ")")
	}
}

//校验表单
func (this *Base) ValidateForm(f form.Form) (params map[string]string) {
	if !f.Validate() {
		for k, v := range f.GetFieldsErrmsg() {
			ourError.PanicError(ourError.FormValidateError, "表单错误，"+k+": "+v)
		}
	}
	params = f.GetFieldsFirstValue()
	return
}

//校验是否登录
func (this *Base) ValidateLogin() (userId string) {
	sso := this.GetRequest().GetCookie("sso")
	if len(sso) == 0 {
		sso = this.GetRequest().GetQuery("sso")
	}

	userHttp := ourHttpUser.NewUser()

	userId, err := userHttp.AnalyzeSso(sso)
	if err != nil {
		fmt.Println("请求用户信息失败，sso：", sso)
		ourError.PanicError(ourError.RequestAppError, err.Error())
	}
	return
}

//校验people_id和user_id是否匹配
func (this *Base) ValiatePeopleIdAndUserId(userId string, peopleId string) {
	people := dao.PeopleDAO.Find(peopleId)
	if people["user_id"] != userId {
		fmt.Println(people, userId)
		ourError.PanicError(ourError.PeopleNotBelongUser)
	}
}

func init() {
	fmt.Println("init controller base")
}
