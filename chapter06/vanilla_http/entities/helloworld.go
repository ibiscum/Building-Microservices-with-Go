package entities

type JsonString string

// HelloWorldResponse defines a response returned from the /helloworld endpoint
type HelloWorldResponse struct {
	Message JsonString `json:"message"`
}

// HelloWorldRequest defines a request sent to the /helloworld endpoint
type HelloWorldRequest struct {
	Name JsonString `json:"name"`
}
