package main

import (
	"context"
	"github.com/astaxie/beego/logs"
	pb "github.com/test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
	// OpenTLS 是否开启TLS认证
	OpenTLS = true
)

// customCredential 自定义认证
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	if OpenTLS {
		return true
	}
	return false
}

func main() {
	var err error
	var opts []grpc.DialOption

	if OpenTLS {
		//TLS连接
		creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "l12h1")
		if err != nil {
			logs.Error("创建TLS认证文件失败 %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(Address, opts...)
	if err != nil {
		logs.Error("获取服务端连接失败 %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	client := pb.NewHelloClient(conn)

	reqBody := new(pb.HelloRequest)
	reqBody.Name = "grpc"

	reponse, err := client.SayHello(context.Background(), reqBody)
	if err != nil {
		logs.Error("调用Rrpc接口出现错误 %v", err)
	}

	logs.Debug(reponse.Message)
}
