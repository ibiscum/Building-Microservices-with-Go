package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/alexcesaro/statsd"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter07/server/entities"
	"github.com/ibiscum/Building-Microservices-with-Go/chapter07/server/httputil"
	"github.com/sirupsen/logrus"
)

type validationHandler struct {
	next   http.Handler
	statsd *statsd.Client
	logger *logrus.Logger
}

// NewValidationHandler creates a new Validation handler with the given statsD client and
// logger.
func NewValidationHandler(statsd *statsd.Client, logger *logrus.Logger, next http.Handler) http.Handler {
	return &validationHandler{next: next, statsd: statsd, logger: logger}
}

func (h *validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request entities.HelloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		h.statsd.Increment(validationFailed)

		message := httputil.SerialzableRequest{Request: r}
		h.logger.WithFields(logrus.Fields{
			"handler": "Validation",
			"status":  http.StatusBadRequest,
			"method":  r.Method,
		}).Info(message.ToJSON())

		http.Error(rw, "Bad request", http.StatusBadRequest)

		return
	}

	type contextKey string
	const contextKeyRequestID contextKey = "name"

	c := context.WithValue(r.Context(), contextKeyRequestID, request.Name)
	r = r.WithContext(c)

	h.statsd.Increment(validationSuccess)
	h.next.ServeHTTP(rw, r)
}
