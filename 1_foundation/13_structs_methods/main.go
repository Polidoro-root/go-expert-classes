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

func (client *Client) Disable() {
	client.IsActive = false
}

func main() {
	joao := Client{Name: "Joao", Age: 19, IsActive: true}

	joao.Disable()

	fmt.Println(joao)

}
