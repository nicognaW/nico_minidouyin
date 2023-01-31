package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pbAuth "nico_minidouyin/gen/douyin/auth"
	pbUser "nico_minidouyin/gen/douyin/user"
	"nico_minidouyin/service/user/app"
)

func main() {
	log.Printf("starting user service...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:40128"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	userService := &app.UserService{}
	authService := &app.AuthService{}
	grpcServer := grpc.NewServer(opts...)
	pbUser.RegisterUserServiceServer(grpcServer, userService)
	pbAuth.RegisterAuthServiceServer(grpcServer, authService)

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
