package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"../../../models/form"
	"../../../models/form/note"
)

type Note struct {
	controllers.Base
}

//新增日记
func (this *Note) AddAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(formPeople.NewAdd(this.GetRequest().GetPosts()))
	params["user_id"] = userId

	result := business.NoteBN.Add(params)
	this.PrintSuccessMessage("", result)
}

//日记列表
func (this *Note) ListAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(form.NewCommon(this.GetRequest().GetQuerys(), []string{"page", "count"}))
	params["user_id"] = userId

	result := business.NoteBN.List(params)
	this.PrintSuccessMessage("", result)
}
