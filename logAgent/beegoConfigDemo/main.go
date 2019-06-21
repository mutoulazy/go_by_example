package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(10000)
	log.SetLogger("console", "")
	log.Trace("trace")
	log.Info("info")
	log.Warn("warning")
	log.Debug("debug")
	log.Critical("critical")
}