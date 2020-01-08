package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pd "github.com/mneko22/gRPC_chat/chat"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("err")
	}
	defer conn.Close()
	c := pd.NewChatServiceClient(conn)
	ctx := context.Background()
	go func() {
		stream, err := c.BloadcastMessage(ctx, &pd.Empty{})
		if err != nil {
			log.Fatal(err)
			ctx.Err()
			return
		}
		for {
			mes, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				ctx.Err()
				return
			}
			log.Print(mes)
		}
	}()
	mes := ""
	for {
		fmt.Scan(&mes)
		r, err := c.SendMes(ctx, &pd.Message{Body: mes})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("%v: %v", r.GetCount(), r.GetBody())
	}

}
