package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	roots := x509.NewCertPool()

	rootCert, err := os.ReadFile("../generate_keys/root_cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	ok := roots.AppendCertsFromPEM(rootCert)
	if !ok {
		panic("failed to parse root certificate")
	}

	applicationCert, err := os.ReadFile("../generate_keys/application_cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	ok = roots.AppendCertsFromPEM(applicationCert)
	if !ok {
		panic("failed to parse root certificate")
	}

	tlsConf := &tls.Config{RootCAs: roots}

	tr := &http.Transport{TLSClientConfig: tlsConf}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:8433")
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
