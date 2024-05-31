package main

import (
	"encoding/json"
	"flag"
	"log"

	"github.com/nats-io/nats.go"
)

type product struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

var natsClient *nats.Conn

var natsServer = flag.String("nats", "", "NATS server URI")

func init() {
	flag.Parse()

}

func main() {
	var err error
	natsClient, err = nats.Connect("nats://" + *natsServer)
	if err != nil {
		log.Fatal(err)
	}
	defer natsClient.Close()

	log.Println("Subscribing to events")
	_, err = natsClient.Subscribe("product.inserted", handleMessage)
	if err != nil {
		log.Fatal(err)
	}
}

func handleMessage(m *nats.Msg) {
	p := product{}
	err := json.Unmarshal(m.Data, &p)
	if err != nil {
		log.Println("Unable to unmarshal event object")
		return
	}

	log.Printf("Received message: %v, %#v", m.Subject, p)
}
