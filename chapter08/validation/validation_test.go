package validation

import (
	"testing"

	"github.com/go-playground/validator"
)

func TestErrorWhenRequestEmailNotPresent(t *testing.T) {
	validate := validator.New()
	request := Request{
		URL: "http://nicholasjackson.io",
	}

	if err := validate.Struct(&request); err == nil {
		t.Error("Should have raised an error")
	}
}

func TestErrorWhenRequestEmailIsInvalid(t *testing.T) {
	validate := validator.New()
	request := Request{
		Email: "something.com",
		URL:   "http://nicholasjackson.io",
	}

	if err := validate.Struct(&request); err == nil {
		t.Error("Should have raised an error")
	}
}

func TestNoErrorWhenRequestNameNotPresent(t *testing.T) {
	validate := validator.New()
	request := Request{
		Email: "myname@address.com",
		URL:   "http://nicholasjackson.io",
	}

	if err := validate.Struct(&request); err != nil {
		t.Error(err)
	}
}
