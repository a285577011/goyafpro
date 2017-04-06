package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"../../../models/form"
	"../../../models/form/people"
	"fmt"
)

type Target struct {
	controllers.Base
}

//增加目标
func (this *Target) AddAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formPeople.NewAddTarget(this.GetRequest().GetPosts()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	params["user_id"] = userId

	result := business.TargetBN.AddTarget(params)
	this.PrintSuccessMessage("", result)
}

//获取目标
func (this *Target) GetAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(form.NewCommon(this.GetRequest().GetQuerys(), []string{"people_id"}))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.TargetBN.GetTarget(params)
	this.PrintSuccessMessage("", map[string]map[string]string{"target": result})
}

func init() {
	fmt.Println("init api controller people")
}
