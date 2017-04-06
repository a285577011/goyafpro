package business

import (
	"../dao"
	"../error"
	//"../http/app"
	//"../mysql"
	//"encoding/json"
	"fmt"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"git.oschina.net/pbaapp/goyaf/lib"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Mission struct {
	Base
}

//列表
func (this *Mission) List(params map[string]string) []map[string]interface{} {
	this.CheckPageAndCount(&params)
	offset, _ := strconv.Atoi(params["offset"])
	count, _ := strconv.Atoi(params["count"])

	where := map[string]string{
		"mission_type": params["mission_type"],
	}
	result := dao.MissionDAO.FetchAll(db.Select{
		Where:  where,
		Count:  count,
		Offset: offset,
	})

	data := make([]map[string]interface{}, len(result))
	i := 0
	for _, row := range result {
		tmp := map[string]interface{}{
			"mission_id":    row["mission_id"],
			"mission_name":  row["mission_name"],
			"mission_desc":  row["mission_desc"],
			"link_share_id": row["link_share_id"],
		}
		//计算任务的总进展
		tmp["total_progress"] = strconv.Itoa(this.CalcMissionProgress(row["mission_id"]))

		//检测用户是否有接受这个任务
		userData := this.MissionStatus(map[string]string{
			"user_id":    params["user_id"],
			"mission_id": row["mission_id"],
		})
		tmp["user_data"] = userData

		//如果有接受任务则只返回单个任务
		if len(userData) > 0 {
			data = make([]map[string]interface{}, 1)
			data[0] = tmp
			return data
		}

		data[i] = tmp
		i = i + 1

	}

	return data
}

//检测任务状态
func (this *Mission) MissionStatus(params map[string]string) map[string]string {
	data := make(map[string]string)
	where := map[string]string{
		"user_id":                params["user_id"],
		"mission_id":             params["mission_id"],
		"(status=1 or status=3)": "db_expression:nil",
	}
	result := dao.MissionLogDAO.FetchAll(db.Select{
		Where: where,
		Count: 1,
	})
	if len(result) == 0 {
		return data
	}
	log := result[0]
	data["log_id"] = log["log_id"]
	data["status"] = log["status"]
	//看看是否到结束时间了
	now := lib.GetNowUnix()
	endTime, _ := strconv.ParseInt(log["end_time"], 10, 64)
	if endTime > 0 && now > endTime {
		data["status"] = "3"
	}
	//计算进度
	data["progress"] = strconv.Itoa(this.CalcLogProgress(log["log_id"]))

	return data
}

//计算接受的任务进展情况
func (this *Mission) CalcLogProgress(logId string) int {
	progress := 0
	log := dao.MissionLogDAO.Find(logId)
	cc := dao.MissionDAO.AnalyseContinueMissionCondition(log["current_completion"])
	for _, v := range cc {
		progress = progress + len(strings.Split(strings.Trim(v, ","), ","))
	}
	return progress
}

//计算任务的总进展
func (this *Mission) CalcMissionProgress(missionId string) int {
	progress := 0
	mission := dao.MissionDAO.Find(missionId)
	c := dao.MissionDAO.AnalyseContinueMissionCondition(mission["mission_condition"])
	for _, v := range c {
		progress = progress + len(strings.Split(strings.Trim(v, ","), ","))
	}
	return progress
}

//接受任务
func (this *Mission) Accept(params map[string]string) int64 {
	//检测是否有任务在进行中
	where := map[string]string{
		"user_id":                params["user_id"],
		"(status=1 or status=3)": "db_expression:nil",
	}
	result := dao.MissionLogDAO.FetchAll(db.Select{
		Where: where,
		Count: 1,
	})
	if len(result) > 0 {
		ourError.PanicError(ourError.HasAcceptMission)
	}

	mission := dao.MissionDAO.Find(params["mission_id"])
	if len(mission) == 0 {
		ourError.PanicError(ourError.MissoinIsNotExists)
	}
	//检测任务是否是continue_child
	if mission["mission_type"] == "continue_child" {
		ourError.PanicError(ourError.MissoinIsNotExists)
	}

	//计算任务结束的时间
	endTime := int64(0)
	if mission["mission_type"] == "continue" {
		endTime = lib.GetTomorrowStartTimeUnix() +
			int64(86400*len(dao.MissionDAO.AnalyseContinueMissionCondition(mission["mission_condition"]))) - 1
	}

	rs := dao.MissionLogDAO.Insert(map[string]string{
		"user_id":    params["user_id"],
		"mission_id": params["mission_id"],
		"status":     "1",
		"start_time": strconv.FormatInt(lib.GetTomorrowStartTimeUnix(), 10),
		"end_time":   strconv.FormatInt(endTime, 10),
		"add_time":   strconv.FormatInt(time.Now().Unix(), 10),
	})
	return rs
}

//任务日志详情
func (this *Mission) LogDetail(params map[string]string) map[string]interface{} {
	log := dao.MissionLogDAO.Find(params["log_id"])
	if len(log) == 0 {
		ourError.PanicError(ourError.MissionLogIsNotExists)
	}
	cc := dao.MissionDAO.AnalyseContinueMissionCondition(log["current_completion"])

	mission := dao.MissionDAO.Find(log["mission_id"])
	if len(mission) == 0 {
		ourError.PanicError(ourError.MissoinIsNotExists)
	}
	data := map[string]interface{}{
		"mission_id":   mission["mission_id"],
		"mission_name": mission["mission_name"],
	}
	//计算当前任务到第几天
	day := this.CalcMissonDay(params["log_id"])

	//计算每天的任务
	c := dao.MissionDAO.AnalyseContinueMissionCondition(mission["mission_condition"])
	days := make([][]map[string]interface{}, len(c))
	for d, v := range c {
		missionIds := strings.Split(strings.Trim(v, ","), ",")
		missions := make([]map[string]interface{}, len(missionIds))
		for k, missionId := range missionIds {
			m := dao.MissionDAO.Find(missionId)
			missions[k] = map[string]interface{}{
				"mission_id":    m["mission_id"],
				"mission_name":  m["mission_name"],
				"mission_desc":  m["mission_desc"],
				"link_share_id": m["link_share_id"],
			}

			userData := map[string]string{}
			dint, _ := strconv.Atoi(d)
			switch {
			case day < dint:
				userData["status"] = "5" //表示任务尚未开始
			case day == dint:
				if strings.Index(cc[d], ","+m["mission_id"]+",") != -1 {
					userData["status"] = "4" //表示任务已经完成
				} else {
					userData["status"] = "3" //表示可以完成了
				}
			case day > dint:
				if strings.Index(cc[d], ","+m["mission_id"]+",") != -1 {
					userData["status"] = "4" //表示任务已经完成
				} else {
					userData["status"] = "6" //表示任务已经过期
				}
			}

			missions[k]["user_data"] = userData
		}
		dint, _ := strconv.Atoi(d)
		days[dint-1] = missions
	}
	data["days"] = days
	return data
}

//完成小任务
func (this *Mission) FinishCCM(params map[string]string) int64 {
	log := dao.MissionLogDAO.Find(params["log_id"])
	if len(log) == 0 {
		ourError.PanicError(ourError.MissionLogIsNotExists)
	}
	if strconv.Itoa(this.CalcMissonDay(params["log_id"])) != params["day"] {
		ourError.PanicError(ourError.NotCurrentDayMission)
	}
	//检测是否有这个小任务
	mission := dao.MissionDAO.Find(log["mission_id"])
	c := dao.MissionDAO.AnalyseContinueMissionCondition(mission["mission_condition"])
	if strings.Index(c[params["day"]], ","+params["mission_id"]+",") == -1 {
		ourError.PanicError(ourError.MissoinIsNotExists)
	}
	//检测是否已经完成了这个任务
	cc := dao.MissionDAO.AnalyseContinueMissionCondition(log["current_completion"])
	if strings.Index(cc[params["day"]], ","+params["mission_id"]+",") != -1 {
		return 1
	}

	cc[params["day"]] = this.AppendFinishMission(log["log_id"], params["mission_id"], params["day"])
	data := map[string]string{
		"current_completion": dao.MissionDAO.EncodeContinueMissionCondition(cc),
	}
	//检测任务的条件是否都达成
	diff := lib.ArrayDiff(strings.Split(strings.Trim(c[strconv.Itoa(len(c))], ","), ","),
		strings.Split(strings.Trim(cc[strconv.Itoa(len(c))], ","), ","))
	if len(diff) == 0 {
		data["status"] = "3"
	}

	dao.MissionLogDAO.Update(data, map[string]string{"log_id": log["log_id"]})

	return 1
}

//完成任务
func (this *Mission) Finish(params map[string]string) string {
	log := dao.MissionLogDAO.Find(params["log_id"])
	if len(log) == 0 {
		ourError.PanicError(ourError.MissionLogIsNotExists)
	}
	if log["status"] != "1" && log["status"] != "3" {
		ourError.PanicError(ourError.MissionLogIsNotExists)
	}

	//看看是否到结束时间了
	now := lib.GetNowUnix()
	endTime, _ := strconv.ParseInt(log["end_time"], 10, 64)
	if endTime > 0 && now > endTime {
		dao.MissionLogDAO.Update(map[string]string{
			"status": "3",
		}, map[string]string{"log_id": log["log_id"]})
		log["status"] = "3"
	}

	if log["status"] != "3" {
		ourError.PanicError(ourError.MissionIsNotFinish)
	}
	dao.MissionLogDAO.Update(map[string]string{
		"status": "4",
	}, map[string]string{"log_id": log["log_id"]})

	progress := this.CalcLogProgress(params["log_id"])
	totalProgress := this.CalcMissionProgress(log["mission_id"])
	r := fmt.Sprintf("%.2f", float64(progress)*100/float64(totalProgress))
	tipsConfig := dao.MissionTipsConfigDAO.FetchAll(db.Select{
		Where: map[string]string{
			"mission_id=?":      log["mission_id"],
			"min<=? and max>=?": r,
		},
		Count: 1,
	})
	if len(tipsConfig) == 0 {
		return "完成任务了"
	}
	configTips := strings.Split(strings.Trim(tipsConfig[0]["tips"], "\r\n"), "\r\n\r\n")
	if !(len(configTips) > 0) {
		return "完成任务了"
	}

	rand.Seed(time.Now().UnixNano())
	tips := configTips[rand.Intn(len(configTips))]
	return tips
}

//计算任务到第几天了
func (this *Mission) CalcMissonDay(logId string) int {
	log := dao.MissionLogDAO.Find(logId)
	startTime, _ := strconv.ParseInt(log["start_time"], 10, 64)
	now := lib.GetNowUnix()
	day := int(math.Ceil((float64(now - startTime)) / 86400))
	return day
}

//追加完成的任务
func (this *Mission) AppendFinishMission(logId string, missionId string, day string) string {
	log := dao.MissionLogDAO.Find(logId)
	cc := dao.MissionDAO.AnalyseContinueMissionCondition(log["current_completion"])

	newCcDay := append(lib.ArrayFilter(strings.Split(strings.Trim(cc[day], ","), ",")), missionId)

	tmp := make([]int, len(newCcDay))
	i := 0
	for _, v := range newCcDay {
		tmp[i], _ = strconv.Atoi(v)
		i = i + 1
	}
	tmp2 := make([]string, len(newCcDay))
	i = 0
	for _, v := range tmp {
		tmp2[i] = strconv.Itoa(v)
		i = i + 1
	}
	return "," + strings.Join(tmp2, ",") + ","
}

var MissionBN *Mission

func init() {
	goyaf.Log("init business mission")
	MissionBN = &Mission{}
}
