package main

import "fmt"

type Address struct {
	Neighborhood string
	Number       string
	City         string
	State        string
}

type Person interface {
	Disable()
}

type Company struct {
	Name string
}

type Client struct {
	Name     string
	Age      int
	IsActive bool
	Address
}

func (company Company) Disable() {}

func (client Client) Disable() {
	client.IsActive = false
}

func Deactivation(person Person) {
	person.Disable()
}

func main() {
	joao := Client{Name: "Joao", Age: 19, IsActive: true}

	company := Company{}

	Deactivation(joao)
	Deactivation(company)

	fmt.Println(joao)

}
