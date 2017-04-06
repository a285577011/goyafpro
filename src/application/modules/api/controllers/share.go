package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"../../../models/form/share"
	"git.oschina.net/pbaapp/goyaf"
)

type Share struct {
	controllers.Base
}

func (this *Share) IndexAction() {
	this.GetResponse().AppendBody("api share index")
}

//增加分享
func (this *Share) AddAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formShare.NewAdd(this.GetRequest().GetPosts()))
	params["user_id"] = userId

	result := business.ShareBN.Add(params)
	this.PrintSuccessMessage("", result)
}

//分享的详情
func (this *Share) ContentAction() {
	params := this.ValidateForm(formShare.NewContent(this.GetRequest().GetQuerys()))

	result := business.ShareBN.Content(params)
	this.PrintSuccessMessage("", result)
}

func init() {
	goyaf.Debug("init api controller share")
}
