package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("收到信号, 停止")
				return
			default:
				fmt.Println("持续运行中")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
	// 发送停止信号
	stop <- true
	time.Sleep(5 * time.Second)
}
