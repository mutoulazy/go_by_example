package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("收到信号, 停止")
				return
			default:
				fmt.Println("持续运行中")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("Done")
	// 使用context发送停止信号
	cancel()
	time.Sleep(5 * time.Second)
}
