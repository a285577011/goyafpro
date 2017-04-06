package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"../../../models/form/people"
	"fmt"
)

type People struct {
	controllers.Base
}

//新增使用者
func (this *People) AddAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(formPeople.NewAdd(this.GetRequest().GetPosts()))
	params["user_id"] = userId

	result := business.PeopleBN.Add(params)
	this.PrintSuccessMessage("", result)
}

//是否激活脂肪秤
func (this *People) IsactivateAction() {
	userId := this.ValidateLogin()
	params := map[string]string{"user_id": userId}

	result := business.PeopleBN.IsActivate(params)
	this.PrintSuccessMessage("", result)
}

//查找使用者用户信息
func (this *People) InfoAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formPeople.NewInfo(this.GetRequest().GetQuerys()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.PeopleBN.Info(params)
	this.PrintSuccessMessage("", result)
}

//查找使用者列表
func (this *People) ListAction() {
	userId := this.ValidateLogin()
	params := map[string]string{"user_id": userId}

	result := business.PeopleBN.List(params)
	this.PrintSuccessMessage("", result)
}

//修改用户
func (this *People) UpdateAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formPeople.NewUpdate(this.GetRequest().GetPosts()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.PeopleBN.Update(params)
	this.PrintSuccessMessage("", result)
}

//删除用户
func (this *People) DeleteAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formPeople.NewInfo(this.GetRequest().GetQuerys()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.PeopleBN.Delete(params)
	this.PrintSuccessMessage("", result)
}

func init() {
	fmt.Println("init api controller people")
}
