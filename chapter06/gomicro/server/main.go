package main

import (
	"context"
	"log"

	kittens "github.com/ibiscum/Building-Microservices-with-Go/chapter06/gomicro/proto"
	"go-micro.dev/v4/server"
	"go-micro.dev/v4/util/cmd"
)

type Kittens struct{}

func (s *Kittens) Hello(ctx context.Context, req *kittens.Request, rsp *kittens.Response) error {
	rsp.Msg = server.DefaultId + ": Hello " + req.Name

	return nil
}

// var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {

	err := cmd.Init()
	if err != nil {
		log.Fatal(err)
	}

	server.Init(
		server.Name("bmigo.micro.Kittens"),
		server.Version("1.0.0"),
		server.Address(":8091"),
	)

	// Register Handlers
	err = server.Handle(
		server.NewHandler(
			new(Kittens),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
