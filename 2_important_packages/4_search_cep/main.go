package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Request error: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on read response: %v\n", err)
		}

		var data ViaCEP

		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on parse response: %v\n", err)
		}

		fmt.Println(data)

		dir, err := os.Getwd()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on get current path: %v\n", err)
		}

		if !strings.HasSuffix(dir, "4_search_cep") {
			dir = path.Join(dir, "4_search_cep")
		}

		file, err := os.Create(path.Join(dir, "city.txt"))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error on create file: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
	}
}
