package main

import (
	"context"
	"fmt"
	"time"
)

var key = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	valueCtx1 := context.WithValue(ctx, key, "No 1")
	valueCtx2 := context.WithValue(ctx, key, "No 2")
	valueCtx3 := context.WithValue(ctx, key, "No 3")
	go watch(valueCtx1)
	go watch(valueCtx2)
	go watch(valueCtx3)

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
	// 发送停止信号
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Value(key), " 收到信号, 停止")
			return
		default:
			fmt.Println(ctx.Value(key), " 持续运行中")
			time.Sleep(2 * time.Second)
		}

	}
}
