package main

import (
	"fmt"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc/client"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
