package business

import (
	"../dao"
	//"../error"
	"../http/app"
	"../mysql"
	"encoding/json"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"strconv"
	"time"
)

type Note struct {
	Base
}

//设置信息
func (this *Note) Add(params map[string]string) int64 {
	_, ok := params["add_time"]
	if !ok {
		params["add_time"] = strconv.FormatInt(time.Now().Unix(), 10)
	}
	data := map[string]string{
		"user_id":      params["user_id"],
		"note_content": params["note_content"],
		"add_time":     params["add_time"],
	}
	noteId := dao.NoteDAO.Insert(data)

	go this.sendDynamic(map[string]string{
		"is_share": params["is_share"],
		"user_id":  params["user_id"],
		"note_id":  strconv.FormatInt(noteId, 10),
	})

	return noteId
}

//发布心情成功后，检测是否发布动态
func (this *Note) sendDynamic(params map[string]string) {
	if params["is_share"] != "1" {
		return
	}
	note := mysql.NoteMysql.Find(params["note_id"])
	var content string = note["note_content"]
	contentJson, _ := json.Marshal(map[string]string{
		"mood": content,
	})

	dh := &ourHttpUser.Dynamic{}
	dh.Add(map[string]string{
		"member_id":       params["user_id"],
		"type":            "15",
		"source_id":       params["note_id"],
		"dynamic_content": string(contentJson),
	})
}

//列表
func (this *Note) List(params map[string]string) []map[string]string {
	this.CheckPageAndCount(&params)
	offset, _ := strconv.Atoi(params["offset"])
	count, _ := strconv.Atoi(params["count"])

	where := map[string]string{
		"user_id": params["user_id"],
	}
	if _, ok := params["start_time"]; ok && len(params["start_time"]) > 0 {
		where["add_time>=?"] = params["start_time"]
	}
	if _, ok := params["end_time"]; ok && len(params["end_time"]) > 0 {
		where["add_time<=?"] = params["end_time"]
	}

	result := dao.NoteDAO.FetchAll(db.Select{
		Where:  where,
		Count:  count,
		Offset: offset,
		Order:  "add_time desc",
	})
	return result
}

var NoteBN *Note

func init() {
	goyaf.Log("init business note")
	NoteBN = &Note{}
}
