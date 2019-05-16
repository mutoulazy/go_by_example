package main 

// 通过Channel完成协程之间的数据共享
func main() {
	var intChan chan int 
	intChan = make(chan int, 10)
	intChan <- 10
}