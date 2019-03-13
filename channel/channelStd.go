package main

import "fmt"
import "time"

// 通道传输数据
func testChannel() {
	message := make(chan string)

	go func() {
		message <- "ping"
	}()

	msg := <-message
	fmt.Println(msg)
}

// 通道缓存
func testChannelBuf() {
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"
	fmt.Println(<-message)
	fmt.Println(<-message)
}

func worker(done chan bool) {
	fmt.Print("Work....")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true
}

// 通道同步
func testChannelSync() {
	done := make(chan bool,1)
	go worker(done)

	fmt.Println(<-done)
}

// 单向通道声明
func ping(pings chan<- string, msg string) {
	pings<-msg
}

// 单向通道声明
func pong(pings <-chan string, pongs chan<- string) {
	msg:= <- pings
	pongs <- msg
}

// 单向通道
func testChannelDirction() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "password")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// 通道选择
func testChannelSelector() {
	c1 := make(chan string)
	c2 := make(chan string) 

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i:=0; i<2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// 对于通道的超时操作
func testChannelTimeout() {
	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second *2)
		c1 <- "time 2"
	}()

	select {
	case res:= <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1) :
		fmt.Println("timeout 1")
	}

	go func() {
		time.Sleep(time.Second *2)
		c1 <- "time 2"
	}()

	select {
	case res:= <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 3) :
		fmt.Println("timeout 3")
	}
}

func testNoBlockingChannel() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no message received")
	}
}

func testChannelClosed() {
	jobs := make(chan int, 5)
	done := make(chan bool)
	
	go func() {
		for {
			j, more := <-jobs
			if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
		}
	}()

	for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
 
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func testRangeChannel() {
	queue := make(chan string, 2)
	queue <- "one";
	queue <- "two";
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
 
func main() {
	
	// testChannel()
	// testChannelBuf()
	// testChannelSync()
	// testChannelDirction()
	// testChannelSelector()
	// testChannelTimeout()
	// testNoBlockingChannel()
	// testChannelClosed()
	testRangeChannel()
}