package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath string
	collectConfs []CollectConf
}

type CollectConf struct {
	logPath string
	topic string
}

func  loadConfig(confTyep string, filename string) (err error) {
	conf, err := config.NewConfig(confTyep, filename)
	if err != nil {
		fmt.Println("new config err, err:", err)
		return
	}
	appConfig = &Config{}
	appConfig.logLevel = conf.String("logs::log_level")
	if len(appConfig.logLevel) == 0 {
		appConfig.logLevel = "debug"
	}
	appConfig.logPath = conf.String("logs::log_path")
	if len(appConfig.logPath) == 0 {
		appConfig.logPath = "D:\\tools\\logs\\logcollect.conf"
	}

	err = loadCollectConfig(conf)
	if err != nil {
		fmt.Println("new config err, err:", err)
		return
	}

	return
}

func  loadCollectConfig(conf config.Configer) (err error) {
	var collectConf CollectConf
	collectConf.logPath = conf.String("collect::log_path")
	if len(collectConf.logPath) == 0 {
		err = errors.New("load collect::log_path err")
		collectConf.logPath = "D:\\tools\\logs\\nginx\\test1.log"
		return
	}
	collectConf.topic = conf.String("collect::topic")
	if len(collectConf.topic) == 0 {
		err = errors.New("load collect::topic err")
		collectConf.topic = "nginx_log"
		return
	}
	appConfig.collectConfs = append(appConfig.collectConfs, collectConf)
	return
}