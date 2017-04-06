package form

import (
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/form"
)

//表单基类
type Base struct {
	form.Base
}

func init() {
	goyaf.Debug("init form base")
}
