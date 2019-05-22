package main

import (
	"net"
	"fmt"
)

// 处理连接协程
func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error, err: ", err)
			return
		}
		fmt.Println("read: ", string(buf))
	}
}

func main() {
	fmt.Println("start server....")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen error, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error, err: ", err)
			continue
		}
		go process(conn) 
	}
}