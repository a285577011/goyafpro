package businessMigrate

import (
	"../../business"
	"../../error"
	"../../http/user"
	"../../mysql"
	"../../mysql/pbayouxing"
	"git.oschina.net/pbaapp/goyaf"
	"git.oschina.net/pbaapp/goyaf/db"
	"git.oschina.net/pbaapp/goyaf/lib"
	"strings"
)

type Migrate struct {
	business.Base
}

//数据迁移
func (this *Migrate) Migrate(params map[string]string) bool {
	userHttp := ourHttpUser.NewUser()
	mushuUserId, err := userHttp.CheckPassword(map[string]string{
		"mobile":   params["mushu_account"],
		"password": params["mushu_password"],
	})
	if err != nil {
		ourError.PanicError(ourError.RequestUserError, err.Error())
	}
	pbaUserId := params["pba_user_id"]
	params["user_id"] = mushuUserId
	//检测用户是否已经有迁移到了其它账号
	r1 := mysql.MigrateMysql.FetchAll(db.Select{
		Where: map[string]string{
			"pba_user_id": pbaUserId,
		},
	})
	if len(r1) > 0 && r1[0]["user_id"] != params["user_id"] {
		ourError.PanicError(ourError.MigrateHaveOtherError)
	}
	r2 := mysql.MigrateMysql.FetchAll(db.Select{
		Where: map[string]string{
			"user_id": params["user_id"],
		},
	})
	if len(r2) > 0 && r2[0]["pba_user_id"] != pbaUserId {
		ourError.PanicError(ourError.MigrateMushuHaveError)
	}

	migrateLog := map[string]string{}
	//记录用户数据
	result := mysql.MigrateMysql.FetchAll(db.Select{
		Where: map[string]string{
			"pba_user_id": pbaUserId,
			"user_id":     params["user_id"],
		},
	})
	if len(result) == 0 {
		migrateLog = map[string]string{
			"pba_user_id": pbaUserId,
			"user_id":     params["user_id"],
			"process":     "0",
		}
		mysql.MigrateMysql.Insert(migrateLog)
	} else {
		migrateLog = result[0]
	}
	switch migrateLog["process"] {
	case "0":
		this.people(pbaUserId, params["user_id"])
		fallthrough
	case "1":
		this.record(pbaUserId, params["user_id"])
		fallthrough
	case "2":
		this.target(pbaUserId, params["user_id"])
		fallthrough
	case "3":
		this.share(pbaUserId, params["user_id"])
		fallthrough
	case "4":
		this.missionlog(pbaUserId, params["user_id"])
		fallthrough
	case "5":
		this.note(pbaUserId, params["user_id"])
	}

	return true
}

//迁移用户数据
func (this *Migrate) people(pbaUserId string, UserId string) bool {
	result := mysqlPbayouxing.PeopleMysql.FetchAll(db.Select{
		Where: map[string]string{
			"user_id": pbaUserId,
		},
	})
	tx, _ := mysql.PeopleMysql.GetAdapter().Begin()
	if len(result) > 0 {
		//修改木薯这边的主使用者为普通使用者
		sql := "update people set people_type=2 where user_id=" + UserId + " and people_type=1"
		goyaf.Debug(sql)
		_, err3 := tx.Exec(sql)
		if err3 != nil {
			tx.Rollback()
			goyaf.Debug(err3)
			ourError.PanicError(ourError.MigratePeopleError)
		}

		sql = "insert into `people` (`people_id`,`user_id`,`people_name`,`people_avatar`,`people_type`,`people_height`,`people_birthday`,`people_sex`,`add_time`,`last_update_time`,`is_delete`,`delete_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + v["people_id"] + "'," + "'" + UserId + "'," + "'" + v["people_name"] + "'," + "'" + v["people_avatar"] + "'," + "'" + v["people_type"] + "'," + "'" + v["people_height"] + "'," + "'" + v["people_birthday"] + "'," + "'" + v["people_sex"] + "'," + "'" + v["add_time"] + "'," + "'" + v["last_update_time"] + "'," + "'" + v["is_delete"] + "'," + "'" + v["delete_time"] + "'" + "),"
		}
		sql = strings.TrimRight(sql, ",")
		goyaf.Debug(sql)
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			goyaf.Debug(err)
			ourError.PanicError(ourError.MigratePeopleError)
		}
	}

	uSql := "update `migrate` set process=1 where pba_user_id=" + pbaUserId + " and user_id=" + UserId
	goyaf.Debug(uSql)
	_, err2 := tx.Exec(uSql)
	if err2 != nil {
		tx.Rollback()
		goyaf.Debug(err2)
		ourError.PanicError(ourError.MigratePeopleError)
	}
	tx.Commit()

	return true
}

