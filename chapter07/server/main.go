package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/alexcesaro/statsd"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter07/server/handlers"
	"github.com/sirupsen/logrus"
)

const port = 8091

func main() {
	statsd, err := createStatsDClient(os.Getenv("STATSD"))
	if err != nil {
		log.Fatal("Unable to create statsD client")
	}

	logger, err := createLogger(os.Getenv("LOGSTASH"))
	if err != nil {
		log.Fatal("Unable to create logstash client")
	}

	setupHandlers(statsd, logger)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func setupHandlers(statsd *statsd.Client, logger *logrus.Logger) {
	validation := handlers.NewValidationHandler(
		statsd,
		logger,
		handlers.NewHelloWorldHandler(statsd, logger),
	)

	bangHandler := handlers.NewPanicHandler(
		statsd,
		logger,
		handlers.NewBangHandler(),
	)

	http.Handle("/helloworld", handlers.NewCorrelationHandler(validation))
	http.Handle("/bang", handlers.NewCorrelationHandler(bangHandler))
}

func createStatsDClient(address string) (*statsd.Client, error) {
	return statsd.New(statsd.Address(address))
}

func createLogger(address string) (*logrus.Logger, error) {
	retryCount := 0

	l := logrus.New()
	hostname, _ := os.Hostname()
	var err error

	// Retry connection to logstash incase the server has not yet come up
	for ; retryCount < 10; retryCount++ {
		conn, err := net.Dial("tcp", address)
		if err == nil {

			hook := logrustash.New(
				conn,
				logrustash.DefaultFormatter(
					logrus.Fields{"hostname": hostname},
				),
			)

			l.Hooks.Add(hook)
			return l, err
		}

		log.Println("Unable to connect to logstash, retrying")
		time.Sleep(1 * time.Second)
	}

	log.Fatal("Unable to connect to logstash")
	return nil, err
}
