package main 

import (
	"fmt"
	"math/rand"
	"../balance"
	"time"
	"os"
)

// 接口实例
// 模拟随机和轮询负载平衡方法
func main() {
	var insts []*balance.Instance;
	for i:=0; i<16; i++ {
		port := 80
		host := fmt.Sprintf("192.168.%d.%d", i, rand.Intn(255)) 
		one := balance.NewInstance(host, port)
		insts = append(insts, one)
	}

	var balancer balance.Balancer
	var conf = "round"

	if len(os.Args) > 1 {
		conf = os.Args[1]
	}

	if conf == "random" {
		balancer = &balance.RandonBalance{}
		fmt.Println("use RandonBalance")
	} else if conf == "round" {
		balancer = &balance.RoundomBalance{}
		fmt.Println("use RoundomBalance")
	} else {
		balancer = &balance.RoundomBalance{}
		fmt.Println("use RoundomBalance")
	}

	for {
		inst, err := balancer.DoBalance(insts)
		if err != nil {
			fmt.Println("do balance error: ", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}