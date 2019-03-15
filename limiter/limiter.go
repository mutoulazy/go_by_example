package main 

import "time"
import "fmt"


// 统一采用打点器限制速率
func limiter1() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i ++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Millisecond * 200	)

	for req := range requests {
		<- limiter
		fmt.Println("request", req, time.Now())
	}
}

// 前三个不受限制 后面的进行限制
func limiter2() {
	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
        burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}


func main() {
	// limiter1()
	limiter2()
}