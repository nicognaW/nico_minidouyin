package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "nico_minidouyin/gen/douyin/feed"
	"nico_minidouyin/service/feed/app"
)

func newServer() *app.FeedService {
	s := &app.FeedService{}
	return s
}

func main() {
	log.Printf("starting feed service...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:40127"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFeedServiceServer(grpcServer, newServer())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
