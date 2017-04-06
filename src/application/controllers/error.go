//错误处理控制器
package controllers

import (
	"../models/error"
	"fmt"
	"reflect"
)

type Error struct {
	Base
}

func (this *Error) ErrorAction(r interface{}) {
	fmt.Println(r)
	if reflect.TypeOf(r).String() == "ourError.Error" {
                ourErr := ourError.Error{}
		ourErr.Errno = reflect.ValueOf(r).FieldByName("Errno").String()
                
                //根据语言返回对应信息
                httpLang := this.GetRequest().GetHeader("LANG", "zh-cn")
                if httpLang == "zh-tw" {
                    ourErr.Errmsg = reflect.ValueOf(r).FieldByName("ErrmsgTw").String()
                } else if httpLang == "en" {
                    ourErr.Errmsg = reflect.ValueOf(r).FieldByName("ErrmsgEn").String()
                } else {
                    ourErr.Errmsg = reflect.ValueOf(r).FieldByName("Errmsg").String()
                }

		this.PrintErrorMessage(ourErr)
		return
	}
	panic(r)
}

func init() {
	fmt.Println("init controller error")
}
