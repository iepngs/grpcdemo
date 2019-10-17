package main

import (
	"encoding/json"
	"log"
	"net"

	pb "grpcdemo/build/go" // 引入编译生成的包

	"grpcdemo/consul"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
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

	respJson := struct {
		Code          uint32 `json:"code"`
		Message       string `json:"message"`
		*respJsonData `json:"data"`
	}{
		Code:         200,
		Message:      "",
		respJsonData: fieldDataInJson,
	}

	message, _ := json.Marshal(respJson)
	resp.Message = string(message)

	return resp, nil
}

// register
func RegisterToConsul() {
	consul.RegitserService("127.0.0.1:8500", &consul.ConsulService{
		Name: "hello",
		Tag:  []string{"hello"},
		IP:   "127.0.0.1",
		Port: 50052,
	})
}

// health
type HealthImpl struct{}

// check
// 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("health checking")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

// watch
func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
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

	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	RegisterToConsul()

	err = s.Serve(listen)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("Listen on " + Address)
}
