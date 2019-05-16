package main

import (
	"fmt"
)

// 判断素数
// 通过Channel完成协程之间的数据共享
func calc(taskChan chan int, resultChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resultChan <- v
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)
	go func() {
		for i := 0; i < 100000; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	for j := 0; j < 8; j++ {
		go calc(intChan, resultChan, exitChan)
	}

	// 等待所有协程完成关闭通道
	go func() {
		for e := 0; e < 8; e++ {
			<-exitChan
			fmt.Println("go route", e, " exit")
		}
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}
