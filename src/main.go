package main

import (
	"./application/controllers"
	"./application/modules/api/controllers"
	"./application/modules/iapi/controllers"
	"./conf"
	"git.oschina.net/pbaapp/goyaf"
)

func main() {
	goyaf.SetConfig(conf.Config)

	goyaf.AddController("/index/index/", controllers.Index{})

	goyaf.AddController("/api/index/", apicontrollers.Index{})
	goyaf.AddController("/api/record/", apicontrollers.Record{})
	goyaf.AddController("/api/people/", apicontrollers.People{})
	goyaf.AddController("/api/note/", apicontrollers.Note{})
	goyaf.AddController("/api/target/", apicontrollers.Target{})
	goyaf.AddController("/api/mission/", apicontrollers.Mission{})
	goyaf.AddController("/api/share/", apicontrollers.Share{})
	goyaf.AddController("/api/migrate/", apicontrollers.Migrate{})

	goyaf.AddController("/iapi/record/", iapicontrollers.Record{})
	goyaf.AddController("/iapi/note/", iapicontrollers.Note{})

	errorController := controllers.Error{}
	goyaf.SetPanicHandleController(errorController)

	goyaf.Run()
}

func init() {
	goyaf.Log("init main")
}

// BUG(chenjiebin): #1: http请求需要记录json格式解析错误

// BUG(chenjiebin): #2: http请求json转换的时候要验证

// BUG(chenjiebin): #3: 生日有效格式校验

// BUG(chenjiebin): #4: 数据分析记录没有的区间

// BUG(chenjiebin): #5: 表单中获取表单，如果表单不存在，panic的处理
