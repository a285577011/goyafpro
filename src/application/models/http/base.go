package ourHttp

import (
	"git.oschina.net/pbaapp/goyaf"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Base struct {
	Url     string
	Method  string
	Params  map[string]string
	Cookies map[string]string
}

func (this *Base) Request() (body string, err error) {
	client := &http.Client{}
	goyaf.Debug(this.Url)
	var req *http.Request

	uv := url.Values{}
	for k, v := range this.Params {
		uv.Add(k, v)
	}
	goyaf.Debug(strings.NewReader(uv.Encode()))
	req, err = http.NewRequest(this.GetMethod(), this.Url, strings.NewReader(uv.Encode()))
	if err != nil {
		return
	}

	if this.GetMethod() == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if len(this.Cookies) > 0 {
		req.Header.Set("Cookie", this.CookiesToString())
	}

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body = string(bodyBytes)

	return
}

//获取请求方法
func (this *Base) GetMethod() (method string) {
	if this.Method == "" {
		this.Method = "GET"
	}
	return this.Method
}

func (this *Base) CookiesToString() string {
	cookieString := ""
	for k, v := range this.Cookies {
		cookieString = cookieString + k + "=" + v + "&"
	}

	return strings.TrimRight(cookieString, "&")
}

// BUG(chenjiebin): #1: 还需要加入超时时间
