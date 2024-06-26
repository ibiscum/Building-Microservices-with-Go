package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
)

var client *http.Client

func main() {
	fmt.Println("Benchmarking application")

	client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
		},
		Timeout: 5 * time.Second,
	}

	b := bench.New(true, 10, 30*time.Second, 15*time.Second, 5*time.Second)
	b.AddOutput(0*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(Request)
}

// Request is executed by benchmarks
func Request() error {

	serverURI := "http://" + os.Getenv("DOCKER_IP") + ":8323/search"

	req, err := http.NewRequest(
		"POST",
		serverURI,
		bytes.NewBuffer([]byte(`{"query":"Fat Freddy's Cat"}`)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed with status: %v", resp.Status)
	}

	return nil
}
