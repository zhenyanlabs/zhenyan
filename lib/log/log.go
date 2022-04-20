package log

import (
	"github.com/beego/beego/v2/core/logs"
)

func NewLogger() *logs.BeeLogger {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)
	//log.Debug("this is a debug message")
	return log
}
