package main

import (
	"testing"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http_json/client"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http_json/server"
)

func BenchmarkHelloWorldHandlerJSONRPC(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = client.PerformRequest()

	}
}

func init() {
	// start the server
	go server.StartServer()
}
