package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

func main() {
	h := md5.New()
	_, err := io.WriteString(h, "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x", h.Sum(nil))
}
