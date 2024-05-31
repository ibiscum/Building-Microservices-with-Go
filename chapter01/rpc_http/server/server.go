package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http/contract"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	err := rpc.Register(helloWorld)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to listen on given port: %s", err))
	}

	log.Printf("Server starting on port %v\n", port)

	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal(err)
	}
}
