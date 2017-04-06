//数据迁移控制器
package apicontrollers

import (
	"../../../controllers"
	"../../../models/business/migrate"
	"../../../models/form/migrate"
	"git.oschina.net/pbaapp/goyaf"
)

type Migrate struct {
	controllers.Base
}

func (this *Migrate) IndexAction() {
	this.GetResponse().AppendBody("api migrate index")
}

func (this *Migrate) MigrateAction() {
	params := this.ValidateForm(formMigrate.NewMigrate(this.GetRequest().GetQuerys()))

	result := businessMigrate.MigrateBN.Migrate(params)
	this.PrintSuccessMessage("", result)
}

//数据迁移信息
func (this *Migrate) IsmigrateAction() {
	params := this.ValidateForm(formMigrate.NewIsMigrate(this.GetRequest().GetQuerys()))

	result := businessMigrate.MigrateBN.IsMigrate(params)
	this.PrintSuccessMessage("", result)
}

//木薯迁移信息
func (this *Migrate) MigrateinfoAction() {
	params := this.ValidateForm(formMigrate.NewIsMigrate(this.GetRequest().GetQuerys()))

	result := businessMigrate.MigrateBN.MigrateInfo(params)
	this.PrintSuccessMessage("", result)
}

func init() {
	goyaf.Debug("init api controller migrate")
}
