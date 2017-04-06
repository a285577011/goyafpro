package dao

import (
	"../mysql"
	"encoding/json"
	"git.oschina.net/pbaapp/goyaf"
)

type Mission struct {
	Base
}

//分析连续任务的完成条件
func (this *Mission) AnalyseContinueMissionCondition(c string) map[string]string {
	cm := make(map[string]string)
	json.Unmarshal([]byte(c), &cm)
	return cm
}

func (this *Mission) EncodeContinueMissionCondition(c map[string]string) string {
	b, _ := json.Marshal(c)
	return string(b)
}

var MissionDAO *Mission

func init() {
	goyaf.Log("init dao mission")

	MissionDAO = &Mission{}
	MissionDAO.mysql = mysql.MissionMysql
}
