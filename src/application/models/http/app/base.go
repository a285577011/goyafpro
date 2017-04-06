package ourHttpUser

import (
	"../../http"
	"encoding/json"
	"errors"
	//"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"reflect"
	"strconv"
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
	if err != nil {
		return
	}

	v := make(map[string]interface{})
	err = json.Unmarshal([]byte(body), &v)
	if err != nil {
		return
	}

	if reflect.ValueOf(v["errno"]).Kind() != reflect.Float64 {
		err = errors.New("http request app.pba.cn: errno error")
		return
	}

	errno := strconv.FormatFloat(reflect.ValueOf(v["errno"]).Float(), 'f', -1, 64)
	if errno != "0" {
		err = errors.New(reflect.ValueOf(v["errmsg"]).String())
		return
	}

	data = v["data"]

	return
}

func (this *Base) GetHost() string {
	if this.Host == "" {
		this.Host = goyaf.GetConfigByKey("url-app.pba.cn")
	}

	return this.Host
}

// BUG(chenjiebin): #1: 需要记录json格式解析错误

// BUG(chenjiebin): #2: json转换的时候要验证
