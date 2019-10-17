package main

import (
	"log"
	"os"
	"time"

	"grpcdemo/consul"

	pb "grpcdemo/build/go" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	target = "consul://127.0.0.1:8500/hello"
	// Address gRPC服务地址
	// Address = "127.0.0.1:50052"
)

func main() {
	consul.Init()
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// 连接
	conn, err := grpc.DialContext(
		ctx,
		target,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithBalancerName("round_robin"))

	if err != nil {
		log.Fatalln("did not connect: %v", err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	if len(os.Args) > 1 {
		reqBody.Name = os.Args[1]
	}

	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, reqBody)
		if err != nil {
			log.Fatalf("could not call SayHello(): %v", err)
		}
		log.Printf("SayHello() response:%v", r.Message)
		time.Sleep(2 * time.Second)
	}

}
