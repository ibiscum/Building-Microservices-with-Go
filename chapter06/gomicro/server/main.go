package main

import (
	"context"
	"flag"

	log "github.com/golang/glog"
	kittens "github.com/ibiscum/Building-Microservices-with-Go/chapter06/gomicro/proto"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
)

type Kittens struct{}

func (s *Kittens) Hello(ctx context.Context, req *kittens.Request, rsp *kittens.Response) error {
	rsp.Msg = server.DefaultId + ": Hello " + req.Name

	return nil
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	cmd.Init()

	server.Init(
		server.Name("bmigo.micro.Kittens"),
		server.Version("1.0.0"),
		server.Address(":8091"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(Kittens),
		),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
