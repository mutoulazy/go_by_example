package main

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"../tailf"
)

func main() {
	// 读取化配置文件
	filename := "D:\\tools\\logs\\logcollect.conf"
	err := loadConfig("ini", filename)
	if err != nil {
		fmt.Println("load config faild, err: %v", err)
		panic("load config faild")
		return
	}

	// 初始化logs
	err = initLogger()
	if err != nil {
		fmt.Println("init logger faild, err: %v", err)
		panic("init logger faild")
		return
	}

	logs.Debug("load config success, conf: %v", appConfig)
	logs.Debug("init success")

	// 初始化tailf
	err = tailf.InitTailf(appConfig.collectConfs, appConfig.chanSize)
	if err != nil {
		logs.Warn("init tailf faild, err: %v", err)
		panic("init tailf faild")
		return
	}

	// 运行服务
	err = runServer()
	if err != nil {
		logs.Warn("run server faild, err: %v", err)
		panic("run server faild")
		return
	}
}