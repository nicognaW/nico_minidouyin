package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "nico_minidouyin/gen/douyin/user"
	"nico_minidouyin/service/user/app"
)

func main() {
	log.Printf("starting user service...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:40128"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	s := &app.UserService{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
