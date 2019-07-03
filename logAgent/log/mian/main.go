package main

import (
	"github.com/astaxie/beego/logs"
	"fmt"
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

	logs.Debug("init success")
}