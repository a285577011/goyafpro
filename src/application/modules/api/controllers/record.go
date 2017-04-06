package apicontrollers

import (
	"../../../controllers"
	"../../../models/business"
	"../../../models/form/record"
	"git.oschina.net/pbaapp/goyaf"
)

type Record struct {
	controllers.Base
}

//增加记录
func (this *Record) AddAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formRecord.NewAdd(this.GetRequest().GetPosts()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.RecordBN.Add(params, this.GetRequest().GetHeader("LANG", "zh-cn"))
	this.PrintSuccessMessage("", result)
}

//列表
func (this *Record) ListAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formRecord.NewList(this.GetRequest().GetQuerys()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.RecordBN.List(params)
	this.PrintSuccessMessage("", result)
}

//记录分析
func (this *Record) AnalyseAction() {
	userId := this.ValidateLogin()
	params := this.ValidateForm(formRecord.NewAnalyse(this.GetRequest().GetPosts()))
	this.ValiatePeopleIdAndUserId(userId, params["people_id"])

	result := business.RecordBN.Analyse(params, this.GetRequest().GetHeader("LANG", "zh-cn"))
	this.PrintSuccessMessage("", result)
}

//各项参数说明
func (this *Record) ParamsAction() {
	params := []map[string]string{
		map[string]string{
		},
	}
        httpLang := this.GetRequest().GetHeader("LANG", "zh-cn")
        if httpLang == "zh-tw" {
            params = []map[string]string{
                    map[string]string{
                            "id":    "1",
                            "title": "脂肪率",
                            "desc":  "脂肪率是指身体成分中，脂肪组织所占的比率。测量脂肪率比单纯的只测量体重更能反映我们身体的脂肪水平（肥胖程度）。所以减肥≠减重，减肥应该减脂(tw)。",
                    },
                    map[string]string{
                            "id":    "2",
                            "title": "基础代谢率",
                            "desc":  "人体在清醒而又极端安静的状态下，不受肌肉活动、环境温度、食物及精神紧张等影响时的能量代谢率。代谢率越高意味着你消耗的热量越多，也越不容易发胖。",
                    },
                    map[string]string{
                            "id":    "3",
                            "title": "水含量",
                            "desc":  "人体含水量的百分比。水分充足有利于新陈代谢，水分过多也易于造成水肿。",
                    },
                    map[string]string{
                            "id":    "4",
                            "title": "BMI",
                            "desc":  "即身体质量指数，简称体质指数又称体重指数，等于体重/身高的平方，是世界公认的一种评定肥胖程度的分级方法，也是目前国际上常用的衡量人体胖瘦程度以及是否健康的一个标准，简单、实用、可反映全身性超重和肥胖。",
                    },
                    map[string]string{
                            "id":    "5",
                            "title": "身体年龄",
                            "desc":  "以基础代谢为标准，综合脂肪率、肌肉量等指标，换算得出的真实身体综合状况。 它直接反应出身体的循环状况、血液流通情况以及营养吸收的程度！",
                    },
                    map[string]string{
                            "id":    "6",
                            "title": "肌肉量",
                            "desc":  "根据人体肌肉总量和体重、身高等相结合得到的比例值，决定人体健康状况和力量大小。",
                    },
                    map[string]string{
                            "id":    "7",
                            "title": "骨量",
                            "desc":  "单位体积内，骨组织——骨矿物质（钙、磷等）和骨基质（骨胶原、蛋白质、无机盐等等）含量。",
                    },
            }
        } else if httpLang == "en" {
            params = []map[string]string{
                    map[string]string{
                            "id":    "1",
                            "title": "脂肪率",
                            "desc":  "脂肪率是指身体成分中，脂肪组织所占的比率。测量脂肪率比单纯的只测量体重更能反映我们身体的脂肪水平（肥胖程度）。所以减肥≠减重，减肥应该减脂。",
                    },
                    map[string]string{
                            "id":    "2",
                            "title": "基础代谢率",
                            "desc":  "人体在清醒而又极端安静的状态下，不受肌肉活动、环境温度、食物及精神紧张等影响时的能量代谢率。代谢率越高意味着你消耗的热量越多，也越不容易发胖。",
                    },
                    map[string]string{
                            "id":    "3",
                            "title": "水含量",
                            "desc":  "人体含水量的百分比。水分充足有利于新陈代谢，水分过多也易于造成水肿。",
                    },
                    map[string]string{
                            "id":    "4",
                            "title": "BMI",
                            "desc":  "即身体质量指数，简称体质指数又称体重指数，等于体重/身高的平方，是世界公认的一种评定肥胖程度的分级方法，也是目前国际上常用的衡量人体胖瘦程度以及是否健康的一个标准，简单、实用、可反映全身性超重和肥胖。",
                    },
                    map[string]string{
                            "id":    "5",
                            "title": "身体年龄",
                            "desc":  "以基础代谢为标准，综合脂肪率、肌肉量等指标，换算得出的真实身体综合状况。 它直接反应出身体的循环状况、血液流通情况以及营养吸收的程度！",
                    },
                    map[string]string{
                            "id":    "6",
                            "title": "肌肉量",
                            "desc":  "根据人体肌肉总量和体重、身高等相结合得到的比例值，决定人体健康状况和力量大小。",
                    },
                    map[string]string{
                            "id":    "7",
                            "title": "骨量",
                            "desc":  "单位体积内，骨组织——骨矿物质（钙、磷等）和骨基质（骨胶原、蛋白质、无机盐等等）含量。",
                    },
            }
        } else {
            params = []map[string]string{
                    map[string]string{
                            "id":    "1",
                            "title": "脂肪率",
                            "desc":  "脂肪率是指身体成分中，脂肪组织所占的比率。测量脂肪率比单纯的只测量体重更能反映我们身体的脂肪水平（肥胖程度）。所以减肥≠减重，减肥应该减脂。",
                    },
                    map[string]string{
                            "id":    "2",
                            "title": "基础代谢率",
                            "desc":  "人体在清醒而又极端安静的状态下，不受肌肉活动、环境温度、食物及精神紧张等影响时的能量代谢率。代谢率越高意味着你消耗的热量越多，也越不容易发胖。",
                    },
                    map[string]string{
                            "id":    "3",
                            "title": "水含量",
                            "desc":  "人体含水量的百分比。水分充足有利于新陈代谢，水分过多也易于造成水肿。",
                    },
                    map[string]string{
                            "id":    "4",
                            "title": "BMI",
                            "desc":  "即身体质量指数，简称体质指数又称体重指数，等于体重/身高的平方，是世界公认的一种评定肥胖程度的分级方法，也是目前国际上常用的衡量人体胖瘦程度以及是否健康的一个标准，简单、实用、可反映全身性超重和肥胖。",
                    },
                    map[string]string{
                            "id":    "5",
                            "title": "身体年龄",
                            "desc":  "以基础代谢为标准，综合脂肪率、肌肉量等指标，换算得出的真实身体综合状况。 它直接反应出身体的循环状况、血液流通情况以及营养吸收的程度！",
                    },
                    map[string]string{
                            "id":    "6",
                            "title": "肌肉量",
                            "desc":  "根据人体肌肉总量和体重、身高等相结合得到的比例值，决定人体健康状况和力量大小。",
                    },
                    map[string]string{
                            "id":    "7",
                            "title": "骨量",
                            "desc":  "单位体积内，骨组织——骨矿物质（钙、磷等）和骨基质（骨胶原、蛋白质、无机盐等等）含量。",
                    },
            }
        }

	this.PrintSuccessMessage("", params)
}

func init() {
	goyaf.Debug("init api controller record")
}
