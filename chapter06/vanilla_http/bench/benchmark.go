package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter06/vanilla_http/entities"
	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
)

func main() {
	fmt.Println("Benchmarking application")

	b := bench.New(true, 10, 100*time.Second, 50*time.Second, 5*time.Second)
	b.AddOutput(70*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(GoMicroRequest)
}

// GoMicroRequest is executed by benchmarks
func GoMicroRequest() error {

	request := entities.HelloWorldRequest{
		Name: "Nic",
	}

	data, _ := json.Marshal(request)

	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 5,
		},
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed with status: %v", resp.Status)
	}

	return nil
}
