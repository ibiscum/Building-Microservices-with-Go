package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/VividCortex/ewma"
)

var ma ewma.MovingAverage
var threshold = 1000 * time.Millisecond
var timeout = 1000 * time.Millisecond
var resetting = false
var resetMutex = sync.RWMutex{}

func main() {
	ma = ewma.NewMovingAverage()

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/health", healthHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func mainHandler(rw http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if !isHealthy() {
		respondServiceUnhealthy(rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Average request time: %f (ms)\n", ma.Value()/1000000)

	duration := time.Since(startTime)

	ma.Add(float64(duration))
}

func healthHandler(rw http.ResponseWriter, r *http.Request) {
	if !isHealthy() {
		rw.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	fmt.Fprint(rw, "OK")
}

func isHealthy() bool {
	return (ma.Value() < float64(threshold))
}

func respondServiceUnhealthy(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusServiceUnavailable)

	resetMutex.RLock()
	defer resetMutex.RUnlock()

	if !resetting {
		go sleepAndResetAverage()
	}
}

func sleepAndResetAverage() {
	resetMutex.Lock()
	resetting = true
	resetMutex.Unlock()

	time.Sleep(timeout)
	ma = ewma.NewMovingAverage()

	resetMutex.Lock()
	resetting = false
	resetMutex.Unlock()
}
