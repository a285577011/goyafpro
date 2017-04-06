package business

import (
	"../dao"
	"../error"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"strconv"
	"time"
)

type People struct {
}

//增加使用者
func (this *People) Add(params map[string]string) int64 {
	_, ok := params["add_time"]
	if !ok {
		params["add_time"] = strconv.FormatInt(time.Now().Unix(), 10)
	}
	params["last_update_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	//如果是主账号，则判断是否已经有了
	if params["people_type"] == "1" {
		result := dao.PeopleDAO.FetchAll(db.Select{
			Where: map[string]string{"people_type": "1", "user_id": params["user_id"]},
			Count: 1,
		})
		if len(result) > 0 {
			ourError.PanicError(ourError.MainPeopleIsExist)
		}
	}

	peopleId := dao.PeopleDAO.Insert(params)
	return peopleId
}

//是否激活
func (this *People) IsActivate(params map[string]string) (peopleId string) {
	result := dao.PeopleDAO.FetchAll(db.Select{
		Where: map[string]string{"people_type": "1", "user_id": params["user_id"]},
		Count: 1,
	})
	if len(result) > 0 {
		peopleId = result[0]["people_id"]
		return
	}
	peopleId = "0"
	return
}

//用户信息
func (this *People) Info(params map[string]string) map[string]string {
	people := dao.PeopleDAO.Find(params["people_id"])
	return people
}

//用户列表
func (this *People) List(params map[string]string) []map[string]string {
	result := dao.PeopleDAO.FetchAll(db.Select{
		Where: map[string]string{
			"user_id":   params["user_id"],
			"is_delete": "0",
		},
	})
	return result
}

//更新使用者信息
func (this *People) Update(params map[string]string) int64 {
	params["last_update_time"] = strconv.FormatInt(time.Now().Unix(), 10)

	where := map[string]string{"people_id=?": params["people_id"]}
	delete(params, "people_id")

	affect := dao.PeopleDAO.Update(params, where)
	return affect
}

//删除用户
func (this *People) Delete(params map[string]string) int64 {
	people := dao.PeopleDAO.Find(params["people_id"])
	if people["is_delete"] == "1" {
		return 1
	}

	data := map[string]string{
		"is_delete":   "1",
		"delete_time": strconv.FormatInt(time.Now().Unix(), 10),
	}
	affect := dao.PeopleDAO.Update(data, map[string]string{"people_id": params["people_id"]})
	return affect
}

var PeopleBN *People

func init() {
	goyaf.Debug("init business people")
	PeopleBN = &People{}
}
