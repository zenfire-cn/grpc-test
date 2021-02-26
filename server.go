package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	pb "learn/grpc-learn/helloworld"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(c context.Context, in *pb.HelloReq) (*pb.HelloRes, error) {
	pr, _ := peer.FromContext(c)
	log.Println("Remote:", pr.Addr, ", Language:", in.Name)
	return &pb.HelloRes{Message: "Hello " + in.Name}, nil
}

func (s *server) Add(c context.Context, in *pb.AddReq) (*pb.AddRes, error) {
	pr, _ := peer.FromContext(c)
	log.Println("Remote:", pr.Addr, ", params:", in)
	return &pb.AddRes{Count: in.A + in.B}, nil
}

func main() {
	if tcpListener, err := net.Listen("tcp", ":6544"); err == nil {
		// var opts []grpc.ServerOption
		// grpcServer := grpc.NewServer(opts...)
		grpcServer := grpc.NewServer()
		pb.RegisterHelloServiceServer(grpcServer, &server{})
		grpcServer.Serve(tcpListener)
	} else {
		log.Fatal(err)
	}
}
