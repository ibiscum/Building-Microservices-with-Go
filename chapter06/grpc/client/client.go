package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	_, err := grpc.NewClient("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("unable to create connection to server: ", err)
	}

	// client := kittens.NewKittensClient(conn)
	// response, err := client.Hello(context.Background(), &kittens.Request{Name: "Nic"})

	// if err != nil {
	// 	log.Fatal("error calling service: ", err)
	// }

	// fmt.Println(response.Msg)
}
