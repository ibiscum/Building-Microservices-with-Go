package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc/contract"
)

const port = 1234

// func main() {
// 	log.Printf("Server starting on port %v\n", port)
// 	StartServer()
// }

func StartServer() {
	helloWorld := &HelloWorldHandler{}
	err := rpc.Register(helloWorld)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(args *contract.HelloWorldRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + args.Name
	return nil
}