//迁移record数据
func (this *Migrate) record(pbaUserId string, UserId string) bool {
	peopleResult := mysql.PeopleMysql.FetchAll(db.Select{
		Columns: "people_id",
		Where: map[string]string{
			"user_id": UserId,
		},
	})

	result := mysqlPbayouxing.RecordMysql.FetchAll(db.Select{
		Where: map[string]string{
			"people_id in (?)": strings.Join(lib.MapColumn(peopleResult, "people_id"), ","),
		},
	})

	tx, _ := mysql.RecordMysql.GetAdapter().Begin()
	if len(result) > 0 {
		sql := "insert into `record` (`record_id`,`people_id`,`people_height`,`record_weight`,`record_bmi`,`record_bmr`,`record_water`,`record_fat`,`record_age`,`record_muscle`,`record_bone`,`add_time`,`day_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + v["record_id"] + "'," + "'" + v["people_id"] + "'," + "'" + v["people_height"] + "'," + "'" + v["record_weight"] + "'," + "'" + v["record_bmi"] + "'," + "'" + v["record_bmr"] + "'," + "'" + v["record_water"] + "'," + "'" + v["record_fat"] + "'," + "'" + v["record_age"] + "'," + "'" + v["record_muscle"] + "'," + "'" + v["record_bone"] + "'," + "'" + v["add_time"] + "'," + "'" + v["day_time"] + "'" + "),"
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

	uSql := "update `migrate` set process=2 where pba_user_id=" + pbaUserId + " and user_id=" + UserId
	goyaf.Debug(uSql)
	_, err2 := tx.Exec(uSql)
	if err2 != nil {
		tx.Rollback()
		goyaf.Debug(err2)
		ourError.PanicError(ourError.MigrateRecordError)
	}
	tx.Commit()

	return true
}

//迁移target
func (this *Migrate) target(pbaUserId string, UserId string) bool {
	peopleResult := mysql.PeopleMysql.FetchAll(db.Select{
		Columns: "people_id",
		Where: map[string]string{
			"user_id": UserId,
		},
	})

	result := mysqlPbayouxing.TargetMysql.FetchAll(db.Select{
		Where: map[string]string{
			"people_id in (?)": strings.Join(lib.MapColumn(peopleResult, "people_id"), ","),
		},
	})

	tx, _ := mysql.TargetMysql.GetAdapter().Begin()
	if len(result) > 0 {
		sql := "insert into `target` (`target_id`,`people_id`,`target_weight`,`target_type`,`is_finish`,`finish_time`,`add_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + v["target_id"] + "'," + "'" + v["people_id"] + "'," + "'" + v["target_weight"] + "'," + "'" + v["target_type"] + "'," + "'" + v["is_finish"] + "'," + "'" + v["finish_time"] + "'," + "'" + v["add_time"] + "'" + "),"
		}
		sql = strings.TrimRight(sql, ",")
		goyaf.Debug(sql)
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			goyaf.Debug(err)
			ourError.PanicError(ourError.MigrateTargetError)
		}
	}

	uSql := "update `migrate` set process=3 where pba_user_id=" + pbaUserId + " and user_id=" + UserId
	goyaf.Debug(uSql)
	_, err2 := tx.Exec(uSql)
	if err2 != nil {
		tx.Rollback()
		goyaf.Debug(err2)
		ourError.PanicError(ourError.MigrateTargetError)
	}
	tx.Commit()

	return true
}

