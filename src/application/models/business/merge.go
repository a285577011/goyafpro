//合并用户的测试和心情数据
package business

import (
	"../error"
	"../mysql"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"strings"
)

type Merge struct {
	Base
}

//设置信息
func (this *Merge) Merge(params map[string]string) int64 {
	where := map[string]string{
		"people_id=?": params["from_people_id"],
	}
	result := mysql.RecordMysql.FetchAll(db.Select{
		Where: where,
	})

	tx, _ := mysql.RecordMysql.GetAdapter().Begin()
	if len(result) > 0 {
		sql := "insert into `record` (`people_id`,`people_height`,`record_weight`,`record_bmi`,`record_bmr`,`record_water`,`record_fat`,`record_age`,`record_muscle`,`record_bone`,`add_time`,`day_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + params["to_people_id"] + "'," + "'" + v["people_height"] + "'," + "'" + v["record_weight"] + "'," + "'" + v["record_bmi"] + "'," + "'" + v["record_bmr"] + "'," + "'" + v["record_water"] + "'," + "'" + v["record_fat"] + "'," + "'" + v["record_age"] + "'," + "'" + v["record_muscle"] + "'," + "'" + v["record_bone"] + "'," + "'" + v["add_time"] + "'," + "'" + v["day_time"] + "'" + "),"
		}
		sql = strings.TrimRight(sql, ",")
		goyaf.Debug(sql)
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			goyaf.Debug(err)
			ourError.PanicError(ourError.MigrateRecordError)
		}
	}
	tx.Commit()

	return int64(len(result))
}

var MergeBN *Merge

func init() {
	goyaf.Log("init business merge")
	MergeBN = &Merge{}
}
