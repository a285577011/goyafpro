package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"fmt"
)

type Index struct {
	controllers.Base
}

func (this *Index) IndexAction() {
	this.GetResponse().AppendBody("api index index")
}

//首页菜单
func (this *Index) MenuAction() {
	menus := []map[string]string{
		map[string]string{
			"id":    "1",
			"type":  "1",
			"title": "每日测试",
			"desc":  "遇到要我保重的人，就逼他上秤！",
		},
		map[string]string{
			"id":    "2",
			"type":  "1",
			"title": "详细分析",
			"desc":  "我不是strong，只是虚胖！",
		},
		map[string]string{
			"id":    "3",
			"type":  "1",
			"title": "我的任务包",
			"desc":  "完成任务，打倒绿茶婊虏获高富帅，从此走向人生巅峰！",
		},
		map[string]string{
			"id":    "4",
			"type":  "1",
			"title": "数据记录",
			"desc":  "做一次是一夜情，长期来，才有资格说真爱",
		},
		map[string]string{
			"id":    "5",
			"type":  "2",
			"title": "订购",
			"desc":  "",
		},
		map[string]string{
			"id":    "6",
			"type":  "2",
			"title": "讨论",
			"desc":  "",
		},
		map[string]string{
			"id":    "7",
			"type":  "2",
			"title": "帮助",
			"desc":  "",
		},
	}

	this.PrintSuccessMessage("", menus)
}

//合并record
func (this *Index) MergeAction() {
	result := business.MergeBN.Merge(map[string]string{
		"from_people_id": "100089",
		"to_people_id":   "100088",
	})
	this.PrintSuccessMessage("", result)
}

func init() {
	fmt.Println("init api controller record")
}
