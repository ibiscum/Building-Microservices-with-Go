package main

import (
	"fmt"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http/client"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http/server"
)

func main() {
	server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}
