package main

import (
	"time"
	"github.com/astaxie/beego/logs"
	"../tailf"
	"../kafka"
)

func runServer() (err error) {
	logs.Info("start server...")
	for {
		msg := tailf.GetOneLine()
		err = sendTokafka(msg)
		if err != nil {
			logs.Error("send msg to kafka error,%v", err)
			time.Sleep(time.Second)
			return
		}
	}
	return
}

func sendTokafka(msg *tailf.TextMsg) (err error) {
	// logs.Debug("msg:%s, topic:%s", msg.Msg, msg.Topic)
	err = kafka.SendMsgToKafka(msg.Msg, msg.Topic)
	return
}