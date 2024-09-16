package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter01/rpc_http_json/contract"
)

func PerformRequest() contract.HelloWorldResponse {
	r, _ := http.Post(
		"http://localhost:1234",
		"application/json",
		bytes.NewBuffer([]byte(`{"id": 1, "method": "HelloWorldHandler.HelloWorld", "params": [{"name":"World"}]}`)),
	)
	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	decoder := json.NewDecoder(r.Body)
	var response contract.HelloWorldResponse
	err := decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
