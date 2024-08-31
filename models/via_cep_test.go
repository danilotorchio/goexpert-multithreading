package models_test

import (
	"testing"

	"github.com/danilotorchio/goexpert/multithreading/models"
)

func TestViaCepResponse_String(t *testing.T) {
	response := models.ViaCepResponse{
		Cep:         "12345-678",
		Logradouro:  "Avenida Paulista",
		Complemento: "Complemento",
		Unidade:     "Unidade",
		Bairro:      "Centro",
		Localidade:  "Sao Paulo",
		Uf:          "Sao Paulo",
		Ibge:        "1234567",
		Gia:         "1234",
		Ddd:         "11",
		Siafi:       "1234",
	}

	expected := "Avenida Paulista, Centro - Sao Paulo, Sao Paulo, 12345-678"
	if response.String() != expected {
		t.Errorf("unexpected String() output, got: %s, want: %s", response.String(), expected)
	}
}
