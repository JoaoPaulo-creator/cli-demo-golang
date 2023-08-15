package main

import (
	"encoding/json"
	"flag"
	"fmt"

	//"io"
	"net/http"
	"os"
)

type Response struct {
	Cep          string `json:"cep"`
	Address_name string `json:"address_name"`
}

var help bool
var request bool

func main() {

	flag.BoolVar(&help, "help", false, "Show help")
	flag.BoolVar(&request, "request", false, "Make an http request")

	flag.Parse()

	res, err := http.Get("https://cep.awesomeapi.com.br/json/80010010")

	if err != nil {
		fmt.Println("Deu ruim")
		return
	}
	defer res.Body.Close()

	if help {
		// Print help message and exit
		flag.PrintDefaults()
		os.Exit(0)
	}

	//body, err := io.ReadAll(res.Body)
	//responseBody := string(body)

	var response Response
	json.NewDecoder(res.Body).Decode(&response)

	if request {
		fmt.Println("Cep do maluco: ", response.Cep)
		fmt.Println("Endereco do maluco : ", response.Address_name)
		os.Exit(0)
	}

	fmt.Println("Welcome to the Go CLI App!")

	if flag.NArg() > 0 {
		fmt.Printf("Arguments: %s\n", flag.Args())
	}
}
