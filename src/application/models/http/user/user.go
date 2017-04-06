package ourHttpUser

import (
	"encoding/json"
	//"errors"
	"crypto/md5"
	"encoding/hex"
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
	this.Uri = "/iapi/user/analyzesso/"
	this.Cookies = map[string]string{"sso": sso}

	var data interface{}
	data, err = this.Base.Request()
	if err != nil {
		return
	}

	v := reflect.ValueOf(data)
	i := v.Interface()
	a := i.(map[string]interface{})

	userId = reflect.ValueOf(a["uid"]).String()

	return
}

//解析sso
func (this *User) GetUserInfo(userId string, fields string) (userInfo map[string]string, err error) {
	this.Uri = "/api/my/getinfo/?user_id=" + userId + "&fields=" + fields

	var data interface{}
	data, err = this.Base.Request()
	if err != nil {
		return
	}

	var b []byte
	b, err = json.Marshal(data)
	if err != nil {
		return
	}

	userInfo = make(map[string]string)
	err = json.Unmarshal(b, &userInfo)
	if err != nil {
		return
	}

	return
}

//校验木薯账号和密码，返回用户编号
func (this *User) CheckPassword(params map[string]string) (userId string, err error) {
	h := md5.New()
	h.Write([]byte(params["password"]))
	this.Uri = "/iapi/user/ilogin/?mobile=" + params["mobile"] + "&password=" + hex.EncodeToString(h.Sum(nil))
	this.Method = "GET"

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

func init() {
	goyaf.Debug("init http user user")
}
