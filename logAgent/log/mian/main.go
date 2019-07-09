package main

import (
	"github.com/astaxie/beego/logs"
	"fmt"
	"../tailf"
	"../kafka"
)

/**
启动zookeeper 		bin/zkServer.cmd
启动kafka	  		./bin/windows/kafka-server-start.bat ./config/server.preperties
启动kafkaClient		kafka-console-consumer.bat  --bootstrap-server localhost:9092 --topic nginx_log --from-beginning
*/
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

	// 初始化kafka
	err = kafka.InitKafka(appConfig.kafkaAddr)
	if err != nil {
		logs.Warn("init kafka faild, err: %v", err)
		panic("init kafka faild")
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