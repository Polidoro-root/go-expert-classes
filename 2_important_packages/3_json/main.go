package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"number"`
	Balance int `json:"balance"` // `json:"-"`
}

func main() {

	account := Account{Number: 1, Balance: 100}

	_, err := json.Marshal(account)

	if err != nil {
		println(err)
	}

	// println(string(res))

	encoder := json.NewEncoder(os.Stdout)

	rawJson := []byte(`{"number":2,"balance":200}`)

	var accountX Account

	json.Unmarshal(rawJson, &accountX)

	encoder.Encode(accountX)
	println(accountX.Balance)
}
