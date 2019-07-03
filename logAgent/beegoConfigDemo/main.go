package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	conf, err := config.NewConfig("ini", "D:\\tools\\logs\\logcollect.conf")
	if err != nil {
		fmt.Println("new config err, err:", err)
		return
	}
	port, err := conf.Int("server::port")
	if err != nil {
		fmt.Println("config read server::port err, err:", err)
		return
	}
	fmt.Println("port:", port)
	log_level, err := conf.Int("log::log_level")
	if err != nil {
		fmt.Println("config read log::log_level err, err:", err)
		return
	}
	fmt.Println("log_level:", log_level)
	log_path := conf.String("log::log_path")
	fmt.Println("log_path:", log_path)
}