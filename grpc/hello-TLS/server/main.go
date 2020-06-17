package main

import (
	"google.golang.org/grpc/credentials"
	"context"
	"github.com/astaxie/beego/logs"
	pb "github.com/test/proto"
	"google.golang.org/grpc"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

var myHelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = "Hello " + in.Name + "."
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		logs.Error("failed to listen: %v", err)
	}
	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		logs.Error("Failed to generate credentials %v", err)
	}
	// 实例化grpc Server
	s := grpc.NewServer(grpc.Creds(creds))
	// 注册HelloService
	pb.RegisterHelloServer(s, myHelloService)

	logs.Debug("Listen on " + Address + " with TLS")
	s.Serve(listen)
}
