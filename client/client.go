package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "skywire/proto"
)

var (
	serverAddr = flag.String("server_addr", "localhost:1967", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	var conn, err = grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewSkywireClient(conn)

	req := pb.SkywireRequest{Meta: &pb.RequestMetadata{Id: 1234}}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.Generate(ctx, &req)
	log.Println(stream)
}
