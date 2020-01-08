package main

import (
	"context"
	"fmt"
	"log"

	pd "github.com/mneko22/gRPC_chat/chat"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

var mes = ""

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("err")
	}
	defer conn.Close()
	c := pd.NewChatServiceClient(conn)
	ctx := context.Background()
	for {
		fmt.Scan(&mes)
		r, err := c.SendMes(ctx, &pd.Message{Body: mes})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Print(r.GetBody(), r.GetCount())
	}
}
