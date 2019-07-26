package main

import (
	"fmt"
	"context"
	"github.com/astaxie/beego/logs"
	pb "github.com/test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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
	// 解析metada中的信息并验证
	metaData, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appid  string
		appkey string
	)
	if val, ok := metaData["appid"]; ok {
		appid = val[0]
	}
	if val, ok := metaData["appkey"]; ok {
		appkey = val[0]
	}
	if appid != "101010" || appkey != "i am key" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", in.Name, appid, appkey)
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

	logs.Debug("Listen on " + Address + " with TLS and Token")
	s.Serve(listen)
}
