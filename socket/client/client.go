package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端向服务端发送数据
func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("Dialing error, err: ", err.Error())
		return
	}

	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Println("write error, err: ", err)
			return
		}
	}
}