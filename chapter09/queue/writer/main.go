package main

import (
	"io"
	"log"
	"net/http"

	"github.com/ibiscum/Building-Microservices-with-Go/chapter09/queue"
)

type Product struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

func main() {
	q, err := queue.NewRedisQueue("redis:6379", "test_queue")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		err := q.Add("new.product", data)
		if err != nil {
			log.Println(err)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	err = http.ListenAndServe(":8000", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}
