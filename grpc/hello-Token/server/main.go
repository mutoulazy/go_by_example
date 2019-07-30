package main

import (
	"context"
	"github.com/astaxie/beego/logs"
	pb "github.com/test/proto"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"net"
	"net/http"
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

func auth(ctx context.Context) error {
	// 解析metada中的信息并验证
	metaData, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
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
		return grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}
	return nil
}

// trace方法
func startTrace() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	go http.ListenAndServe(":50051", nil)
	logs.Debug("Trace listen on 50051")
}

func main() {
	var opts []grpc.ServerOption
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		logs.Error("failed to listen: %v", err)
	}
	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		logs.Error("Failed to generate credentials %v", err)
	}
	opts = append(opts, grpc.Creds(creds))

	// 注册interceptor
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err = auth(ctx)
		if err != nil {
			return
		}
		return handler(ctx, req)
	}
	opts = append(opts, grpc.UnaryInterceptor(interceptor))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)
	// 注册HelloService
	pb.RegisterHelloServer(s, myHelloService)

	// 开启trace
	grpc.EnableTracing = true
	go startTrace()

	logs.Debug("Listen on " + Address + " with TLS and Token and + Interceptor")
	s.Serve(listen)
}
