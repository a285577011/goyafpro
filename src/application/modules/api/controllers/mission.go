package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	//"../../../models/form"
	"../../../models/form/mission"
	"git.oschina.net/pbaapp/goyaf"
)

type Mission struct {
	controllers.Base
}

//首页任务列表
func (this *Mission) ListAction() {
	userId := this.ValidateLogin()

	params := map[string]string{}
	params["user_id"] = userId
	params["mission_type"] = "continue"

	result := business.MissionBN.List(params)
	this.PrintSuccessMessage("", result)
}

//接受任务
func (this *Mission) AcceptAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(formMission.NewAccept(this.GetRequest().GetQuerys()))
	params["user_id"] = userId

	result := business.MissionBN.Accept(params)
	this.PrintSuccessMessage("", result)
}

//接到的任务详情
func (this *Mission) LogdetailAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(formMission.NewLogDetail(this.GetRequest().GetQuerys()))
	params["user_id"] = userId

	result := business.MissionBN.LogDetail(params)
	this.PrintSuccessMessage("", result)
}

//完成子任务
func (this *Mission) FinishccmAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(formMission.NewFinishCCM(this.GetRequest().GetQuerys()))
	params["user_id"] = userId

	result := business.MissionBN.FinishCCM(params)
	this.PrintSuccessMessage("", result)
}

//完成任务
func (this *Mission) FinishAction() {
	userId := this.ValidateLogin()

	params := this.ValidateForm(formMission.NewFinish(this.GetRequest().GetQuerys()))
	params["user_id"] = userId

	result := business.MissionBN.Finish(params)
	this.PrintSuccessMessage("", result)
}

func init() {
	goyaf.Log("init api controller mission")
}
