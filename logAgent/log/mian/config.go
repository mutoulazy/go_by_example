package main

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"../tailf"
)

var (
	appConfig *Config
)

type Config struct {
	logLevel string
	logPath string
	chanSize int 
	collectConfs []tailf.CollectConf
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
	appConfig.chanSize,err = conf.Int("logs::chan_size")
	if err != nil {
		appConfig.chanSize = 100
	}

	err = loadCollectConfig(conf)
	if err != nil {
		fmt.Println("new config err, err:", err)
		return
	}

	return
}

func  loadCollectConfig(conf config.Configer) (err error) {
	var collectConf tailf.CollectConf
	collectConf.LogPath = conf.String("collect::log_path")
	if len(collectConf.LogPath) == 0 {
		err = errors.New("load collect::log_path err")
		collectConf.LogPath = "D:\\tools\\logs\\nginx\\test1.log"
		return
	}
	collectConf.Topic = conf.String("collect::topic")
	if len(collectConf.Topic) == 0 {
		err = errors.New("load collect::topic err")
		collectConf.Topic = "nginx_log"
		return
	}
	appConfig.collectConfs = append(appConfig.collectConfs, collectConf)
	return
}