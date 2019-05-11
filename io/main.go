package main

import (
	"bufio"
	"fmt"
	"os"
)

type student struct {
	Name  string
	Age   int
	Score float32
}

// 从缓冲区读取终端输入
func bufferRead() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string faild, err : ", err)
		return
	}
	fmt.Printf("read success, string : %s \n", str)
}

// 从缓冲区读取文件内容
func fileBufferRead() {
	file, err := os.Open("./test.log")
	if err != nil {
		fmt.Println("open file err: ", err)
		return
	}
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string faild, err : ", err)
		return
	}
	fmt.Printf("read success, string : %s \n", str)

	file.Close()
}

func main() {
	var str = "std18 18 88.8"
	var su student
	fmt.Sscanf(str, "%s %d %f", &su.Name, &su.Age, &su.Score)
	fmt.Println(su)

	// bufferRead()
	// fileBufferRead()
}
