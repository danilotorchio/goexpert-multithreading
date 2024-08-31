package models

import "fmt"

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (a *BrasilApiResponse) String() string {
	return fmt.Sprintf("%s, %s - %s, %s, %s", a.Street, a.Neighborhood, a.City, a.State, a.Cep)
}
