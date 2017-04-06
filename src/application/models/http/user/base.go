package ourHttpUser

import (
	"../../http"
	"encoding/json"
	"errors"
	//"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"reflect"
)

type Base struct {
	Host string //主机地址
	Uri  string //请求的uri
	ourHttp.Base
}

func (this *Base) Request() (data interface{}, err error) {
	this.Url = this.GetHost() + this.Uri
	var body string
	body, err = this.Base.Request()
	goyaf.Debug("body:", body)

	if err != nil {
		return
	}

	v := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &v)
	if err != nil {
		return
	}

	if reflect.ValueOf(v["errno"]).Kind() != reflect.String {
		err = errors.New("http request user.mushu.cn: errno error")
		return
	}

	if reflect.ValueOf(v["errno"]).String() != "0" {
		err = errors.New(reflect.ValueOf(v["errmsg"]).String())
		return
	}

	data = v["data"]

	return
}

func (this *Base) GetHost() string {
	if this.Host == "" {
		this.Host = goyaf.GetConfigByKey("url-user.mushu.cn")
	}

	return this.Host
}
