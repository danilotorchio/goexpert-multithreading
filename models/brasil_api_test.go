package models_test

import (
	"testing"

	"github.com/danilotorchio/goexpert/multithreading/models"
)

func TestBrasilApiResponse_String(t *testing.T) {
	response := models.BrasilApiResponse{
		Cep:          "12345-678",
		State:        "Sao Paulo",
		City:         "Sao Paulo",
		Neighborhood: "Centro",
		Street:       "Avenida Paulista",
		Service:      "API",
	}

	expected := "Avenida Paulista, Centro - Sao Paulo, Sao Paulo, 12345-678"
	if response.String() != expected {
		t.Errorf("unexpected String() output, got: %s, want: %s", response.String(), expected)
	}
}
