package main

import (
	"fmt"
	"os"
	"time"

	proto "github.com/ibiscum/Building-Microservices-with-Go/chapter06/grpc/proto"
	"github.com/nicholasjackson/bench"
	"github.com/nicholasjackson/bench/output"
	"github.com/nicholasjackson/bench/util"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conn *grpc.ClientConn

func main() {
	conn, _ = grpc.NewClient("consul.acet.io:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	fmt.Println("Benchmarking application")

	b := bench.New(true, 400, 300*time.Second, 90*time.Second, 5*time.Second)
	b.AddOutput(301*time.Second, os.Stdout, output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./output.txt"), output.WriteTabularData)
	b.AddOutput(1*time.Second, util.NewFile("./error.txt"), output.WriteErrorLogs)
	b.AddOutput(1*time.Second, util.NewFile("./output.png"), output.PlotData)
	b.RunBenchmarks(GrpcRequest)
}

// GrpcRequest is executed by benchmarks
func GrpcRequest() error {
	client := proto.NewKittensClient(conn)
	_, err := client.Hello(context.TODO(), &proto.Request{Name: "Nic"})

	if err != nil {
		return err
	}

	return nil
}
