package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"testing"
)

type Response struct {
	Message string
}

func BenchmarkHelloHandlerVariable(b *testing.B) {
	b.ResetTimer()

	var writer = io.Discard
	response := Response{Message: "Hello world"}

	for i := 0; i < b.N; i++ {
		data, _ := json.Marshal(response)
		fmt.Fprint(writer, string(data))
	}
}

func BenchmarkHelloHandlerEncoder(b *testing.B) {
	b.ResetTimer()

	var writer = io.Discard
	response := Response{Message: "Hello world"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(response)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkHelloHandlerEncoderReference(b *testing.B) {
	b.ResetTimer()

	var writer = io.Discard
	response := Response{Message: "Hello world"}

	for i := 0; i < b.N; i++ {
		encoder := json.NewEncoder(writer)
		err := encoder.Encode(&response)
		if err != nil {
			log.Fatal(err)
		}
	}
}
