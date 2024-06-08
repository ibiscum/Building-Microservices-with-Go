package main

import (
	"context"
	"fmt"
	"os"
	"time"

	kittens "github.com/ibiscum/Building-Microservices-with-Go/chapter06/gomicro/proto"
	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

var c client.Client

func main() {
	fmt.Println("Benchmarking application")

	service := micro.NewService()
	service.Init()

	b := bench.New(true, 400, 300*time.Second, 90*time.Second, 5*time.Second)
	b.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(GoMicroRequest)
}

// GoMicroRequest is executed by benchmarks
func GoMicroRequest() error {

	// Create new request to service go.micro.srv.example, method Example.Call
	request := c.NewRequest("bmigo.micro.Kittens", "Kittens.Hello", &kittens.Request{Name: "Nic"})
	response := &kittens.Response{}

	// err := c.Call(context.TODO(), request, response, "consul.acet.io:8091",
	err := c.Call(context.TODO(), request, response)
	if err != nil {
		return err
	}

	return nil
}
