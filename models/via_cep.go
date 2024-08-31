package models

import "fmt"

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (a *ViaCepResponse) String() string {
	return fmt.Sprintf("%s, %s - %s, %s, %s", a.Logradouro, a.Bairro, a.Localidade, a.Uf, a.Cep)
}
