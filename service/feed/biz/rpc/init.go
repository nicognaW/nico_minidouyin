package rpc

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "nico_minidouyin/gen/proto/gen/feed"
)

func RpcInit() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:9919"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	s := NewServer()
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFeedServiceServer(grpcServer, s)
	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
}
