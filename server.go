package main

import (
    "encoding/json"
    "log"
    "net"

    pb "grpcdemo/build/go" // 引入编译生成的包

    "golang.org/x/net/context"
    "google.golang.org/grpc"
)

const (
    // Address gRPC服务地址
    Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService ...
var HelloService = helloService{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    resp := new(pb.HelloReply)

    type respJsonData struct {
        Content string `json:"content"`
    }
    fieldDataInJson := new(respJsonData)
    fieldDataInJson.Content = "Hello " + in.Name + "."

    log.Println(in.Name)

    // fieldDataInJson := &respJsonData{
        // Content: "Hello " + in.Name + ".",
    // }

    respJson := struct{
        Code uint32 `json:"code"`
        Message string `json:"message"`
        *respJsonData `json:"data"`
    }{
        Code : 200,
        Message: "",
        respJsonData: fieldDataInJson,
    }
    
    message, _ := json.Marshal(respJson)
    resp.Message = string(message)

    return resp, nil
}

func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    // 实例化grpc Server
    s := grpc.NewServer()

    // 注册HelloService
    pb.RegisterHelloServer(s, HelloService)

    log.Println("Listen on " + Address)

    s.Serve(listen)
}

