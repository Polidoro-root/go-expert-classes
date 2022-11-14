package main

import "fmt"

type Address struct {
	Neighborhood string
	Number       string
	City         string
	State        string
}

type Client struct {
	Name     string
	Age      int
	IsActive bool
	Address
}

func main() {
	joao := Client{Name: "Joao", Age: 19, IsActive: true}

	joao.IsActive = false

	joao.Number = "10"

	fmt.Println(joao)

}
