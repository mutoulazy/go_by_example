package main

import (
	"../balance"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// 接口实例
// 模拟随机和轮询负载均衡方法
func main() {
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		port := 80
		host := fmt.Sprintf("192.168.%d.%d", i, rand.Intn(255))
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	var conf = "round"

	if len(os.Args) > 1 {
		conf = os.Args[1]
	}

	// 日志输入到文件中
	file, err := os.OpenFile("./run.log", os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}

	for {
		inst, err := balance.DoBalance(conf, insts)
		if err != nil {
			// fmt.Println("do balance error: ", err)
			fmt.Fprint(file, "do balance error\n")
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}

	file.Close()
}
