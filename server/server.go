package main

import (
	"context"
	"log"
	"net"

	pd "github.com/mneko22/gRPC_chat/chat"
	"google.golang.org/grpc"
)

const port = ":50051"

var count int32 = 0

type server struct {
	pd.UnimplementedChatServiceServer
}

func (s server) SendMes(ctx context.Context, req *pd.Message) (*pd.ReMessage, error) {
	count++
	log.Printf("Count: %v, Receive: %v", count, req.GetBody())
	return &pd.ReMessage{Body: req.GetBody(), Count: count}, nil
}

func main() {
	var lis, err = net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen")
		return
	}
	var s = grpc.NewServer()
	pd.RegisterChatServiceServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal("failed to up server")
		return
	}
}
