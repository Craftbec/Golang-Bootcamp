package main

import (
	"context"
	"fmt"
	"log"
	"military/transmitter"
	"time"

	"io"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	client := transmitter.NewTransmitterClient(conn)
	tr := transmitter.TransmitterRequest{}
	stream, err := client.Transmitter(context.Background(), &tr)
	if err != nil {
		log.Fatalln(err)
	}

	for i := 1; ; i++ {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%d: %s %f %v\n", i, resp.GetSessionId(), resp.GetFrequency(), resp.GetTimestamp().AsTime())
		time.Sleep(time.Millisecond * 1)
	}
}
