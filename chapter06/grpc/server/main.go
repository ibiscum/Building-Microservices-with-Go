package main

import (
	"fmt"
	"log"
	"net"

	// context "golang.org/x/net/context"

	"google.golang.org/grpc"
	// kittens "github.com/ibiscum/Building-Microservices-with-Go/chapter06/grpc/proto/kittens"
)

// type kittenServer struct {
// 	// kittens.UnimplementedKittensServer
// }

// func (k *kittenServer) Hello(ctx context.Context, request *kittens.Request) (*kittens.Response, error) {
// 	response := &kittens.Response{}
// 	response.Msg = fmt.Sprintf("Hello %v", request.Name)

// 	return response, nil
// }

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	// kittens.RegisterKittensServer(grpcServer, &kittenServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
