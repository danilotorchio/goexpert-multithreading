package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/danilotorchio/goexpert/multithreading/models"
	"github.com/danilotorchio/goexpert/multithreading/utils"
)

const (
	ResponseTimeout = time.Second * 1

	BrasilApiUrl = "https://brasilapi.com.br/api/cep/v1"
	ViaCepUrl    = "http://viacep.com.br/ws"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run main.go <cep>")
	}

	re := regexp.MustCompile("[0-9]+")
	reMatches := re.FindAllString(os.Args[1], -1)
	cep := strings.Join(reMatches, "")

	if len(cep) != 8 {
		log.Fatal("cep must be 8 digits")
	}

	ctx, cancel := context.WithTimeout(context.Background(), ResponseTimeout)
	defer cancel()

	chBrasilApi := make(chan *models.BrasilApiResponse)
	chViaCep := make(chan *models.ViaCepResponse)

	go GetAddressByBrasilApi(cep, chBrasilApi)
	go GetAddressByViaCep(cep, chViaCep)

	select {
	case res := <-chBrasilApi:
		fmt.Println("Brasil API")
		fmt.Printf("Resposta: %v\n", res.String())
	case res := <-chViaCep:
		fmt.Println("Via CEP")
		fmt.Printf("Endereço: %v\n", res.String())
	case <-ctx.Done():
		fmt.Println("timeout")
	}
}

func GetAddressByBrasilApi(cep string, ch chan<- *models.BrasilApiResponse) {
	url := fmt.Sprintf("%s/%s", BrasilApiUrl, cep)

	data, err := utils.RequestData(url)
	if err != nil {
		log.Fatal(err)
	}

	var address models.BrasilApiResponse
	if err := json.Unmarshal(data, &address); err != nil {
		log.Fatal(err)
	}

	ch <- &address
}

func GetAddressByViaCep(cep string, ch chan<- *models.ViaCepResponse) {
	url := fmt.Sprintf("%s/%s/json", ViaCepUrl, cep)

	data, err := utils.RequestData(url)
	if err != nil {
		log.Fatal(err)
	}

	var address models.ViaCepResponse
	if err := json.Unmarshal(data, &address); err != nil {
		log.Fatal(err)
	}

	ch <- &address
}
