package iapicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"../../../models/form/iapi/note"
	"git.oschina.net/pbaapp/goyaf"
)

type Note struct {
	controllers.Base
}

//用户数据列表
func (this *Note) ListAction() {
	params := this.ValidateForm(formIapiNote.NewList(this.GetRequest().GetQuerys()))

	result := business.NoteBN.List(params)
	this.PrintSuccessMessage("", result)
}

func init() {
	goyaf.Log("init api controller note")
}
