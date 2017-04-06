package iapicontrollers

import (
	"../../../controllers"
	"../../../models/business/iapi"
	"../../../models/form/iapi/record"
	"git.oschina.net/pbaapp/goyaf"
)

type Record struct {
	controllers.Base
}

//用户数据列表
func (this *Record) ListAction() {
	params := this.ValidateForm(formIapiRecord.NewList(this.GetRequest().GetQuerys()))

	result := businessIapi.RecordBN.List(params)
	this.PrintSuccessMessage("", result)
}

func init() {
	goyaf.Log("init api controller record")
}
