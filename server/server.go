//go:generate protoc -Iproto --go_out=plugins=grpc:proto proto/skywire.proto

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "skywire/proto"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 1967, "The server port")
)

type skywireServer struct {
	pb.UnimplementedSkywireServer
}

func (s *skywireServer) Generate(ctx context.Context, req *pb.SkywireRequest) (*pb.SkywireResponse, error) {
	res := pb.SkywireResponse{Meta: &pb.RequestMetadata{Id: 0}}
	res.GetMeta().Id = req.GetMeta().Id
	log.Println(res)
	return &res, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	swServer := &skywireServer{}
	pb.RegisterSkywireServer(grpcServer, swServer)
	grpcServer.Serve(lis)
}
