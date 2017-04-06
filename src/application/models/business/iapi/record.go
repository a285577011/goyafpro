package businessIapi

import (
	"../../business"
	"../../dao"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"strconv"
)

type Record struct {
	business.Base
}

//数据列表
func (this *Record) List(params map[string]string) []map[string]string {
	this.CheckPageAndCount(&params)
	offset, _ := strconv.Atoi(params["offset"])
	count, _ := strconv.Atoi(params["count"])

	//如果没有传递people_id，则从user_id获取主使用者
	if len(params["people_id"]) == 0 {
		result := dao.PeopleDAO.FetchAll(db.Select{
			Where: map[string]string{
				"user_id":     params["user_id"],
				"people_type": "1",
				"is_delete":   "0",
			},
		})
		if len(result) == 0 {
			return make([]map[string]string, 0)
		}
		params["people_id"] = result[0]["people_id"]
	}

	where := map[string]string{}
	where["people_id=?"] = params["people_id"]

	if _, ok := params["start_time"]; ok && len(params["start_time"]) > 0 {
		where["add_time>=?"] = params["start_time"]
	}
	if _, ok := params["end_time"]; ok && len(params["end_time"]) > 0 {
		where["add_time<=?"] = params["end_time"]
	}

	result := dao.RecordDAO.FetchAll(db.Select{
		Where:  where,
		Count:  count,
		Offset: offset,
		Order:  "add_time desc",
	})
	return result
}

var RecordBN *Record

func init() {
	goyaf.Log("init business iapi record")
	RecordBN = &Record{}
}
