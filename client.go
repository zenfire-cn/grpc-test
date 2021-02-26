package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learn/grpc-learn/helloworld"
	"log"
	"time"
)

func main() {
	// 发起连接
	if c, err := grpc.Dial("127.0.0.1:6544", grpc.WithInsecure()); err != nil {
		log.Fatal(err)
	} else {
		defer c.Close()
		// 初始化客户端
		client := pb.NewHelloServiceClient(c)
		// 初始化上下文、设置连接超时
		ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

		// grpc调用
		helloRes, err := client.SayHello(ctx, &pb.HelloReq{Name: "Golang"})
		log.Println(helloRes, err)
		AddRes, err := client.Add(ctx, &pb.AddReq{A: 10, B: 22})
		log.Println(AddRes, err)
	}
}
