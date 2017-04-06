package ourError

import (
	"git.oschina.net/pbaapp/goyaf"
)

//错误类型字典
var errors map[string]Error

//定义错误码
const (
	FormValidateError = "10020"
	RequestAppError   = "10030"
	SystemError       = "10040"
	RequestUserError  = "10050"

	MainPeopleIsExist   = "11010"
	PeopleIsNotExists   = "11020"
	PeopleNotBelongUser = "12010"

	NotHaveAnyRecord = "13010"

	HasAcceptMission      = "14010"
	MissoinIsNotExists    = "14020"
	MissionLogIsNotExists = "14030"
	MissionIsNotFinish    = "14040"
	MissionIsFinish       = "14050"
	NotCurrentDayMission  = "14060"

	ShareIsNotExists = "15010"

	MigratePeopleError     = "16001"
	MigrateRecordError     = "16002"
	MigrateTargetError     = "16003"
	MigrateShareError      = "16004"
	MigrateMissionLogError = "16005"
	MigrateNoteError       = "16006"
	MigrateHaveOtherError  = "16007"
	MigrateMushuHaveError  = "16008"
)

//错误示例
type Error struct {
	Errno  string
	Errmsg string
	ErrmsgTw string
	ErrmsgEn string
}

func init() {
	errors = map[string]Error{
		"-1":    Error{Errno: "-1", Errmsg: "未知错误", ErrmsgTw: "未知錯誤", ErrmsgEn: "error"},
		"10020": Error{Errno: "10020", Errmsg: "表单验证失败", ErrmsgTw: "表單驗證失敗", ErrmsgEn: "Form validation fails"},
		"10030": Error{Errno: "10030", Errmsg: "请求app失败", ErrmsgTw: "請求app失敗", ErrmsgEn: "App failed request"},
		"10040": Error{Errno: "10040", Errmsg: "系统发生错误，请联系客服", ErrmsgTw: "系統發生錯誤，請聯繫客服", ErrmsgEn: "System error, pls contact customer service"},
		"10050": Error{Errno: "10050", Errmsg: "系统发生错误，请联系客服", ErrmsgTw: "系統發生錯誤，請聯繫客服", ErrmsgEn: "System error, pls contact customer service"},

		"11010": Error{Errno: "11010", Errmsg: "主使用账户已经存在", ErrmsgTw: "主使用賬戶已經存在", ErrmsgEn: "Use main account already exists"},
		"11020": Error{Errno: "11020", Errmsg: "使用者不存在", ErrmsgTw: "使用者不存在", ErrmsgEn: "User does not exist"},
		"12010": Error{Errno: "12010", Errmsg: "people_id和当前user_id不符合", ErrmsgTw: "people id和當前user_id不符合", ErrmsgEn: "people id and does not meet current user_id"},
		"13010": Error{Errno: "13010", Errmsg: "没有测试记录", ErrmsgTw: "沒有測試記錄", ErrmsgEn: "No test records"},

		"14010": Error{Errno: "14010", Errmsg: "已经接受了任务", ErrmsgTw: "已經接受了任務", ErrmsgEn: "He has accepted the task"},
		"14020": Error{Errno: "14020", Errmsg: "任务不存在", ErrmsgTw: "任務不存在", ErrmsgEn: "Task does not exist"},
		"14030": Error{Errno: "14030", Errmsg: "任务日志不存在", ErrmsgTw: "任務日誌不存在", ErrmsgEn: "Task logs do not exist"},
		"14040": Error{Errno: "14040", Errmsg: "任务还未达成完成条件", ErrmsgTw: "任務還未達成完成條件", ErrmsgEn: "Task has not yet reached the completion condition"},
		"14050": Error{Errno: "14050", Errmsg: "任务已经结束了", ErrmsgTw: "任務已經結束了", ErrmsgEn: "Mission has ended"},
		"14060": Error{Errno: "14060", Errmsg: "不是当天的任务不能完成", ErrmsgTw: "不是當天的任務不能完成", ErrmsgEn: "Not the task of the day can not be completed"},

		"15010": Error{Errno: "15010", Errmsg: "分享不存在", ErrmsgTw: "分享不存在", ErrmsgEn: "Share does not exist"},

		"16001": Error{Errno: "16001", Errmsg: "迁移用户数据失败", ErrmsgTw: "遷移用戶數據失敗", ErrmsgEn: "Failed to migrate user data"},
		"16002": Error{Errno: "16002", Errmsg: "迁移测试数据失败", ErrmsgTw: "遷移測試數據失敗", ErrmsgEn: "Failed to migrate target data"},
		"16003": Error{Errno: "16003", Errmsg: "迁移目标数据失败", ErrmsgTw: "遷移目標數據失敗", ErrmsgEn: "Failed to migrate mood data"},
		"16004": Error{Errno: "16004", Errmsg: "迁移分享数据失败", ErrmsgTw: "遷移分享數據失敗", ErrmsgEn: "Failed to migrate share data"},
		"16005": Error{Errno: "16005", Errmsg: "迁移任务数据失败", ErrmsgTw: "遷移任務數據失敗", ErrmsgEn: "Failed to migrate task data"},
		"16006": Error{Errno: "16006", Errmsg: "迁移心情数据失败", ErrmsgTw: "遷移心情數據失敗", ErrmsgEn: "Failed to migrate mood data"},
		"16007": Error{Errno: "16007", Errmsg: "该PBA的用户数据已经迁移到木薯其他账号了", ErrmsgTw: "的用戶數據已經遷移到木薯其他賬號了", ErrmsgEn: "The PBA user data has been migrated to other accounts of the cassava"},
		"16008": Error{Errno: "16008", Errmsg: "你的木薯已经账号从其他PBA账号迁入数据了", ErrmsgTw: "你的木薯已經賬號從其他PBA賬號遷入數據了", ErrmsgEn: "Your account has been moved cassava data from other accounts of the PBA"},
	}
}

func CheckSystemError(err error) {
	if err != nil {
		goyaf.ErrorLog.Println(err)
		PanicError(SystemError)
	}
}

func PanicError(errno string, errmsg ...string) {
	err, ok := errors[errno]
	if !ok {
		panic(errors["-1"])
	}

	if len(errmsg) > 0 {
		err.Errmsg = errmsg[0]
		err.ErrmsgTw = errmsg[0]
		err.ErrmsgEn = errmsg[0]
	}
	panic(err)
}

//获取错误
func GetError(errno string, errmsg ...string) Error {
	err, ok := errors[errno]
	if !ok {
		return errors["-1"]
	}

	if len(errmsg) > 0 {
		err.Errmsg = errmsg[0]
	}
	return err
}
