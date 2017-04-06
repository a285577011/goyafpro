package ourHttpUser

import (
	"git.oschina.net/pbaapp/goyaf"
	"reflect"
)

type User struct {
	Base
}

//解析sso
func (this *User) AnalyzeSso(sso string) (userId string, err error) {
	//userId = "1"
	//return
	this.Uri = "/iapi/member/analyzesso/"
	this.Cookies = map[string]string{"sso": sso}

	var data interface{}
	data, err = this.Base.Request()
	if err != nil {
		return
	}

	v := reflect.ValueOf(data)
	i := v.Interface()
	a := i.(map[string]interface{})

	userId = reflect.ValueOf(a["member_id"]).String()

	return
}

func (this *User) CheckPassword(params map[string]string) (userId string, err error) {
	this.Uri = "/iapi/member/checkpassword/"
	this.Method = "POST"
	this.Params = map[string]string{
		"mobile":   params["pba_account"],
		"password": params["pba_password"],
	}

	var data interface{}
	data, err = this.Base.Request()
	if err != nil {
		return
	}
	userId = reflect.ValueOf(data).String()
	return
}

func NewUser() *User {
	user := &User{}
	return user
}

var UserHttp *User

func init() {
	goyaf.Debug("init http app user")
	UserHttp = &User{}
}
