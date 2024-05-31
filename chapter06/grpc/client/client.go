package main

import (
	"fmt"
	"log"

	proto "github.com/ibiscum/Building-Microservices-with-Go/chapter06/grpc/proto"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Unable to create connection to server: ", err)
	}

	client := proto.NewKittensClient(conn)
	response, err := client.Hello(context.Background(), &proto.Request{Name: "Nic"})

	if err != nil {
		log.Fatal("Error calling service: ", err)
	}

	fmt.Println(response.Msg)
}
