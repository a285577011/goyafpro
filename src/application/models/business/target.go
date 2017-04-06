package business

import (
	"../dao"
	"../error"
	"../http/app"
	"../mysql"
	"encoding/json"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"math"
	"strconv"
	"time"
)

type Target struct {
	DynamicShareConfig []map[string]string
}

//检测目标是否达成
func (this *Target) CheckTarget(params map[string]string) (isFinish bool) {
	target := this.GetTarget(map[string]string{"people_id": params["people_id"]})
	if len(target) == 0 {
		return
	}
	record := dao.RecordDAO.Find(params["record_id"])
	if len(record) == 0 {
		return
	}

	isFinish = false
	//增重
	if target["target_type"] == "1" &&
		record["record_weight"] >= target["target_weight"] {
		isFinish = true
	}
	//减肥
	if target["target_type"] == "2" &&
		record["record_weight"] <= target["target_weight"] {
		isFinish = true
	}

	if isFinish {
		dao.TargetDAO.Update(map[string]string{
			"is_finish":   "1",
			"finish_time": strconv.FormatInt(time.Now().Unix(), 10),
		}, map[string]string{"target_id": target["target_id"]})
	}
	return
}

//增加目标
func (this *Target) AddTarget(params map[string]string) int64 {
	_, ok := params["add_time"]
	if !ok {
		params["add_time"] = strconv.FormatInt(time.Now().Unix(), 10)
	}

	//获取上一条测试记录，检测是增肥还是减重
	result := dao.RecordDAO.FetchAll(db.Select{
		Where: map[string]string{
			"people_id=?": params["people_id"],
		},
		Count: 1,
		Order: "add_time desc",
	})
	if len(result) == 0 {
		ourError.PanicError(ourError.NotHaveAnyRecord)
	}

	data := map[string]string{
		"target_weight": params["target_weight"],
		"people_id":     params["people_id"],
		"add_time":      params["add_time"],
	}
	if result[0]["record_weight"] > params["target_weight"] {
		data["target_type"] = "2"
	} else {
		data["target_type"] = "1"
	}

	targetId := dao.TargetDAO.Insert(data)

	go this.sendDynamic(map[string]string{
		"is_share":  params["is_share"],
		"record_id": result[0]["record_id"],
		"user_id":   params["user_id"],
		"target_id": strconv.FormatInt(targetId, 10),
	})

	return targetId
}

//发布目标成功后，检测是否发布动态
func (this *Target) sendDynamic(params map[string]string) {
	if params["is_share"] != "1" {
		return
	}
	record := mysql.RecordMysql.Find(params["record_id"])
	target := mysql.TargetMysql.Find(params["target_id"])

	var content string
	rw, _ := strconv.Atoi(record["record_weight"])
	tw, _ := strconv.Atoi(target["target_weight"])
	abs := math.Abs(float64((tw*100)-rw) / float64(100))
	for _, v := range this.DynamicShareConfig {
		min, _ := strconv.ParseFloat(v["min"], 64)
		max, _ := strconv.ParseFloat(v["max"], 64)
		if min <= abs && abs <= max {
			content = v["content"]
			break
		}
	}
	goyaf.Debug("abs", abs)
	if len(content) == 0 {
		goyaf.Error("增加目标没有内容页发布到动态", params)
		return
	}
	contentJson, _ := json.Marshal(map[string]string{
		"target": content,
	})

	dh := &ourHttpUser.Dynamic{}
	dh.Add(map[string]string{
		"member_id":       params["user_id"],
		"type":            "25",
		"source_id":       params["target_id"],
		"dynamic_content": string(contentJson),
	})
}

//获取目标
func (this *Target) GetTarget(params map[string]string) map[string]string {
	result := dao.TargetDAO.FetchAll(db.Select{
		Where: map[string]string{"people_id": params["people_id"], "is_finish=?": "0"},
		Order: "add_time desc",
		Count: 1,
	})
	if len(result) > 0 {
		return result[0]
	}

	return make(map[string]string)
}

var TargetBN *Target

func init() {
	goyaf.Debug("init business target")
	TargetBN = &Target{}
	config := make([]map[string]string, 4)
	config[0] = map[string]string{"config_id": "1", "min": "0", "max": "1.9", "content": "维持现有的体型也不错，加油哦。"}
	config[1] = map[string]string{"config_id": "2", "min": "2.0", "max": "4.9", "content": "目标虽小，却不易达成，干巴爹！"}
	config[2] = map[string]string{"config_id": "3", "min": "5.0", "max": "11.9", "content": "这个目标可不太容易达到，给自己定的时间不要太短哦！"}
	config[3] = map[string]string{"config_id": "4", "min": "12.0", "max": "100000", "content": "你确定？一口吃不了胖子，要注意科学健身哦。"}
	TargetBN.DynamicShareConfig = config
}