//迁移share
func (this *Migrate) share(pbaUserId string, UserId string) bool {
	result := mysqlPbayouxing.ShareMysql.FetchAll(db.Select{
		Where: map[string]string{
			"user_id": pbaUserId,
		},
	})
	tx, _ := mysql.ShareMysql.GetAdapter().Begin()
	if len(result) > 0 {
		sql := "insert into `share` (`share_id`,`user_id`,`share_type`,`share_content`,`add_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + v["share_id"] + "'," + "'" + UserId + "'," + "'" + v["share_type"] + "'," + "'" + v["share_content"] + "'," + "'" + v["add_time"] + "'" + "),"
		}
		sql = strings.TrimRight(sql, ",")
		goyaf.Debug(sql)
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			goyaf.Debug(err)
			ourError.PanicError(ourError.MigrateShareError)
		}
	}

	uSql := "update `migrate` set process=4 where pba_user_id=" + pbaUserId + " and user_id=" + UserId
	goyaf.Debug(uSql)
	_, err2 := tx.Exec(uSql)
	if err2 != nil {
		tx.Rollback()
		goyaf.Debug(err2)
		ourError.PanicError(ourError.MigrateShareError)
	}
	tx.Commit()

	return true
}

//迁移missionlog
func (this *Migrate) missionlog(pbaUserId string, UserId string) bool {
	result := mysqlPbayouxing.MissionLogMysql.FetchAll(db.Select{
		Where: map[string]string{
			"user_id": pbaUserId,
		},
	})
	tx, _ := mysql.MissionLogMysql.GetAdapter().Begin()
	if len(result) > 0 {
		sql := "insert into `mission_log` (`log_id`,`user_id`,`mission_id`,`current_completion`,`status`,`start_time`,`end_time`,`add_time`,`finish_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + v["log_id"] + "'," + "'" + UserId + "'," + "'" + v["mission_id"] + "'," + "'" + v["current_completion"] + "'," + "'" + v["status"] + "'," + "'" + v["start_time"] + "'," + "'" + v["end_time"] + "'," + "'" + v["add_time"] + "'," + "'" + v["finish_time"] + "'" + "),"
		}
		sql = strings.TrimRight(sql, ",")
		goyaf.Debug(sql)
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			goyaf.Debug(err)
			ourError.PanicError(ourError.MigrateMissionLogError)
		}
	}

	uSql := "update `migrate` set process=5 where pba_user_id=" + pbaUserId + " and user_id=" + UserId
	goyaf.Debug(uSql)
	_, err2 := tx.Exec(uSql)
	if err2 != nil {
		tx.Rollback()
		goyaf.Debug(err2)
		ourError.PanicError(ourError.MigrateMissionLogError)
	}
	tx.Commit()

	return true
}

//迁移note
func (this *Migrate) note(pbaUserId string, UserId string) bool {
	result := mysqlPbayouxing.NoteMysql.FetchAll(db.Select{
		Where: map[string]string{
			"user_id": pbaUserId,
		},
	})
	tx, _ := mysql.NoteMysql.GetAdapter().Begin()
	if len(result) > 0 {
		sql := "insert into `note` (`note_id`,`user_id`,`note_content`,`add_time`) values "
		for _, v := range result {
			sql = sql + "(" + "'" + v["note_id"] + "'," + "'" + UserId + "'," + "'" + v["note_content"] + "'," + "'" + v["add_time"] + "'" + "),"
		}
		sql = strings.TrimRight(sql, ",")
		goyaf.Debug(sql)
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			goyaf.Debug(err)
			ourError.PanicError(ourError.MigrateNoteError)
		}
	}

	uSql := "update `migrate` set process=6 where pba_user_id=" + pbaUserId + " and user_id=" + UserId
	goyaf.Debug(uSql)
	_, err2 := tx.Exec(uSql)
	if err2 != nil {
		tx.Rollback()
		goyaf.Debug(err2)
		ourError.PanicError(ourError.MigrateNoteError)
	}
	tx.Commit()

	return true
}

//用户数据是否迁移
func (this *Migrate) IsMigrate(params map[string]string) bool {
	result := mysql.MigrateMysql.FetchAll(db.Select{
		Where: map[string]string{
			"pba_user_id": params["pba_user_id"],
		},
	})
	if len(result) == 0 {
		return false
	}
	if result[0]["process"] != "0" {
		return true
	}
	return false
}

//获取木薯迁移信息
func (this *Migrate) MigrateInfo(params map[string]string) map[string]string {
	result := mysql.MigrateMysql.FetchAll(db.Select{
		Where: map[string]string{
			"pba_user_id": params["pba_user_id"],
		},
	})
	if len(result) == 0 {
		return map[string]string{}
	}
	return result[0]
}

var MigrateBN *Migrate

func init() {
	goyaf.Debug("init business migrateBN")
	MigrateBN = &Migrate{}
}
