package main

import (
	"google.golang.org/grpc/credentials"
	"context"
	"github.com/astaxie/beego/logs"
	pb "github.com/test/proto"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

func main() {
	// TLS认证
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "l12h1")
	if err != nil {
		logs.Error("Failed to create TLS credentials %v", err)
	}
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		logs.Error("conn grpc error,err:", err)
	}
	defer conn.Close()
	// 初始化客户端
	client := pb.NewHelloClient(conn)
	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	resp, err := client.SayHello(context.Background(), reqBody)
	if err != nil {
		logs.Error("exec SayHello error,err:", err)
	}
	
	logs.Debug(resp.Message)
}
