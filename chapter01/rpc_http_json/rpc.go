package main

import "github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http_json/server"

// To execute a request with this server run the below command on your command line
// curl -X POST -H "Content-Type: application/json" -d '{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}' http://localhost:1234
func main() {
	server.StartServer()
}
